package db

import "fmt"

type Repository interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(userRequest *User) error
	UpdatePassword(info *User) error
	DeleteUser(id int) error
}

func (d *DB) GetUserByEmail(email string) (*User, error) {
	rows, err := d.db.Query("SELECT id, email, password FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}

	var userInfo User
	if !rows.Next() {
		return nil, fmt.Errorf("user email not found")
	}
	rows.Scan(&userInfo.ID, &userInfo.Email, &userInfo.Password)
	return &userInfo, nil
}

func (d *DB) CreateUser(userRequest *User) error {
	_, err := d.db.Exec(`
	INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`,
		userRequest.ID, userRequest.Email, userRequest.Password)

	return err
}

func (d *DB) UpdatePassword(info *User) error {
	_, err := d.db.Exec(`UPDATE users SET password=$1 WHERE id=$2`, info.Password, info.ID)
	return err
}

func (d *DB) DeleteUser(id int) error {
	_, err := d.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
