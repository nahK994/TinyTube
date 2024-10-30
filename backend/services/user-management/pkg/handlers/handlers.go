package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"user-management/pkg/db"
	"user-management/pkg/mq"
)

type Handler struct {
	repo db.Repository
	msg  mq.MessageProcessor
}

func GetHandler(userRepo db.Repository, msg mq.MessageProcessor) *Handler {
	return &Handler{repo: userRepo, msg: msg}
}

func writeErrorResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *Handler) HandleUserActions(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		writeErrorResponse(w, http.StatusUnauthorized, "Invalid or missing user ID")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodDelete:
		h.delete(w, userId)
	case http.MethodPut:
		h.update(w, r.Body, userId)
	case http.MethodGet:
		h.get(w, userId)
	default:
		writeErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *Handler) delete(w http.ResponseWriter, id int) {
	if err := h.repo.DeleteUser(id); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}
	h.msg.PublishMessage(mq.MessageAction{
		ActionType: mq.UserDelete,
		Message:    id,
	})
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) update(w http.ResponseWriter, data io.ReadCloser, id int) {
	var userInfo db.UserUpdateRequest
	if err := json.NewDecoder(data).Decode(&userInfo); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.repo.UpdateUser(id, &userInfo)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) get(w http.ResponseWriter, id int) {
	user, err := h.repo.GetUserDetails(id)
	if err != nil {
		writeErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userRequest db.User
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.repo.Register(&userRequest)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to register user")
		return
	}

	h.msg.PublishMessage(mq.MessageAction{
		ActionType: mq.UserCreate,
		Message: mq.CreateMessage{
			Email:    user.Email,
			Id:       user.ID,
			Password: userRequest.Password,
		},
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}
