package main

import (
	"fmt"
	"log"
	"video-upload/pkg/app"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := app.GetConfig()
	// db, err := db.InitDB(conf.DB)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// handler := handlers.GetHandler(db)

	r := gin.Default()

	srvAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	fmt.Println("Starting auth service on", srvAddress)
	log.Fatal(r.Run(srvAddress))
}
