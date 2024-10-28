package db

import "fmt"

type Repository interface {
	GetUserByEmail(email string) (*UserDetails, error)
	CreateUser(userRequest *UserCreate) error
	UpdatePassword(info *PasswordUpdate) error
	DeleteUser(id int) error
}

func (d *DB) GetUserByEmail(email string) (*UserDetails, error) {
	rows, err := d.db.Query("SELECT id, password FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}

	var userInfo UserDetails
	if !rows.Next() {
		return nil, fmt.Errorf("user email not found")
	}
	rows.Scan(&userInfo.ID, &userInfo.Password)
	return &userInfo, nil
}

func (d *DB) CreateUser(userRequest *UserCreate) error {
	_, err := d.db.Exec(`
	INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`,
		userRequest.ID, userRequest.Email, userRequest.Password)

	if err != nil {
		return err
	}

	return nil
}

func (d *DB) UpdatePassword(info *PasswordUpdate) error {
	_, err := d.db.Exec(`UPDATE users SET password=$1 WHERE email=$2`, info.Password, info.Email)
	return err
}

func (d *DB) DeleteUser(id int) error {
	_, err := d.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
