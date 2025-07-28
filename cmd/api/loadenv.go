package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func LoadEnv() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		viper.AutomaticEnv()
		log.Info().Msg("ENV loaded from AutomaticEnv()")
	} else {
		log.Info().Msg("ENV loaded from .env")
	}
}
