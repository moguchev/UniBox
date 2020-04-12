package repository

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/moguchev/UniBox/internal/app/models"
	"github.com/moguchev/UniBox/internal/app/users"
	psql "github.com/moguchev/UniBox/pkg/psql"
)

type usersRepository struct {
	Conn *sqlx.DB
}

// NewUsersRepository will create an object that represent the users.Repository interface
func NewUsersRepository(Conn *sqlx.DB) users.Repository {
	return &usersRepository{Conn}
}

func (repo *usersRepository) CreateUser(ctx context.Context, user models.NewUser) (models.User, error) {
	rID := ctx.Value(models.CtxKey(models.ReqIDKey))
	now := time.Now()
	query := `INSERT INTO Users (nickname, password_digest, email, firstname, lastname, registration_time)
				VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;`
	_, err := repo.Conn.ExecContext(ctx, query,
		user.Nickname, user.Password, user.Email, user.Firstname, user.Lastname, now)

	if err != nil {
		var target string
		errorType := models.NoType

		if e, ok := err.(pgx.PgError); ok {
			switch e.Code {
			case psql.UniqueViolation:
				errorType = models.AlreadyExists
				target = strings.Split(e.ConstraintName, "_")[1]
				if target == "pkey" {
					target = "nickname"
				}
				break
			case psql.CheckViolation:
				errorType = models.Invalid
				target = strings.Split(e.ConstraintName, "_")[0]
				break
			default:
				errorType = models.Internal
				psql.Print(e)
				break
			}
		} else {
			errorType = models.Internal
		}
		if errorType == models.Internal {
			log.WithFields(log.Fields{
				"request_id": rID,
				"place":      "repository",
				"action":     "exec",
				"query":      query,
				"args":       user,
			}).Error(err)
		}
		err = models.Error{
			Type:     errorType,
			Target:   target,
			Message:  err.Error(),
			Original: err,
		}
	}

	usr := models.User{
		Nickname:         user.Nickname,
		Email:            user.Email,
		Firstname:        user.Firstname,
		Lastname:         user.Lastname,
		RegistrationTime: now,
	}

	return usr, err
}
