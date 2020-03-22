package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/moguchev/UniBox/internal/app/server"
)

func init() {
	viper.SetConfigFile("configs/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	if viper.GetBool("debug") {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
		log.Info("Service RUN on DEBUG mode")
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	server.RunServer()
}
