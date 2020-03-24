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

// ErrorType -
type ErrorType uint

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
	Email     string `json:"email"     db:"email"     valid:"email,required"`
	Nickname  string `json:"nickname"  db:"nickname"  valid:"matches(^[A-Za-z0-9_.]*$),required"`
	Firstname string `json:"firstname" db:"firstname" valid:"utfletter,runelength(1|50),required"`
	Lastname  string `json:"lastname"  db:"lastname"  valid:"utfletter,runelength(1|50),required"`
	Password  string `json:"password,omitempty" valid:"-"`
}
