package discord

import (
	"fmt"
	"github.com/gtuk/discordwebhook"
	"golang-aws-lambda/src/pkg/crawl"
	"log"
	"time"
)

type Discord struct {
	url string
}

func NewDiscord(url string) *Discord {
	return &Discord{
		url: url,
	}
}

func (d Discord) GetProperties() discordwebhook.Message {
	crawl, err := crawl.NewCrawl()
	if err != nil {
		fmt.Println("failed to init crawl")
		log.Fatal(err)
	}

	properties, err := crawl.ConstructDataBank.GetProperties(d.url)
	if err != nil {
		fmt.Println("failed to get properties")
		log.Fatal(err)
	}

	//todo:  textの整形
	fields := make([]discordwebhook.Field, 0)
	for _, property := range properties {
		inline := true
		v := fmt.Sprintf(">>> %s / %s / %s", property.Address, property.Square, property.DetailUrl)
		field := discordwebhook.Field{
			Name:   &property.Title,
			Value:  &v,
			Inline: &inline,
		}
		fields = append(fields, field)
	}

	today := time.Now()
	month := int(today.Month())
	day := today.Day()
	date := fmt.Sprintf(" %d/%d", month, day)
	botName := "まーうんだよ"
	headerTitle := "New" + date + " 建設地 上位10件"

	footerText := "Sent By CAT@不動産を買う仕事してました"
	footer := &discordwebhook.Footer{
		Text: &footerText,
	}

	message := discordwebhook.Message{
		Username: &botName,
		Embeds: &[]discordwebhook.Embed{
			{
				Title:  &headerTitle,
				Url:    &d.url,
				Fields: &fields,
				Footer: footer,
			},
		},
	}
	return message
}
