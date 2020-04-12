/*
 * Copyright (C) 2020. Leonid Moguchev
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
