package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *zerolog.Logger
}

func main() {

	var cfg config

	LoadEnv()
	LoggerConfig()

	cfg.port = viper.GetInt("SERVER_PORT")
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	app := &application{
		config: cfg,
		logger: &logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  viper.GetDuration("SERVER_IDLE_TIMEOUT"),
		ReadTimeout:  viper.GetDuration("SERVER_READ_TIMEOUT"),
		WriteTimeout: viper.GetDuration("SERVER_WRITE_TIMEOUT"),
	}

	log.Info().Msgf("Starting server on port %d", cfg.port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Error().Msgf("Error starting server: %v", err)
	}
}
