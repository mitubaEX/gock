package main

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func getEnv() (string, error) {
	if env := os.Getenv("WEBHOOK_URL"); env != "" {
		return env, nil
	} else {
		return "", errors.New("WEBHOOK_URL not found")
	}
}

func main() {
	webhookURL, err := getEnv()
	if err != nil {
		panic(err)
	}
	log.Println(webhookURL)

	execCommand := strings.Join(os.Args[1:], " ")
	log.Println(execCommand)
	result, err := exec.Command("sh", "-c", execCommand).Output()
	if err != nil {
		panic(err)
	}
	log.Println(string(result))

	postData := `{"text": "` + "```$ " + execCommand + "\n" + string(result) + "```" + `"}`

	req, err := http.NewRequest(
		"POST",
		webhookURL,
		bytes.NewBuffer([]byte(postData)),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
