package middlewares

import (
	"auth-service/pkg/app"
	"fmt"
	"net/http"
	"strings"

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

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		valid, userId := validateJWT(tokenString)
		if !valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		fmt.Println(r.RequestURI, r.Method)
		fmt.Println(userId)

		next.ServeHTTP(w, r)
	})
}
