package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const storagePath = "./storage/"

func StoreHandler(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, "Failed to parse file")
		return
	}

	filename := filepath.Join(storagePath, header.Filename)
	if err := c.SaveUploadedFile(header, filename); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to save file")
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("file stored %s\n", header.Filename))
}

func ReplicateHandler(c *gin.Context) {
	filename := c.PostForm("filename")
	src := filepath.Join(storagePath, filename)
	dst := filepath.Join(storagePath, "replica_"+filename)

	input, err := os.Open(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to open source file")
		return
	}
	defer input.Close()

	output, err := os.Create(dst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to create replica")
		return
	}
	defer output.Close()

	io.Copy(output, input)
	c.JSON(http.StatusOK, "File replicated")
}

func RetrieveHandler(c *gin.Context) {
	filename := c.Param("filename")
	filepath := filepath.Join(storagePath, filename)

	file, err := os.Open(filepath)
	if err != nil {
		c.JSON(http.StatusNotFound, "File not found")
		return
	}
	defer file.Close()

	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	io.Copy(c.Writer, file)
}
