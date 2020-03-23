package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/logrusadapter"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/moguchev/UniBox/internal/app/users"
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
	Port       uint16        `mapstructure:"port"`
	User       string        `mapstructure:"user"`
	Pass       string        `mapstructure:"pass"`
	Name       string        `mapstructure:"name"`
	MaxConn    int           `mapstructure:"max_connections"`
	AcqTimeout time.Duration `mapstructure:"acquire_timeout"`
}

type usecases struct {
	users users.Usecase
}

// ConnectToDB - возвращает коннект к БД
func ConnectToDB(opt MainDBOptions) (*sqlx.DB, error) {
	config := pgx.ConnConfig{
		Host:     opt.Host,
		Port:     opt.Port,
		Database: opt.Name,
		User:     opt.User,
		Password: opt.Pass,
		Logger:   logrusadapter.NewLogger(&log.Logger{}),
		LogLevel: pgx.LogLevelInfo,
		RuntimeParams: map[string]string{
			"standard_conforming_strings": "on",
		},
		PreferSimpleProtocol: true, // протокол без prepare запроса
	}

	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: opt.MaxConn,
		AcquireTimeout: opt.AcqTimeout * time.Second,
	})

	if err != nil {
		connPool.Close()
		return nil, err
	}

	nativeDB := stdlib.OpenDBFromPool(connPool)
	return sqlx.NewDb(nativeDB, "pgx"), nil
}

// NewRouter - returns gorilla mux router
func NewRouter(ucs usecases) *mux.Router {
	router := mux.NewRouter()
	router = router.PathPrefix("/api/").Subrouter()

	mw := middleware.InitMiddleware()
	router.Use(mw.RequestIDMiddleware)
	router.Use(mw.AccessLogMiddleware)
	router.Use(mw.CORSMiddleware)
	router.Use(mw.RecoverMiddleware)

	_usersHttpDeliver.NewUsersHandler(router, ucs.users)

	return router
}

// RunServer - запуск сервера
func RunServer(opt Options) {
	conn, err := ConnectToDB(opt.Database)
	if err != nil {
		log.Fatal(err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}

	uRepo := _usersRepo.NewUsersRepository(conn)
	uCase := _usersUcase.NewUsersUsecase(uRepo, opt.Context.Timeout*time.Second)

	ucs := usecases{
		users: uCase,
	}

	router := NewRouter(ucs)

	srv := &http.Server{
		Addr:         opt.Server.Address,
		Handler:      router,
		ReadTimeout:  opt.Server.ReadTimeout * time.Second,
		WriteTimeout: opt.Server.WriteTimeout * time.Second,
	}

	log.Infof("Server started at %s", opt.Server.Address)
	log.Fatal(srv.ListenAndServe())
}
