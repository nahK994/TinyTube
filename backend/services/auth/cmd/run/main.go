package main

import (
	"auth-service/pkg/app"
	"auth-service/pkg/db"
	"auth-service/pkg/handlers"
	"auth-service/pkg/security"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf := app.GetConfig()
	srvAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	db, err := db.InitDB(conf.Database)
	if err != nil {
		log.Fatal(err)
	}
	handler := handlers.GetHandler(db)

	r := mux.NewRouter()

	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.Use(security.Middleware)
	userRouter.HandleFunc("/{id}", handler.DeleteUser).Methods(http.MethodDelete)
	userRouter.HandleFunc("/{id}", handler.GetProfile).Methods(http.MethodGet)

	r.HandleFunc("/register", handler.RegisterUser).Methods(http.MethodPost)
	r.HandleFunc("/users", handler.UserList).Methods(http.MethodGet)
	r.HandleFunc("/login", handler.LoginUser).Methods(http.MethodPost)

	fmt.Println("Starting auth service on", srvAddress)
	log.Fatal(http.ListenAndServe(srvAddress, r))
}
