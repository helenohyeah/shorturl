package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	ghandlers "github.com/gorilla/handlers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/helen/lumen5_miniurl/config"
	"github.com/helen/lumen5_miniurl/db"
	"github.com/helen/lumen5_miniurl/handlers"
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

	// Connect to redis
	// rdb := store.InitializeStore()
	// defer func() {
	// 	// Close Redis Pool when the app exited
	// 	if err := rdb.RedisDB.Close(); err != nil {
	// 		log.Error().Err(err).Msg("Error closing redis")
	// 	}
	// }()

	dB := db.DB{}
	dB.Connect(*cfg)
	defer func() {
		// Close DB when app exits
		if err := dB.DB.Close(); err != nil {
			log.Error().Err(err).Msg("Error closing DB connection")
		}
	}()
	if err := dB.SeedDB(); err != nil {
		log.Error().Err(err).Msg("Error seeding db")
	}

	log.Info().Msg("Setting up routes...")

	r := chi.NewRouter()
	middlewares := []func(http.Handler) http.Handler{
		// hlog.NewHandler(log.Logger),
		// hlog.URLHandler("url"),
		// hlog.MethodHandler("method"),
	}
	r.Use(middlewares...)

	urlHandle := &handlers.URLHandle{DB: &dB}
	redirectHandle := &handlers.RedirectHandle{DB: &dB}
	usersHandle := &handlers.UserAccountHandle{DB: &dB}

	// Unauthenticated routes
	r.Get("/{shortURL}", redirectHandle.Redirect)

	r.Post("/shorten_url", urlHandle.CreateShortURL)

	r.Post("/register", usersHandle.Register)
	r.Post("/acct_login", usersHandle.Login)

	// Authenticated
	// users/{userID}/urls - get urls for user

	var handler http.Handler = r
	if cfg.IsDevEnv() {
		fmt.Println("setting cors")
		corsOptions := []ghandlers.CORSOption{
			ghandlers.AllowedMethods([]string{
				http.MethodGet,
				http.MethodPut,
				http.MethodHead,
				http.MethodPost,
				http.MethodDelete,
				http.MethodOptions,
			}),
			ghandlers.AllowedHeaders([]string{
				"Accept",
				"Content-Type",
			}),
			ghandlers.AllowedOrigins([]string{"http://localhost:3000"}),
		}

		handler = ghandlers.CORS(corsOptions...)(handler)
	}

	server := &http.Server{
		Addr: cfg.ListenAddress,
		// WriteTimeout:      time.Second * 600,
		// ReadHeaderTimeout: time.Second * 60,
		// ReadTimeout:       time.Second * 60,
		// IdleTimeout:       time.Second * 120,
		Handler: handler,
	}

	// Start the HTTP server.
	log.Info().Msgf("HTTP server listening on %v", cfg.ListenAddress)
	if err := server.ListenAndServe(); err != nil {
		log.Info().Err(err).Msg("Terminating server...")
	}
}
