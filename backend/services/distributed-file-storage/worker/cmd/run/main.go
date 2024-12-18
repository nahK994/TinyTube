package main

import (
	"dfs-worker/pkg/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/store", handlers.StoreHandler)
	r.POST("/replicate", handlers.ReplicateHandler)
	r.GET("/retrieve/:filename", handlers.RetrieveHandler)

	log.Println("Worker Service is running on port 8081")
	log.Fatal(r.Run(":8081"))
}
