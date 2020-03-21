package users

import (
	"github.com/google/uuid"
)

// Service - business logic
type Service interface {
	CreateUser() (uuid.UUID, error)
}
