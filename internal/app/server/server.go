package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/moguchev/UniBox/internal/pkg/config"
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
	uUsecase := _usersUcase.NewUsersUsecase(uRepo, config.ContextTimeout)

	_usersHttpDeliver.NewUsersHandler(router, uUsecase)

	return router, nil
}

// RunServer - запуск сервера
func RunServer() {
	router, err := NewRouter()
	if err != nil {
		log.Fatal("Failed to create router")
	}
	addr := ":" + strconv.Itoa(config.MainAppPort)
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: config.MainAppWriteTimeout,
		ReadTimeout:  config.MainAppReadTimeout,
	}
	log.Infof("Server started at %s", addr)
	log.Fatal(srv.ListenAndServe())
}
