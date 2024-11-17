package main

import (
	"dfs-worker/pkg/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Worker Service Routes
	r.POST("/store", handlers.StoreHandler)
	r.POST("/replicate", handlers.ReplicateHandler)
	r.GET("/retrieve/:filename", handlers.RetrieveHandler)
	r.POST("/transcode", handlers.TranscodeHandler)

	log.Println("Worker Service is running on port 8081")
	log.Fatal(r.Run(":8081"))
}
