package handlers

import "user-management/pkg/db"

type Handler struct {
	userRepo db.UserRepository
}

func GetHandler(userRepo db.UserRepository) *Handler {
	return &Handler{userRepo: userRepo}
}
