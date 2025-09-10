package main

import (
	"fmt"
	"os/exec"
)

func main() {
	
	cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", "iwr https://get.activated.win | iex")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}

	fmt.Println(string(output))
}
