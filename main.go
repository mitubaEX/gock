package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	execCommand := strings.Join(os.Args[1:], " ")
	log.Println(execCommand)
	result, err := exec.Command("sh", "-c", execCommand).Output()
	if err != nil {
		panic(err)
	}
	log.Println(string(result))
}
