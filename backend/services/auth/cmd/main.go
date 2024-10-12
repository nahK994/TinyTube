package main

import (
	"auth-service/internal/db"
	"auth-service/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := db.InitDB("user", "password", "127.0.0.1", 5432, "auth_db")
	if err != nil {
		log.Fatal(err)
	}
	handler := handlers.GetHandler(db)

	router := mux.NewRouter()
	// Public Routes (no middleware required)
	router.HandleFunc("/register", handler.RegisterUser).Methods("POST")
	router.HandleFunc("/", handler.UserList).Methods("GET")

	fmt.Println("Starting auth service on 127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
