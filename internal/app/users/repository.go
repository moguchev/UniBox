package users

import (
	"github.com/google/uuid"
)

// Repository -
type Repository interface {
	CreateUser() bool
}
