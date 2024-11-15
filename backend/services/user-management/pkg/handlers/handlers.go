package handlers

import (
	"net/http"
	"user-management/pkg/db"
	"user-management/pkg/mq"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo db.Repository
	msg  mq.MessageProcessor
}

func GetHandler(userRepo db.Repository, msg mq.MessageProcessor) *Handler {
	return &Handler{repo: userRepo, msg: msg}
}

// HandleUserActions handles actions like GET, PUT, DELETE for a user.
func (h *Handler) HandleUserActions(c *gin.Context) {
	userId, ok := c.Get("userId") // Get user ID from context (set by middleware).
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing user ID"})
		return
	}

	id := userId.(int) // Ensure the type is correct.

	switch c.Request.Method {
	case http.MethodDelete:
		h.delete(c, id)
	case http.MethodPut:
		h.update(c, id)
	case http.MethodGet:
		h.get(c, id)
	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}

func (h *Handler) delete(c *gin.Context, id int) {
	if err := h.repo.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	h.msg.PublishMessage(mq.MessageAction{
		ActionType: mq.UserDelete,
		Message:    id,
	})

	c.Status(http.StatusNoContent)
}

func (h *Handler) update(c *gin.Context, id int) {
	var userInfo db.UserUpdateRequest
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := h.repo.UpdateUser(id, &userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) get(c *gin.Context, id int) {
	user, err := h.repo.GetUserDetails(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) RegisterUser(c *gin.Context) {
	var userRequest db.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := h.repo.Register(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}
