package db

type UserRepository interface {
	GetUserByEmail(email string) (*User, error)
}
