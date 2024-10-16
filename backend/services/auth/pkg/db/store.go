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

func (d *DB) DeleteUser(id int) error {
	_, err := d.db.Exec("delete from users where id=$1", id)
	return err
}

func (d *DB) GetUserDetails(id int) (*UserResponse, error) {
	rows, err := d.db.Query("select name, email, profile_pic from users where id=$1", id)
	if err != nil {
		return nil, err
	}

	var user UserResponse
	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}
	rows.Scan(&user.Name, &user.Email, &user.ProfilePic)
	return &user, nil
}

func (d *DB) UpdateUser(id int, userUpdateInfo *UserUpdateInfo) (*UserResponse, error) {
	var updatedUser UserResponse

	err := d.db.QueryRow(`
		UPDATE users
		SET name=$1, profile_pic=$2
		WHERE id=$3
		RETURNING name, email, profile_pic
	`, userUpdateInfo.Name, userUpdateInfo.ProfilePic, id).Scan(
		&updatedUser.Name,
		&updatedUser.Email,
		&updatedUser.ProfilePic,
	)

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

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

func (d *DB) List() ([]User, error) {
	rows, err := d.db.Query("SELECT id, name, email, profile_pic, created_at, password FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.ProfilePic, &user.CreatedAt, &user.Password)
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
