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

func transcode(src, dst string) error {
	input, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer input.Close()

	output, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer output.Close()

	// Simulate transcoding (copying file as is for simplicity)
	_, err = io.Copy(output, input)
	return err
}

func TranscodeHandler(c *gin.Context) {
	filename := c.PostForm("filename")
	src := filepath.Join(storagePath, filename)
	dst := filepath.Join(storagePath, "transcoded_"+filename)

	if err := transcode(src, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to transcode file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File transcoded", "transcoded_file": dst})
}
