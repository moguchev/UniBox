package models

import "fmt"

const (
	// NoType - Unknown error
	NoType = ErrorType(iota)
	// BadRequest -
	BadRequest
	// NotFound -
	NotFound
	// AlreadyExists -
	AlreadyExists
	// Internal -
	Internal
)

// ReqIDKey -
const ReqIDKey = "rID"

// ErrorType -
type ErrorType uint

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

// CtxKey - ctx key
type CtxKey string

// User - пользователь
type User struct {
	Email     string `json:"email"     db:"email"`
	Nickname  string `json:"nickname"  db:"nickname"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname"  db:"lastname"`
	Password  string `json:"password,omitempty"`
}
