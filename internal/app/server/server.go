package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	_usersHttpDeliver "github.com/moguchev/UniBox/internal/app/users/delivery/http"
	_usersRepo "github.com/moguchev/UniBox/internal/app/users/repository"
	_usersUcase "github.com/moguchev/UniBox/internal/app/users/usecase"
	"github.com/moguchev/UniBox/internal/pkg/middleware"
)

// Options - опции для запуска сервера
type Options struct {
	Server   MainServerOptions
	Database MainDBOptions
	Context  ContextOptions
}

// ContextOptions -
type ContextOptions struct {
	Timeout time.Duration `mapstructure:"timeout"`
}

// MainServerOptions -
type MainServerOptions struct {
	Address      string        `mapstructure:"address"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

// MainDBOptions -
type MainDBOptions struct {
	Host       string        `mapstructure:"host"`
	Port       string        `mapstructure:"port"`
	User       string        `mapstructure:"user"`
	Pass       string        `mapstructure:"pass"`
	Name       string        `mapstructure:"name"`
	MaxConn    int           `mapstructure:"max_connections"`
	AcqTimeout time.Duration `mapstructure:"acquire_timeout"`
}

// ConnectToDB - возвращает коннект к БД
func ConnectToDB(opt MainDBOptions) (*sqlx.DB, error) {
	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		opt.User, opt.Pass, opt.Host, opt.Port, opt.Name)
	log.Debugf("DB URI: %s", uri)
	config, err := pgx.ParseURI(uri)
	if err != nil {
		return nil, err
	}

	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: opt.MaxConn,
		AcquireTimeout: opt.AcqTimeout * time.Second,
	})

	if err != nil {
		log.Fatalf("Failed to create connections pool: %s", err)
	}

	nativeDB := stdlib.OpenDBFromPool(connPool)
	return sqlx.NewDb(nativeDB, "pgx"), nil
}

// NewRouter - returns router
func NewRouter(opt Options) (*mux.Router, error) {
	router := mux.NewRouter()
	router = router.PathPrefix("/api/").Subrouter()

	mw := middleware.InitMiddleware()
	router.Use(mw.RequestIDMiddleware)
	router.Use(mw.AccessLogMiddleware)
	router.Use(mw.CORSMiddleware)
	router.Use(mw.RecoverMiddleware)

	conn, err := ConnectToDB(opt.Database)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	uRepo := _usersRepo.NewUsersRepository(conn)
	uCase := _usersUcase.NewUsersUsecase(uRepo, opt.Context.Timeout*time.Second)
	_usersHttpDeliver.NewUsersHandler(router, uCase)

	return router, nil
}

// RunServer - запуск сервера
func RunServer(opt Options) {
	router, err := NewRouter(opt)
	if err != nil {
		log.Fatalf("Failed to create router: %s", err)
	}

	srv := &http.Server{
		Addr:         opt.Server.Address,
		Handler:      router,
		ReadTimeout:  opt.Server.ReadTimeout * time.Second,
		WriteTimeout: opt.Server.WriteTimeout * time.Second,
	}
	log.Infof("Server started at %s", opt.Server.Address)
	log.Fatal(srv.ListenAndServe())
}
