package usecase

import (
	"context"
	"time"

	"github.com/moguchev/UniBox/internal/app/models"
	"github.com/moguchev/UniBox/internal/app/users"
)

type userUsecase struct {
	usersRepo      users.Repository
	contextTimeout time.Duration
}

// NewUsersUsecase will create new an userUsecase object representation of user.Usecase interface
func NewUsersUsecase(repo users.Repository, timeout time.Duration) users.Usecase {
	return &userUsecase{
		usersRepo:      repo,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, user models.User, password string) error {
	return nil
}
