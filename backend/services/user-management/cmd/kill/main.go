package main

import (
	"fmt"
	"os/exec"
	"user-management/pkg/app"
)

func main() {
	port := app.GetConfig().App.Port
	cmdStr := fmt.Sprintf("sudo kill -9 $(sudo lsof -t -i:%d)", port)

	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Process on port", port, "was killed successfully.")
	}
}
