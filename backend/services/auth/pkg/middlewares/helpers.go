package middlewares

import (
	"auth-service/pkg/app"

	"github.com/golang-jwt/jwt/v5"
)

func validateJWT(tokenString string) (bool, int) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return app.GetConfig().App.JWT_secret_key, nil
	})

	if err != nil || !token.Valid {
		return false, -1
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if id, ok := claims["userId"].(float64); ok {
			return true, int(id)
		}
	}

	return false, -1
}
