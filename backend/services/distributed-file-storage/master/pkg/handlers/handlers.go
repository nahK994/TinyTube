package api

import (
	"dfs-master/pkg/db"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

var workerNodes = []string{
	"http://localhost:8081",
	"http://localhost:8082",
	"http://localhost:8083",
	"http://localhost:8084",
	"http://localhost:8085",
}

// UploadHandler handles video uploads and distributes to a worker
func UploadHandler(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse file"})
		return
	}

	// Save file metadata
	fileID := db.SaveFileMetadata(header.Filename)

	// Choose a worker node
	worker := workerNodes[fileID%len(workerNodes)]

	// Forward the file to the worker
	fileContent, _ := header.Open()
	defer fileContent.Close()
	resp, err := http.Post(fmt.Sprintf("%s/store", worker), c.ContentType(), fileContent)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store file on worker"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File uploaded to %s", worker), "fileID": fileID})
}

// ReplicateHandler replicates a file to another worker node
func ReplicateHandler(c *gin.Context) {
	var req struct {
		FileID int `json:"file_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Select the next worker for replication
	worker := workerNodes[(req.FileID+1)%len(workerNodes)]
	fileMetadata := db.GetFileMetadata(req.FileID)

	// Forward replication request to the worker
	data := url.Values{}
	data.Set("filename", fileMetadata)

	resp, err := http.PostForm(fmt.Sprintf("%s/replicate", worker), data)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to replicate file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File replicated to %s", worker)})
}

// RetrieveHandler retrieves a file from the appropriate worker
func RetrieveHandler(c *gin.Context) {
	fileID, _ := strconv.Atoi(c.Param("fileID"))
	worker := workerNodes[fileID%len(workerNodes)]
	resp, err := http.Get(fmt.Sprintf("%s/retrieve/%d", worker, fileID))
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve file"})
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	io.Copy(c.Writer, resp.Body)
	resp.Body.Close()
}
