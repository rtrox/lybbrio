package commands

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	chi_middleware "github.com/go-chi/chi/middleware"
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
	"lybbrio/internal/metrics"
	"lybbrio/internal/middleware"
	"lybbrio/internal/task"
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

func rootRun(cmd *cobra.Command, args []string) {
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
	cal, error := calibre.NewCalibreSQLite(conf.CalibreDBPath)
	cal = cal.WithLogger(&log.Logger).(*calibre.CalibreSQLite)

	if error != nil {
		log.Fatal().Err(error).Msg("Failed to initialize Calibre")
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
	workerPool := task.NewWorkerPool(
		client,
		&task.WorkerPoolConfig{
			Ctx:         schedulerVC,
			NumWorkers:  conf.Task.Workers,
			QueueLength: conf.Task.QueueLength,
		},
	)
	workerPool.Start()

	scheduler := task.NewScheduler(
		client,
		&task.SchedulerConfig{
			Ctx:       schedulerVC,
			WorkQueue: workerPool.WorkQueue(),
			Cadence:   conf.Task.Cadence,
		},
	)
	scheduler.RegisterTasks(calibre_tasks.TaskMap(cal, client))
	scheduler.Start()

	// Auth Provider
	jwtProvider, err := auth.NewJWTProvider(conf.JWTSecret, conf.JWTIssuer, conf.JWTExpiry)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize JWT Provider")
	}

	// GraphQL
	graphqlHandler := handler.NewDefaultServer(graph.NewSchema(client))
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

	r.Mount("/auth", auth.Routes(client, jwtProvider))
	r.Route("/graphql", func(r chi.Router) {
		r.With(
			auth.Middleware(jwtProvider),
			middleware.ViewerContextMiddleware(client),
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
