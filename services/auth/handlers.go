package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

// Register new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := HashPassword(user.Password)
	user.Password = hashedPassword

	err := db.QueryRow(`INSERT INTO users (name, email, password, profile_pic) 
                        VALUES ($1, $2, $3, $4) RETURNING id, created_at`,
		user.Name, user.Email, user.Password, user.ProfilePic).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	refreshToken, _ := GenerateRefreshToken(user.Email)
	db.Exec("UPDATE users SET refresh_token = $1 WHERE email = $2", refreshToken, user.Email)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

// Login user and return tokens
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var dbPassword string

	json.NewDecoder(r.Body).Decode(&user)

	err := db.QueryRow("SELECT password FROM users WHERE email = $1", user.Email).Scan(&dbPassword)
	if err == sql.ErrNoRows || !CheckPasswordHash(user.Password, dbPassword) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, _ := GenerateJWT(user.Email)
	refreshToken, _ := GenerateRefreshToken(user.Email)

	db.Exec("UPDATE users SET refresh_token = $1 WHERE email = $2", refreshToken, user.Email)

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
	valid, email := ValidateJWT(refreshToken)
	if !valid {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	var dbRefreshToken string
	err := db.QueryRow("SELECT refresh_token FROM users WHERE email = $1", email).Scan(&dbRefreshToken)
	if err != nil || dbRefreshToken != refreshToken {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	accessToken, _ := GenerateJWT(email)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": accessToken,
	})
}

// Get user profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	// JWT token is validated in the middleware, now fetch user profile
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
	_, email := ValidateJWT(token)

	var user User
	err := db.QueryRow("SELECT id, name, email, profile_pic, created_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.ProfilePic, &user.CreatedAt)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
