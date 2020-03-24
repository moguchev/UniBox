package repository

import (
	"context"

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

func (repo *usersRepository) CreateUser(ctx context.Context, user models.User) error {
	rID := ctx.Value(models.CtxKey("rID"))
	query := `INSERT INTO Users (nickname, password_digest, email, firstname, lastname)
				VALUES ($1, $2, $3, $4, $5);`
	_, err := repo.Conn.ExecContext(ctx, query,
		user.Nickname, user.Password, user.Email, user.Firstname, user.Lastname)

	if err != nil {
		log.WithFields(log.Fields{"request_id": rID}).Debug(err)
		if e, ok := err.(pgx.PgError); ok {
			switch e.Code {
			case psql.UniqueViolation:
				err = models.Error{
					Type:     models.AlreadyExists,
					Target:   e.ConstraintName,
					Message:  "Already exists",
					Original: err,
				}
				log.Print(e.Code)
				log.Print(e.ColumnName)
				log.Print(e.ConstraintName)
				log.Print(e.DataTypeName)
				log.Print(e.Detail)
				log.Print(e.File)
				log.Print(e.Hint)
				log.Print(e.InternalPosition)
				log.Print(e.InternalQuery)
				log.Print(e.Line)
				log.Print(e.Message)
				log.Print(e.Position)
				log.Print(e.Routine)
				log.Print(e.SchemaName)
				log.Print(e.Severity)
				log.Print(e.TableName)
				log.Print(e.Where)
			default:
				log.WithFields(log.Fields{"request_id": rID}).Error(err)
				break
			}
		} else {
			log.WithFields(log.Fields{"request_id": rID}).Error(err)
			err = models.Error{
				Type:     models.Internal,
				Original: err,
			}
		}
	}

	return err
}
