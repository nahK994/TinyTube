package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-management/pkg/db"
	"user-management/pkg/mq"
)

type Handler struct {
	userRepo db.Repository
	msg      mq.MessageProcessor
}

func GetHandler(userRepo db.Repository, msg mq.MessageProcessor) *Handler {
	return &Handler{userRepo: userRepo, msg: msg}
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

	//ToDo: Need to complete it
	h.msg.PublishMessage("")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}
