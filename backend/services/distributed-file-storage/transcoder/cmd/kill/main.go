package main

import (
	"dfs-transcoder/pkg/app"
	"fmt"
	"os/exec"
)

func main() {
	port := app.GetConfig().Port
	cmdStr := fmt.Sprintf("sudo kill -9 $(sudo lsof -t -i:%d)", port)

	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Process on port", port, "was killed successfully.")
	}
}
