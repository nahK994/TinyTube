package db

import (
	"fmt"
)

func (d *DB) Register(user *User) error {
	db := d.db
	_, err := db.Exec(`INSERT INTO users (name, email, password, profile_pic) 
	VALUES ($1, $2, $3, $4)`, user.Name, user.Email, user.Password, user.ProfilePic)
	if err != nil {
		return err
	}

	return db.QueryRow("select id, created_at from users where email=$1", user.Email).Scan(&user.ID, &user.CreatedAt)
}

func (d *DB) List() ([]User, error) {
	rows, err := d.db.Query("SELECT id, name, email, profile_pic, created_at FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.ProfilePic, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return users, nil
}
