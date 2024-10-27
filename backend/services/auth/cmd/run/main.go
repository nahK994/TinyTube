package main

import (
	"auth-service/pkg/app"
	"auth-service/pkg/db"
	"auth-service/pkg/handlers"
	"auth-service/pkg/mq"
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

	mq, err := mq.InitMQ(conf.MQ)
	if err != nil {
		mq.Close()
		log.Fatal(err)
	}
	err = mq.Start(handler)
	if err != nil {
		mq.Close()
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/login", handler.LoginUser).Methods(http.MethodPost)

	srvAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	fmt.Println("Starting auth service on", srvAddress)
	log.Fatal(http.ListenAndServe(srvAddress, r))
}
