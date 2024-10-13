package main

import (
	"fmt"
	"log"
	"net/http"
	"user-management/pkg/app"

	"github.com/gorilla/mux"
)

func main() {
	conf := app.GetConfig()
	router := mux.NewRouter()

	srvAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	fmt.Println("Starting auth service on", srvAddress)
	log.Fatal(http.ListenAndServe(srvAddress, router))
}
