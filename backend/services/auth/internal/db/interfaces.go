package db

type UserRepository interface {
	Register(user *User) error
	List() ([]User, error)
}
