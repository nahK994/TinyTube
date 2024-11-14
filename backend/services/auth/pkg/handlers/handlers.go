package handlers

import (
	"auth-service/pkg/db"
	"auth-service/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo db.Repository
}

func GetHandler(userRepo db.Repository) *Handler {
	return &Handler{repo: userRepo}
}

func (h *Handler) ChangePassword(c *gin.Context) {
	id, _ := c.Get("userId")

	var reqBody db.ChangePasswordRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var hashedPassword string
	err := utils.HashPassword(reqBody.Password, &hashedPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.repo.UpdatePassword(&db.User{
		ID:       id.(int),
		Password: hashedPassword,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "password updated")
}

func (h *Handler) LoginUser(c *gin.Context) {
	var reqBody struct {
		Email    string
		Password string
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.repo.GetUserByEmail(reqBody.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if !checkPasswordHash(reqBody.Password, user.Password) {
		c.JSON(http.StatusBadRequest, "email or password mismatch")
		return
	}

	accessToken, err1 := generateJWT(user.ID)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, err1.Error())
		return
	}
	refreshToken, err2 := generateRefreshToken(user.ID)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, err2.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
