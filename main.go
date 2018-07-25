/*
Copyright 2018 Jonathan Monette

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jmoney8080/go-gadget-slack"
)

var (
	// Info Logger
	Info *log.Logger
	// Warning Logger
	Warning *log.Logger
	// Error Logger
	Error *log.Logger

	slackClient *slack.Client

	gifs []string
)

func init() {

	Info = log.New(os.Stdout,
		"[INFO]: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"[WARNING]: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stderr,
		"[ERROR]: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	slackClient = slack.New(http.Client{Timeout: 10 * time.Second}, os.Getenv("SLACK_WEBHOOK"))

	gifs = strings.Split(os.Getenv("GIFS"), ",")
}

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest handling the request
func HandleRequest(ctx context.Context) error {
	payload := slack.Payload{
		Channel: os.Getenv("SLACK_CHANNEL"),
		Attachments: []slack.Attachment{
			{
				Text:     "FORCE PULL!!!!!",
				ImageURL: gifs[rand.Intn(len(gifs)-1)],
			},
		},
	}

	status, err := slackClient.Send(payload)
	if err != nil {
		Error.Println(err)
	} else {
		Info.Println(status)
	}

	return nil
}
