package main

import (
	"fmt"
	"os/exec"
)

func main() {
	targetDir := "/home/vk"
	ls_cmd := exec.Command("ls", "-lah")
	ls_cmd.Dir = targetDir
	output, err := ls_cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running ls with args %s", output)
		return
	}
	fmt.Printf(string(output))
}
