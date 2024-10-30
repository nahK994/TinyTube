package main

import (
	"auth-service/pkg/app"
	"auth-service/pkg/db"
	"auth-service/pkg/handlers"
	"auth-service/pkg/mq"
	"auth-service/pkg/security"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf := app.GetConfig()
	db, err := db.InitDB(conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.GetHandler(db)

	mq, err := mq.InitMQ(conf.MQ, db)
	if err != nil {
		mq.Close()
		log.Fatal(err)
	}
	err = mq.Start()
	if err != nil {
		mq.Close()
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/login", handler.LoginUser).Methods(http.MethodPost)

	secureRoute := r.PathPrefix("/change-password").Subrouter()
	secureRoute.Use(security.Middleware)
	secureRoute.HandleFunc("", handler.ChangePassword).Methods(http.MethodPut)

	srvAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	fmt.Println("Starting auth service on", srvAddress)
	log.Fatal(http.ListenAndServe(srvAddress, r))
}
