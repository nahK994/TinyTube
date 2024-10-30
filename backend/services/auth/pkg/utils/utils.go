package utils

import (
	"auth-service/pkg/app"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, hashedPassword *string) error {
	cost := app.GetConfig().App.Bcrypt_password_cost
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}

	*hashedPassword = string(bytes)
	return nil
}
