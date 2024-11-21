package main

import (
	"dfs-transcoder/pkg/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/transcode", handlers.TranscodeHandler)

	log.Println("Worker Service is running on port 8081")
	log.Fatal(r.Run(":8081"))
}
