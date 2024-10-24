package main

import (
	"fmt"
	"log"
	"net/http"
	"user-management/pkg/app"
	"user-management/pkg/db"
	"user-management/pkg/handlers"
	"user-management/pkg/mq"
	"user-management/pkg/security"

	"github.com/gorilla/mux"
)

func main() {
	conf := app.GetConfig()
	db, err := db.InitDB(conf.DB)
	if err != nil {
		log.Fatal(err)
	}
	mq, err := mq.InitMQ(conf.MQ)
	if err != nil {
		log.Fatal(err)
	}
	handler := handlers.GetHandler(db, mq)

	r := mux.NewRouter()
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.Use(security.Middleware)
	userRouter.HandleFunc("/{id}", handler.DeleteUser).Methods(http.MethodDelete)
	userRouter.HandleFunc("/{id}", handler.GetProfile).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", handler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/register", handler.RegisterUser).Methods(http.MethodPost)

	srvAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	fmt.Println("Starting auth service on", srvAddress)
	log.Fatal(http.ListenAndServe(srvAddress, r))
}
