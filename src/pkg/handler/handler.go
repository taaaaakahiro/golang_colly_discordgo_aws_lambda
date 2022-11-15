package handler

import (
	"fmt"
	"github.com/gtuk/discordwebhook"
	"golang-aws-lambda/src/pkg/discord"
	"log"
	"os"
)

func Handler() {
	// env
	hook := os.Getenv("HOOK_REAL_ESTATE")
	url := os.Getenv("TARGET_URL")
	d := discord.NewDiscord(url)
	message := d.GetProperties()

	err := discordwebhook.SendMessage(hook, message)
	if err != nil {
		fmt.Println("failed to send message")
		log.Fatal(err)
	}
}
