package db

import (
	"fmt"
)

type Repository interface {
	Register(user *UserRequest) (*UserResponse, error)
	List() ([]UserResponse, error)
	DeleteUser(id int) error
	GetUserDetails(id int) (*UserResponse, error)
	UpdateUser(id int, userUpdateInfo *UserUpdateInfo) (*UserUpdateInfo, error)
}

func (d *DB) Register(userRequest *UserRequest) (*UserResponse, error) {
	var userResponse UserResponse
	err := d.db.QueryRow(`
	INSERT INTO users (name, email, profile_pic) 
	VALUES ($1, $2, $3, $4)
	RETURNING id, name, email, profile_pic, created_at`,
		userRequest.Name, userRequest.Email, userRequest.ProfilePic).Scan(
		&userResponse.ID,
		&userResponse.Name,
		&userResponse.Email,
		&userResponse.ProfilePic,
		&userResponse.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func (d *DB) DeleteUser(id int) error {
	_, err := d.db.Exec("delete from users where id=$1", id)
	return err
}

func (d *DB) GetUserDetails(id int) (*UserResponse, error) {
	rows, err := d.db.Query("select name, email, profile_pic, created_at from users where id=$1", id)
	if err != nil {
		return nil, err
	}

	var user UserResponse
	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}
	rows.Scan(&user.Name, &user.Email, &user.ProfilePic, &user.CreatedAt)
	return &user, nil
}

func (d *DB) UpdateUser(id int, userUpdateInfo *UserUpdateInfo) (*UserUpdateInfo, error) {
	var updatedUser UserUpdateInfo

	err := d.db.QueryRow(`
		UPDATE users
		SET name=$1, profile_pic=$2
		WHERE id=$3
		RETURNING name, email, profile_pic
	`, userUpdateInfo.Name, userUpdateInfo.ProfilePic, id).Scan(
		&updatedUser.Name,
		&updatedUser.ProfilePic,
	)

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (d *DB) List() ([]UserResponse, error) {
	rows, err := d.db.Query("SELECT id, name, email, profile_pic, created_at FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []UserResponse
	for rows.Next() {
		var user UserResponse
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
