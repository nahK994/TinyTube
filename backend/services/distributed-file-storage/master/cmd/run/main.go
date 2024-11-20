package main

import (
	"dfs-master/pkg/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define Master Service routes
	r.POST("/upload", handlers.UploadHandler)
	r.POST("/replicate", handlers.ReplicateHandler)
	r.GET("/retrieve/:fileID", handlers.RetrieveHandler)

	log.Println("Master Service running on port 8080")
	log.Fatal(r.Run(":8080"))
}
