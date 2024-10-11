package handlers

import (
	"auth-service/pkg/db"
	"auth-service/pkg/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

// Register new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user db.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	err := db.Instance.QueryRow(`INSERT INTO users (name, email, password, profile_pic) 
                        VALUES ($1, $2, $3, $4) RETURNING id, created_at`,
		user.Name, user.Email, user.Password, user.ProfilePic).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	refreshToken, _ := utils.GenerateRefreshToken(user.Email)
	db.Instance.Exec("UPDATE users SET refresh_token = $1 WHERE email = $2", refreshToken, user.Email)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

// Login user and return tokens
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user db.User
	var dbPassword string

	json.NewDecoder(r.Body).Decode(&user)

	err := db.Instance.QueryRow("SELECT password FROM users WHERE email = $1", user.Email).Scan(&dbPassword)
	if err == sql.ErrNoRows || !utils.CheckPasswordHash(user.Password, dbPassword) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, _ := utils.GenerateJWT(user.Email)
	refreshToken, _ := utils.GenerateRefreshToken(user.Email)

	db.Instance.Exec("UPDATE users SET refresh_token = $1 WHERE email = $2", refreshToken, user.Email)

	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Refresh access token using refresh token
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var tokens map[string]string
	json.NewDecoder(r.Body).Decode(&tokens)

	refreshToken := tokens["refresh_token"]
	valid, email := utils.ValidateJWT(refreshToken)
	if !valid {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	var dbRefreshToken string
	err := db.Instance.QueryRow("SELECT refresh_token FROM users WHERE email = $1", email).Scan(&dbRefreshToken)
	if err != nil || dbRefreshToken != refreshToken {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	accessToken, _ := utils.GenerateJWT(email)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": accessToken,
	})
}

// Get user profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	// JWT token is validated in the middleware, now fetch user profile
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
	_, email := utils.ValidateJWT(token)

	var user db.User
	err := db.Instance.QueryRow("SELECT id, name, email, profile_pic, created_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.ProfilePic, &user.CreatedAt)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
