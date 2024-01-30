package commands

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"entgo.io/contrib/entgql"
	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"lybbrio/internal/auth"
	"lybbrio/internal/calibre"
	calibre_tasks "lybbrio/internal/calibre/tasks"
	"lybbrio/internal/config"
	"lybbrio/internal/db"
	"lybbrio/internal/graph"
	"lybbrio/internal/handler"
	"lybbrio/internal/metrics"
	"lybbrio/internal/middleware"
	"lybbrio/internal/scheduler"
	"lybbrio/internal/viewer"
)

type AppInfo struct {
	Name      string
	Version   string
	Revision  string
	BuildTime string
}

var (
	appInfo = AppInfo{}
	conf    = &config.Config{}
	rootCmd = &cobra.Command{
		Use:   appInfo.Name,
		Short: "Lybbr.io: A modern web API and UI for Calibre",
		Long:  `Lybbr.io is a modern web API and UI for Calibre, written in golang & React.`,
		Run:   rootRun,
	}
)

func Execute(a AppInfo) error {
	appInfo = a
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig, initLogger)

	config.RegisterFlags(rootCmd.PersistentFlags())
}

func initConfig() {
	var err error
	conf, err = config.LoadConfig(rootCmd.PersistentFlags())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n\n", err)
		err := rootCmd.Usage()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to print usage")
		}
		os.Exit(1)
	}
	if err := conf.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n\n", err)
		err := rootCmd.Usage()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to print usage")
		}
		os.Exit(1)
	}
}

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	lvl, err := zerolog.ParseLevel(conf.LogLevel)
	if err != nil {
		// should be unreachable as config is validated in initConfig()
		log.Fatal().Err(err).Msg("Failed to parse log level")
	}
	zerolog.SetGlobalLevel(lvl)
	if conf.LogFormat == "text" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

func rootRun(_ *cobra.Command, _ []string) {
	var srv http.Server

	schedulerCtx := context.Background()
	idleConnsClosed := make(chan struct{})

	go func() {
		sigchan := make(chan os.Signal, 1)

		signal.Notify(sigchan, os.Interrupt)
		signal.Notify(sigchan, syscall.SIGTERM)
		sig := <-sigchan
		log.Info().
			Str("signal", sig.String()).
			Msg("Stopping in response to signal")
		ctx, cancel := context.WithTimeout(context.Background(), conf.GracefulShutdownTimeout)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal().Err(err).Msg("Failed to gracefully close http server")
		}
		close(idleConnsClosed)
	}()

	log.Info().
		Str("name", appInfo.Name).
		Str("version", appInfo.Version).
		Str("build_time", appInfo.BuildTime).
		Str("revision", appInfo.Revision).
		Msg("App Started.")
	log.Debug().Interface("config", conf).Msg("Loaded config")

	appFunc := metrics.AppInfoGaugeFunc(metrics.AppInfoOpts{
		Name:      appInfo.Name,
		Version:   appInfo.Version,
		BuildTime: appInfo.BuildTime,
		Revision:  appInfo.Revision,
	})

	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(appFunc)

	if conf.GoCollector {
		reg.MustRegister(collectors.NewGoCollector())
	}
	if conf.ProcessCollector {
		reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	}

	// Calibre
	cal, err := calibre.NewCalibreSQLite(conf.CalibreLibraryPath)
	cal = cal.WithLogger(&log.Logger)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize Calibre")
	}

	// Database
	client, err := db.Open(&conf.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("opening ent client")
	}
	// TODO: Better Migrations
	if err := client.Schema.Create(
		context.Background(),
	); err != nil {
		log.Fatal().Err(err).Msg("migrating ent client")
	}

	// Task Scheduler
	schedulerVC := viewer.NewSystemAdminContext(schedulerCtx)
	workerPool := scheduler.NewWorkerPool(
		client,
		&scheduler.WorkerPoolConfig{
			Ctx:         schedulerVC,
			NumWorkers:  conf.Task.Workers,
			QueueLength: conf.Task.QueueLength,
		},
	)
	workerPool.Start()

	scheduler := scheduler.NewScheduler(
		client,
		&scheduler.SchedulerConfig{
			Ctx:       schedulerVC,
			WorkQueue: workerPool.WorkQueue(),
			Cadence:   conf.Task.Cadence,
		},
	)
	scheduler.RegisterTasks(calibre_tasks.TaskMap(cal, client))
	scheduler.Start()

	// Auth Provider
	var kc auth.KeyContainer
	switch conf.JWT.SigningMethod {
	case "HS512":
		kc, err = auth.NewHS512KeyContainer(conf.JWT.HMACSecret)
	case "RS512":
		kc, err = auth.NewRS512KeyContainer(conf.JWT.RSAPrivateKey, conf.JWT.RSAPublicKey)
	}
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize Key Container")
	}

	jwtProvider, err := auth.NewJWTProvider(kc, conf.JWT.Issuer, conf.JWT.Expiry, conf.JWT.RefreshExpiry)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize JWT Provider")
	}

	// GraphQL
	graphqlHandler := graphql_handler.NewDefaultServer(graph.NewSchema(
		client,
		conf.Argon2ID,
	))
	graphqlHandler.Use(
		entgql.Transactioner{TxOpener: client},
	)

	// HTTP
	r := chi.NewRouter()

	r.Use(middleware.DefaultStructuredLogger()) // Must be first, as it initializes the log ctx.
	r.Use(chi_middleware.RealIP)
	r.Use(chi_middleware.RedirectSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.Prometheus(reg))
	r.Use(chi_middleware.Recoverer)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"status": "ok"})
	})

	r.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	r.Mount("/auth", handler.AuthRoutes(client, jwtProvider, conf.Argon2ID))

	r.With(
		middleware.ViewerContextMiddleware(jwtProvider),
	).Mount("/download",
		handler.DownloadRoutes(client))

	r.Route("/graphql", func(r chi.Router) {
		r.With(
			middleware.ViewerContextMiddleware(jwtProvider),
			middleware.SuperRead,
		).Handle("/", graphqlHandler)
		r.Handle("/playground", playground.Handler("Lybbrio GraphQL playground", "/graphql"))
	})

	srv.Addr = fmt.Sprintf("%s:%d", conf.Interface, conf.Port)
	srv.Handler = r

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}

	<-idleConnsClosed
	log.Info().Msg("App Stopped.")
}
