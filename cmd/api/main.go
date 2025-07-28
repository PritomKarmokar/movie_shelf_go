package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"movie_shelf_go/cmd/config"
	"net/http"
)

func main() {

	config.LoadEnv()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", viper.GetInt("SERVER_PORT")),
		Handler:      mux,
		IdleTimeout:  viper.GetDuration("SERVER_IDLE_TIMEOUT"),
		ReadTimeout:  viper.GetDuration("SERVER_READ_TIMEOUT"),
		WriteTimeout: viper.GetDuration("SERVER_WRITE_TIMEOUT"),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}
