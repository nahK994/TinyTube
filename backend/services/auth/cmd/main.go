package main

import (
	"auth-service/internal/config"
	"auth-service/internal/db"
	"auth-service/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	db, err := db.InitDB(conf.Database)
	if err != nil {
		log.Fatal(err)
	}
	handler := handlers.GetHandler(db)

	router := mux.NewRouter()
	router.HandleFunc("/register", handler.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", handler.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}", handler.GetProfile).Methods(http.MethodGet)
	router.HandleFunc("/users", handler.UserList).Methods(http.MethodGet)

	srvAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	fmt.Println("Starting auth service on", srvAddress)
	log.Fatal(http.ListenAndServe(srvAddress, router))
}
