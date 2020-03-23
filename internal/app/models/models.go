package models

// CtxKey - ctx key
type CtxKey string

// User - пользователь
type User struct {
	Email     string `json:"email"     db:"email"     valid:"email,required"`
	Nickname  string `json:"nickname"  db:"nickname"  valid:"matches(^[A-Za-z0-9_.]*$),required"`
	Firstname string `json:"firstname" db:"firstname" valid:"utfletter,runelength(1|50),required"`
	Lastname  string `json:"lastname"  db:"lastname"  valid:"utfletter,runelength(1|50),required"`
	Password  string `json:"password,omitempty" valid:"-"`
	// Password must contain at least one letter, at least one number, and be longer than six charaters.
}
