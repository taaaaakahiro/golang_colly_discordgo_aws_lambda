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

func (d Discord) GetMessage() discordwebhook.Message {
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

	botName := "建築看板通知 Bot"
	footerText := "Sent By CAT@不動産を買う仕事してました"
	footer := &discordwebhook.Footer{
		Text: &footerText,
	}

	// 新着なし
	if len(properties) == 0 {
		content := "本日の新着はありません"
		m := discordwebhook.Message{
			Username: &botName,
			Content:  &content,
			Embeds:   &[]discordwebhook.Embed{},
		}
		return m
	}

	//新着あり
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

	title := makeTitle(len(properties))
	message := discordwebhook.Message{
		Username: &botName,
		Embeds: &[]discordwebhook.Embed{
			{
				Title:  &title,
				Url:    &d.url,
				Fields: &fields,
				Footer: footer,
			},
		},
	}
	return message
}

func makeTitle(cnt int) string {
	today := time.Now()
	month := int(today.Month())
	day := today.Day()
	date := fmt.Sprintf(" %d/%d", month, day)
	return fmt.Sprintf("New %s 建設地 上位%d件", date, cnt)
}
