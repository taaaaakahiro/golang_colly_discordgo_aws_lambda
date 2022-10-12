package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gtuk/discordwebhook"
)

type Request struct {
	ID   float64 `json: "id"`
	Name string  `json: "name"`
	Like string  `json: "like"`
}

type Response struct {
	YouName string `json: "name"`
	YouLike string `json: "like"`
}

func main() {
	lambda.Start(handler)
}

func handler() {

	var username = "BotUser"
	var content = "This is a test message"
	var url = os.Getenv("HOOK_URL")

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
	return
}
