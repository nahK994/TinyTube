package db

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	ProfilePic string `json:"profile_pic"`
	CreatedAt  string `json:"created_at"`
}

type UserResponse struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	ProfilePic string `json:"profile_pic"`
	CreatedAt  string `json:"created_at"`
}
