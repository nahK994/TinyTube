package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define Master Service routes
	// r.POST("/upload", api.UploadHandler)
	// r.POST("/replicate", api.ReplicateHandler)
	// r.GET("/retrieve/:fileID", api.RetrieveHandler)

	log.Println("Master Service running on port 8080")
	log.Fatal(r.Run(":8080"))
}
