package db

type User struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	ProfilePic string `json:"profile_pic"`
}

type UserResponse struct {
	ID int `json:"id"`
	User
	CreatedAt string `json:"created_at"`
}

type UserUpdateRequest struct {
	Name       string `json:"name"`
	ProfilePic string `json:"profilePic"`
}
