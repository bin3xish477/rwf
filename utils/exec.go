package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func getCmdPath(cmd string) (string, bool) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	return path, true
}

func ExecuteCmd(cmd string, args []string) {
	if path, exists := getCmdPath(cmd); exists {
		c := exec.Command(path, args...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	} else {
		log.Fatalf("%s is not in your PATH...", cmd)
	}
}
