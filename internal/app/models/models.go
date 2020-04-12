package models

import "fmt"

// ErrorType -
type ErrorType uint

// CtxKey - ctx key
type CtxKey string

const (
	// NoType - Unknown error
	NoType = ErrorType(iota)
	// BadRequest -
	BadRequest
	// NotFound -
	NotFound
	// AlreadyExists -
	AlreadyExists
	// Invalid -
	Invalid
	// Internal -
	Internal
)

// ReqIDKey -
const ReqIDKey = "rID"

// ErrorMessage -
type ErrorMessage struct {
	Target      string            `json:"target"`
	Message     string            `json:"message"`
	ContextInfo map[string]string `json:"internal"`
}

// Error - own realisation
type Error struct {
	Type        ErrorType
	Original    error
	Target      string
	Message     string
	ContextInfo map[string]string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Target, e.Message)
}

// User - пользователь
type User struct {
	Email     string `json:"email"     db:"email"`
	Nickname  string `json:"nickname"  db:"nickname"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname"  db:"lastname"`
}

// NewUser - новый пользователь
type NewUser struct {
	Email     string `json:"email"     db:"email"`
	Nickname  string `json:"nickname"  db:"nickname"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname"  db:"lastname"`
	Password  string `json:"password"  db:"password_digest"`
}
