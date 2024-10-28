package db

import (
	"fmt"
)

type Repository interface {
	Register(user *UserRequest) (*UserResponse, error)
	DeleteUser(id int) error
	GetUserDetails(id int) (*UserResponse, error)
	UpdateUser(id int, userUpdateInfo *UserUpdateInfo) (*UserUpdateInfo, error)
}

func (d *DB) Register(userRequest *UserRequest) (*UserResponse, error) {
	var userResponse UserResponse
	err := d.db.QueryRow(`
	INSERT INTO users (name, email, profile_pic) 
	VALUES ($1, $2, $3)
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
	rows, err := d.db.Query("select id, name, email, profile_pic, created_at from users where id=$1", id)
	if err != nil {
		return nil, err
	}

	var user UserResponse
	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}
	rows.Scan(&user.ID, &user.Name, &user.Email, &user.ProfilePic, &user.CreatedAt)
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
