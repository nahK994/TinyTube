package middlewares

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

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

		vars := mux.Vars(r)
		resourceId, _ := strconv.Atoi(vars["id"])
		if userId != resourceId {
			http.Error(w, "Permission denied", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
