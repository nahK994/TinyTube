package handlers

import (
	"auth-service/pkg/db"
	"encoding/json"
	"net/http"
)

type Handler struct {
	repo db.Repository
}

func GetHandler(userRepo db.Repository) *Handler {
	return &Handler{repo: userRepo}
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Email    string
		Password string
	}

	json.NewDecoder(r.Body).Decode(&reqBody)
	user, err := h.repo.GetUserByEmail(reqBody.Email)
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

func (h *Handler) CreateUser(userCreate db.UserCreate) error {
	hashedPassword, err := hashPassword(userCreate.Password)
	if err != nil {
		return err
	}

	user := db.UserCreate{
		ID:       userCreate.ID,
		Email:    userCreate.Email,
		Password: hashedPassword,
	}
	err = h.repo.CreateUser(&user)
	return err
}

func (h *Handler) UpdatePassword(userPassword db.PasswordUpdate) error {
	hashedPassword, err := hashPassword(userPassword.Password)
	if err != nil {
		return err
	}

	err = h.repo.UpdatePassword(&db.PasswordUpdate{
		Email:    userPassword.Email,
		Password: hashedPassword,
	})
	return err
}

func (h *Handler) DeleteUser(id int) error {
	err := h.repo.DeleteUser(id)
	return err
}
