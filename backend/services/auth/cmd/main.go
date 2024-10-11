package main

import (
	"auth-service/pkg/db"
	"auth-service/pkg/handlers"
	"auth-service/pkg/middlewares"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting auth service on 127.0.0.1:8000")
	db.InitDB()

	router := mux.NewRouter()

	// Public Routes (no middleware required)
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	router.HandleFunc("/token/refresh", handlers.RefreshToken).Methods("POST")

	// Protected Routes (middleware for JWT validation)
	router.Handle("/profile", middlewares.JWTMiddleware(http.HandlerFunc(handlers.GetProfile))).Methods("GET")
	// router.Handle("/update", JWTMiddleware(http.HandlerFunc(UpdateUser))).Methods("PUT")
	// router.Handle("/delete", JWTMiddleware(http.HandlerFunc(DeleteUser))).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
