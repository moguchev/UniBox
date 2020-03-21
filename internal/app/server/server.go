package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/moguchev/UniBox/internal/app/server/handler"
	"github.com/moguchev/UniBox/internal/pkg/config"
	"github.com/moguchev/UniBox/internal/pkg/middleware"

	log "github.com/sirupsen/logrus"
)

// NewRouter - returns router
func NewRouter() (*mux.Router, error) {
	router := mux.NewRouter()
	router = router.PathPrefix("/api/").Subrouter()

	h := handler.Handler{}
	mw := middleware.InitMiddleware()

	router.Use(mw.RequestIDMiddleware)
	router.Use(mw.AccessLogMiddleware)
	router.Use(mw.CORSMiddleware)
	router.Use(mw.RecoverMiddleware)

	router.HandleFunc("/user", h.CreateUser).Methods(http.MethodPost, http.MethodOptions)

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
