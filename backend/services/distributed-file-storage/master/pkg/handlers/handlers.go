package handlers

import (
	"dfs-master/pkg/app"
	"dfs-master/pkg/db"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse file"})
		return
	}

	fileID := db.SaveFileMetadata(header.Filename)

	config := app.GetConfig()
	worker := config.Workers[fileID%len(config.Workers)]

	fileContent, _ := header.Open()
	defer fileContent.Close()
	resp, err := http.Post(fmt.Sprintf("%s/store", worker), c.ContentType(), fileContent)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store file on worker"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File uploaded to %s", worker), "fileID": fileID})
}

func ReplicateHandler(c *gin.Context) {
	var req struct {
		FileID int `json:"file_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	config := app.GetConfig()
	worker := config.Workers[req.FileID%len(config.Workers)]
	fileMetadata := db.GetFileMetadata(req.FileID)

	data := url.Values{}
	data.Set("filename", fileMetadata)

	resp, err := http.PostForm(fmt.Sprintf("%s/replicate", worker), data)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to replicate file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File replicated to %s", worker)})
}

func RetrieveHandler(c *gin.Context) {
	fileID, _ := strconv.Atoi(c.Param("fileID"))
	config := app.GetConfig()
	worker := config.Workers[fileID%len(config.Workers)]
	resp, err := http.Get(fmt.Sprintf("%s/retrieve/%d", worker, fileID))
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve file"})
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	io.Copy(c.Writer, resp.Body)
	resp.Body.Close()
}
