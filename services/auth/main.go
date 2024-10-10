package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting auth service on 127.0.0.1:8000")
	InitDB()

	router := mux.NewRouter()

	// Public Routes (no middleware required)
	router.HandleFunc("/register", RegisterUser).Methods("POST")
	router.HandleFunc("/login", LoginUser).Methods("POST")
	router.HandleFunc("/token/refresh", RefreshToken).Methods("POST")

	// Protected Routes (middleware for JWT validation)
	router.Handle("/profile", JWTMiddleware(http.HandlerFunc(GetProfile))).Methods("GET")
	// router.Handle("/update", JWTMiddleware(http.HandlerFunc(UpdateUser))).Methods("PUT")
	// router.Handle("/delete", JWTMiddleware(http.HandlerFunc(DeleteUser))).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
