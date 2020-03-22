package users

// Repository - database level
type Repository interface {
	CreateUser() bool
}
