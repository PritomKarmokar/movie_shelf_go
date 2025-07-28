package config

import (
	"github.com/spf13/viper"
	"log"
)

func LoadEnv() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		viper.AutomaticEnv()
		log.Println("ENV loaded from AutomaticEnv()")
	} else {
		log.Println("ENV loaded from .env")
	}
}
