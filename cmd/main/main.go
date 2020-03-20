package main

import (
	"github.com/moguchev/UniBox/internal/app/server"
	"github.com/moguchev/UniBox/internal/pkg/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	if config.Debug {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	server.RunServer()
}
