package main

import (
	"auth-service/pkg/app"
	"auth-service/pkg/db"
	"auth-service/pkg/handlers"
	"auth-service/pkg/mq"
	"auth-service/pkg/security"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	r.POST("/login", handler.LoginUser)

	secureRoute := r.Group("/change-password")
	secureRoute.Use(security.AuthMiddleware())
	secureRoute.PUT("", handler.ChangePassword)

	srvAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	fmt.Println("Starting auth service on", srvAddress)
	log.Fatal(r.Run(srvAddress))
}
