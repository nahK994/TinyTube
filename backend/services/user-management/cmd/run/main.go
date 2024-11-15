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

	"github.com/gin-gonic/gin"
)

func main() {
	conf := app.GetConfig()
	db, err := db.InitDB(conf.DB)
	if err != nil {
		log.Fatal(err)
	}
	mq, err := mq.InitMQ(conf.MQ)
	if err != nil {
		mq.Close()
		log.Fatal(err)
	}
	handler := handlers.GetHandler(db, mq)

	r := gin.Default()
	authRoutes := r.Group("/users", security.MiddlewareManager()) // Protect routes with auth middleware
	{
		authRoutes.GET("/:id", handler.HandleUserActions)
		authRoutes.PUT("/:id", handler.HandleUserActions)
		authRoutes.DELETE("/:id", handler.HandleUserActions)
	}

	r.POST("/register", handler.RegisterUser) // Public route
	srvAddress := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	fmt.Println("Starting user-management service on", srvAddress)
	log.Fatal(http.ListenAndServe(srvAddress, r))
}
