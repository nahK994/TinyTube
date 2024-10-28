package handlers

import (
	"encoding/json"
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
	info := mq.MessageAction{
		ActionType: mq.UserDelete,
		Message:    id,
	}
	h.msg.PublishMessage(info)
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

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userRequest db.UserRequest
	json.NewDecoder(r.Body).Decode(&userRequest)

	user, err := h.userRepo.Register(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	info := mq.MessageAction{
		ActionType: mq.UserCreate,
		Message: mq.CreateMessage{
			Email:    user.Email,
			Id:       user.ID,
			Password: userRequest.Password,
		},
	}
	h.msg.PublishMessage(info)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}
