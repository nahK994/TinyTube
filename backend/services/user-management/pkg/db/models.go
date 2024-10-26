package db

type UserRequest struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	ProfilePic string `json:"profile_pic"`
}

type User struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	ProfilePic string `json:"profile_pic"`
}

type UserResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	ProfilePic string `json:"profile_pic"`
	CreatedAt  string `json:"created_at"`
}

type UserUpdateInfo struct {
	Name       string `json:"name"`
	ProfilePic string `json:"profilePic"`
}
