package utils

import (
	"auth-service/pkg/app"

	"github.com/dgrijalva/jwt-go"
)

func ValidateJWT(tokenString string) (bool, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return app.GetConfig().App.JWT_secrey_key, nil
	})

	if err != nil {
		return false, ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims["email"].(string)
	}

	return false, ""
}
