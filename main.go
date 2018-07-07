package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
	"os/exec"
	"strings"
)

var channelName = flag.String("c", "", "target channel name")

func getEnv() (string, error) {
	if env := os.Getenv("SLACK_TOKEN"); env != "" {
		return env, nil
	} else {
		return "", errors.New("token is not found")
	}
}

func main() {
	// fetch token from env variable
	slackToken, err := getEnv()
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()

	// exec command
	execCommand := strings.Join(flag.Args(), " ")
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
		log.Fatal(err)
	}
	for _, channel := range channels {
		if channel.Name != *channelName {
			continue
		}

		params := slack.PostMessageParameters{}
		_, _, err := api.PostMessage(channel.ID, resultOfExecutedCommand, params)
		if err != nil {
			log.Fatal(err)
		}
	}
}
