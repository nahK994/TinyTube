package db

import "fmt"

func (d *DB) GetUserByEmail(email string) (*User, error) {
	rows, err := d.db.Query("select id, name, email, profile_pic, password from users where email=$1", email)
	if err != nil {
		return nil, err
	}

	var user User
	if !rows.Next() {
		return nil, fmt.Errorf("user email not found")
	}
	rows.Scan(&user.ID, &user.Name, &user.Email, &user.ProfilePic, &user.Password)
	return &user, nil
}
