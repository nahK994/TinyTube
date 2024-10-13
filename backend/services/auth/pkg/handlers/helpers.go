package handlers

import (
	"auth-service/pkg/app"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func getId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	key := vars["id"]
	if id, err := strconv.Atoi(key); err != nil {
		return -1, err
	} else {
		return id, nil
	}
}

func hashPassword(password string) (string, error) {
	cost := app.GetConfig().App.Bcrypt_password_cost
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Duration(app.GetConfig().App.JWT_exp_minutes) * time.Minute).Unix(),
	})

	jwtSecretKey := app.GetConfig().App.JWT_secrey_key
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	return tokenString, err
}

func generateRefreshToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Duration(app.GetConfig().App.RefreshToken_exp_hours) * time.Hour).Unix(),
	})

	jwtSecretKey := app.GetConfig().App.JWT_secrey_key
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	return tokenString, err
}
