package handlers

import (
	"net/http"
	"strconv"
	"user-management/pkg/app"

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
