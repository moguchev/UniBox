package users

import (
	"context"

	"github.com/moguchev/UniBox/internal/app/models"
)

// Usecase - business logic
type Usecase interface {
	CreateUser(ctx context.Context, user models.User) error
}
