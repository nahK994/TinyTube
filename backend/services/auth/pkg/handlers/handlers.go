package handlers

import (
	"auth-service/pkg/db"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	userRepo db.UserRepository
}

func GetHandler(userRepo db.UserRepository) *Handler {
	return &Handler{userRepo: userRepo}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user db.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := hashPassword(user.Password)
	user.Password = hashedPassword
	if err := h.userRepo.Register(&user); err != nil {
		fmt.Println(err)
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

func (h *Handler) UserList(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepo.List()
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.userRepo.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var userInfo db.UserUpdateInfo
	json.NewDecoder(r.Body).Decode(&userInfo)

	user, err := h.userRepo.UpdateUser(id, &userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Login user and return tokens
func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Email    string
		Password string
	}

	json.NewDecoder(r.Body).Decode(&reqBody)
	user, err := h.userRepo.GetUserByEmail(reqBody.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !checkPasswordHash(reqBody.Password, user.Password) {
		http.Error(w, "email or password mismatch", http.StatusBadRequest)
		return
	}

	accessToken, err1 := generateJWT(user.ID)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	refreshToken, err2 := generateRefreshToken(user.ID)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// // Refresh access token using refresh token
// func RefreshToken(w http.ResponseWriter, r *http.Request) {
// 	var tokens map[string]string
// 	json.NewDecoder(r.Body).Decode(&tokens)

// 	refreshToken := tokens["refresh_token"]
// 	valid, email := utils.ValidateJWT(refreshToken)
// 	if !valid {
// 		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
// 		return
// 	}

// 	var dbRefreshToken string
// 	err := db.Instance.QueryRow("SELECT refresh_token FROM users WHERE email = $1", email).Scan(&dbRefreshToken)
// 	if err != nil || dbRefreshToken != refreshToken {
// 		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
// 		return
// 	}

// 	accessToken, _ := utils.GenerateJWT(email)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"access_token": accessToken,
// 	})
// }

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err1 := h.userRepo.GetUserDetails(id)

	w.Header().Set("Content-Type", "application/json")
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
