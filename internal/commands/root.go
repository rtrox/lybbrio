package commands

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	chi_middleware "github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"lybbrio/docs"
	"lybbrio/internal/config"
	"lybbrio/internal/handlers"
	"lybbrio/internal/metrics"
	"lybbrio/internal/middleware"
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
	cobra.OnInitialize(initConfig, initLogger, initDocs)

	rootCmd.PersistentFlags().String("log-level", "info", "Log level")
	rootCmd.PersistentFlags().String("log-format", "text", "Log format (text, json)")
	rootCmd.PersistentFlags().Duration("graceful-shutdown-timeout", 1*time.Second, "Graceful shutdown timeout prior to killing the process")
	rootCmd.PersistentFlags().String("base-url", "http://localhost:8080", "Base URL")
	rootCmd.PersistentFlags().Bool("go-collector", false, "Enable Go prometheus collector")
	rootCmd.PersistentFlags().Bool("process-collector", false, "Enable process prometheus collector")
	rootCmd.PersistentFlags().Int("port", 8080, "Port to Listen On")
	rootCmd.PersistentFlags().String("interface", "127.0.0.1", "Interface")

}

func initDocs() {
	docs.SwaggerInfo.Title = rootCmd.Short
	docs.SwaggerInfo.Description = rootCmd.Long
	docs.SwaggerInfo.Version = "0.1.0" // API Version, differs from app version
	docs.SwaggerInfo.Host = conf.BaseURL
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
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

	// Database

	// Stores

	// Auth Provider

	// HTTP
	r := chi.NewRouter()

	r.Use(chi_middleware.RealIP)
	r.Use(chi_middleware.RedirectSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.Prometheus(reg))
	r.Use(middleware.DefaultStructuredLogger())
	r.Use(chi_middleware.Recoverer)

	r.Get("/healthz", handlers.Health)
	r.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	r.Handle("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	srv.Addr = fmt.Sprintf("%s:%d", conf.Interface, conf.Port)
	srv.Handler = r

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}

	<-idleConnsClosed
	log.Info().Msg("App Stopped.")
}
