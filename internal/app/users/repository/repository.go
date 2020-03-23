package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/moguchev/UniBox/internal/app/models"
	"github.com/moguchev/UniBox/internal/app/users"
)

type usersRepository struct {
	Conn *sqlx.DB
}

// NewUsersRepository will create an object that represent the users.Repository interface
func NewUsersRepository(Conn *sqlx.DB) users.Repository {
	return &usersRepository{Conn}
}

func (repo *usersRepository) CreateUser(ctx context.Context, user models.User) error {
	rID := ctx.Value(models.CtxKey("rID"))
	query := `INSERT INTO Users (nickname, password_digest, email, firstname, lastname)
				VALUES ($1, $2, $3, $4, $5);`
	res, err := repo.Conn.ExecContext(ctx, query,
		user.Nickname, user.Password, user.Email, user.Firstname, user.Lastname)
	if err != nil {
		log.WithFields(log.Fields{
			"request_id": rID,
		}).Debug(err)
		if e, ok := err.(pgx.PgError); ok {
			switch e.Code {
			case codes.ForeignKeyViolation, codes.NotNullViolation:
				if e.ConstraintName == "thread_author_fkey" {
					err = errors.New(messages.UserNotFound)
				} else {
					err = errors.New(messages.ForumNotFound)
				}
				break
			case codes.UniqueViolation:
				err = errors.New(messages.ThreadAlreadyExists)
				break
			default:
				log.Println(e.Code)
				err = errors.New(messages.ForumNotFound)
			}
		} else {
			err = errors.New(messages.ForumNotFound)
		}
	}

	return err
}
