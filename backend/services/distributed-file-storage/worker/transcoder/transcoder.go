package transcoder

import (
	"fmt"
	"io"
	"os"
)

// Transcode simulates video transcoding
func Transcode(src, dst string) error {
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
