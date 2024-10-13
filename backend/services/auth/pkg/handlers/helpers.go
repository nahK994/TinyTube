package handlers

import (
	"auth-service/pkg/app"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func generateJWT(id int) (string, error) {
	now := time.Now()
	expTime := now.Add(time.Duration(app.GetConfig().App.JWT_exp_minutes) * time.Minute)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    id,
		"iss":    "TinyTube",
		"exp":    expTime.Unix(),
		"iat":    now.Unix(),
		"userId": id,
	})

	tokenString, err := claims.SignedString(app.GetConfig().App.JWT_secret_key)
	return tokenString, err
}

func generateRefreshToken(id int) (string, error) {
	expTime := time.Now().Add(time.Duration(app.GetConfig().App.RefreshToken_exp_hours) * time.Hour).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": id,
		"iss":    "TinyTube",
		"exp":    expTime,
		"iat":    time.Now().Unix(),
	})

	jwtSecretKey := app.GetConfig().App.JWT_secret_key
	tokenString, err := claims.SignedString(jwtSecretKey)
	return tokenString, err
}
