package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/moguchev/UniBox/internal/app/server"
)

func init() {
	viper.SetConfigFile("../configs/config.json")
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
	serverOptions := server.MainServerOptions{}
	err := viper.UnmarshalKey("main_server", &serverOptions)
	if err != nil {
		log.Fatal(err)
	}

	dbOptions := server.MainDBOptions{}
	viper.UnmarshalKey("main_database", &dbOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctxOptions := server.ContextOptions{}
	viper.UnmarshalKey("context", &ctxOptions)
	if err != nil {
		log.Fatal(err)
	}

	server.RunServer(server.Options{
		Server:   serverOptions,
		Database: dbOptions,
		Context:  ctxOptions,
	})
}
