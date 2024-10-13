package main

import (
	"auth-service/pkg/app"
	"auth-service/pkg/db"
	"auth-service/pkg/handlers"
	"auth-service/pkg/middlewares"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf := app.GetConfig()
	db, err := db.InitDB(conf.Database)
	if err != nil {
		log.Fatal(err)
	}
	handler := handlers.GetHandler(db)

	router := mux.NewRouter()
	router.Handle("/users/{id}", middlewares.JWTMiddleware(http.HandlerFunc(handler.DeleteUser))).Methods(http.MethodDelete)
	router.Handle("/users/{id}", middlewares.JWTMiddleware(http.HandlerFunc(handler.GetProfile))).Methods(http.MethodGet)

	router.HandleFunc("/register", handler.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/users", handler.UserList).Methods(http.MethodGet)
	router.HandleFunc("/login", handler.LoginUser).Methods(http.MethodPost)

	srvAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	fmt.Println("Starting auth service on", srvAddress)
	log.Fatal(http.ListenAndServe(srvAddress, router))
}
