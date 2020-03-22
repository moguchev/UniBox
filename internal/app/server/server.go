package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/moguchev/UniBox/internal/pkg/middleware"

	_usersHttpDeliver "github.com/moguchev/UniBox/internal/app/users/delivery/http"
	_usersRepo "github.com/moguchev/UniBox/internal/app/users/repository"
	_usersUcase "github.com/moguchev/UniBox/internal/app/users/usecase"
)

// NewRouter - returns router
func NewRouter() (*mux.Router, error) {
	router := mux.NewRouter()
	router = router.PathPrefix("/api/").Subrouter()

	mw := middleware.InitMiddleware()
	router.Use(mw.RequestIDMiddleware)
	router.Use(mw.AccessLogMiddleware)
	router.Use(mw.CORSMiddleware)
	router.Use(mw.RecoverMiddleware)

	uRepo := _usersRepo.NewUsersRepository(nil)

	timeoutContext := time.Duration(viper.GetInt64("context.timeout")) * time.Second
	uUsecase := _usersUcase.NewUsersUsecase(uRepo, timeoutContext)

	_usersHttpDeliver.NewUsersHandler(router, uUsecase)

	return router, nil
}

// RunServer - запуск сервера
func RunServer() {
	router, err := NewRouter()
	if err != nil {
		log.Fatalf("Failed to create router: %s", err)
	}
	addr := viper.GetString("main_server.address")
	writeTimeout := viper.GetInt64("main_server.write_timeout")
	readTimeout := viper.GetInt64("main_server.read_timeout")

	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
	}
	log.Infof("Server started at %s", addr)
	log.Fatal(srv.ListenAndServe())
}
