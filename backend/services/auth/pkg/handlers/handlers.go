package handlers

import (
	"auth-service/pkg/db"
	"auth-service/pkg/utils"
	"encoding/json"
	"net/http"
)

type Handler struct {
	repo db.Repository
}

func GetHandler(userRepo db.Repository) *Handler {
	return &Handler{repo: userRepo}
}

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var reqBody db.ChangeUpdateRequest
	json.NewDecoder(r.Body).Decode(&reqBody)

	var hashedPassword string
	err := utils.HashPassword(reqBody.Password, &hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.repo.UpdatePassword(&db.ChangeUpdateRequest{
		Id:       reqBody.Id,
		Password: hashedPassword,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("password updated")
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
