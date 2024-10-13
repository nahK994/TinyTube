package db

type UserRepository interface {
	Register(user *User) error
	List() ([]User, error)
	DeleteUser(id int) error
	GetUserDetails(id int) (*UserResponse, error)
}
