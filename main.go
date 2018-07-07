package main

import (
	"errors"
	"github.com/nlopes/slack"
	"log"
	"os"
	"strings"
	"os/exec"
	"fmt"
)

func getEnv() (string, error) {
	if env := os.Getenv("SLACK_TOKEN"); env != "" {
		return env, nil
	} else {
		return "", errors.New("SLACK_TOKEN is not found.")
	}
}

func main() {
	// fetch token from env variable
	slackToken, err := getEnv()
	if err != nil {
		log.Fatal(err)
	}

	// exec command
	execCommand := strings.Join(os.Args[1:], " ")
	result, err := exec.Command("sh", "-c", execCommand).Output()
	if err != nil {
		log.Fatal(err)
	}

	// print result
	fmt.Printf("%v", string(result))

	resultOfExecutedCommand := "```$ " + execCommand + "\n" + string(result) + "```"

	// post result of executed command to slack
	api := slack.New(slackToken)
	channels, err := api.GetChannels(false)
	if err != nil {
		log.Fatal("%s\n", err)
	}
	for _, channel := range channels {
		if channel.Name != "bots" {
			continue
		}

		params := slack.PostMessageParameters{}
		_, _, err := api.PostMessage(channel.ID, resultOfExecutedCommand, params)
		if err != nil {
			log.Fatal("%s\n", err)
		}
	}
}
