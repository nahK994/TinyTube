package db

import "time"

type AuthTable struct {
	UserId       int       `json:"userId"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	IssueTime    time.Time `json:"issue_time"`
}

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	ProfilePic string `json:"profile_pic"`
	CreatedAt  string `json:"created_at"`
}
