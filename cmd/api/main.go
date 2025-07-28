package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io"
	"movie_shelf_go/cmd/config"
	"net/http"
)

func main() {

	config.LoadEnv()
	config.LoggerConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	port := viper.GetInt("SERVER_PORT")
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		IdleTimeout:  viper.GetDuration("SERVER_IDLE_TIMEOUT"),
		ReadTimeout:  viper.GetDuration("SERVER_READ_TIMEOUT"),
		WriteTimeout: viper.GetDuration("SERVER_WRITE_TIMEOUT"),
	}

	log.Info().Msgf("Starting server on port %d", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Error().Msgf("Error starting server: %v", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}
