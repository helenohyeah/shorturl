package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/helen/lumen5_miniurl/config"
)

// Command-line flags.
var (
	configPath string
)

func main() {
	// Parse command-line flags
	flag.StringVar(&configPath, "config", "config.yml", "Path to the configuration `file`")
	flag.Parse()

	// Set up logger
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log.Logger = zerolog.New(output).With().Timestamp().Caller().Logger()

	// Load config
	log.Info().Msgf("Loading configuration from: %s", configPath)
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to load config file: %v", err)
	}

	log.Info().Msg("Setting up routes...")

	r := chi.NewRouter()
	middlewares := []func(http.Handler) http.Handler{
		// hlog.NewHandler(log.Logger),
		// hlog.URLHandler("url"),
		// hlog.MethodHandler("method"),
	}
	r.Use(middlewares...)

	// Unauthenticated routes
	// redirect
	// create short url

	// register
	// login
	// logout

	// Authenticated routes
	// users/{userID}/urls - get urls for user

	server := &http.Server{
		Addr: cfg.ListenAddress,
		// WriteTimeout:      time.Second * 600,
		// ReadHeaderTimeout: time.Second * 60,
		// ReadTimeout:       time.Second * 60,
		// IdleTimeout:       time.Second * 120,
		Handler: r,
	}

	// Start the HTTP server.
	log.Info().Msgf("HTTP server listening on %v", cfg.ListenAddress)
	if err := server.ListenAndServe(); err != nil {
		log.Info().Err(err).Msg("Terminating server...")
	}
}
