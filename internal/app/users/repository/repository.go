package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	
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

func (repo *usersRepository) CreateUser(ctx context.Context, user models.User, password string) error {
	return nil
}
