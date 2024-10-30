package security

import (
	"context"
	"net/http"
	"strings"
	"user-management/pkg/app"

	"github.com/golang-jwt/jwt/v5"
)

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return app.GetConfig().App.JWT_secret_key, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userId := int(claims["sub"].(float64))
			ctx := context.WithValue(r.Context(), "userId", userId)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
	})
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type middlewareFunc func(http.Handler) http.Handler

		// NOTE: add additional middlewares here
		middlewares := []middlewareFunc{
			authenticate,
		}

		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		next.ServeHTTP(w, r)
	})
}
