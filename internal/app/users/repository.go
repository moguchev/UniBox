package users

import (
	"context"

	"github.com/moguchev/UniBox/internal/app/models"
)

// Repository - database level
type Repository interface {
	CreateUser(ctx context.Context, user models.NewUser) (models.User, error)
}
