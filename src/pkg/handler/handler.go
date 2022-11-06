package handler

import (
	"fmt"
	"github.com/gtuk/discordwebhook"
	"golang-aws-lambda/src/pkg/crawl"
	"log"
	"os"
	"strings"
	"time"
)

func Handler() {
	// env
	hook := os.Getenv("HOOK_ESC_KEY")
	url := os.Getenv("URL")

	crawl, err := crawl.NewCrawl()
	if err != nil {
		fmt.Println("failed to init crawl")
		log.Fatal(err)
	}

	properties, err := crawl.ConstructDataBank.GetProperties(url)
	if err != nil {
		fmt.Println("failed to get properties")
		log.Fatal(err)
	}

	text := make([]string, 0)
	//todo:  textの整形
	for _, property := range properties {
		text = append(text,
			"```",
			property.Title,
			property.DetailUrl,
			property.Address,
			property.Square,
			"```",
		)
	}
	t := strings.Join(text, ",")

	today := time.Now()
	month := int(today.Month())
	day := today.Day()
	date := fmt.Sprintf(" %d/%d", month, day)
	botName := "まーうんだよ"
	headerTitle := "New" + date + " 新着物件"
	footerText := "Sent By CAT@不動産を買う仕事してました"
	footer := &discordwebhook.Footer{
		Text: &footerText,
	}

	message := discordwebhook.Message{
		Username: &botName,
		Embeds: &[]discordwebhook.Embed{
			{Title: &headerTitle, Url: &url, Description: &t, Footer: footer},
		},
	}

	err = discordwebhook.SendMessage(hook, message)
	if err != nil {
		fmt.Println("failed to send message")
		log.Fatal(err)
	}
}
