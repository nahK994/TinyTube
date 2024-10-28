package db

import "time"

type AuthTable struct {
	UserId       int       `json:"userId"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	IssueTime    time.Time `json:"issue_time"`
}

type UserCreate struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserDetails struct {
	ID       int    `json:"id"`
	Password string `json:"password,omitempty"`
}

type PasswordUpdate struct {
	Id       int    `json:"email"`
	Password string `json:"password,omitempty"`
}
