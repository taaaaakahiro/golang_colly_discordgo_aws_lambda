package handler

import (
	"fmt"
	"github.com/gtuk/discordwebhook"
	"golang-aws-lambda/src/pkg/crawl"
	"golang-aws-lambda/src/pkg/domain/entity"
)

func Handler(hook string, url string) error {
	item, err := crawl.GetItem(url)
	if err != nil {
		fmt.Errorf("failed to crawl. %v", err)
	}

	user := &entity.BotUser{
		UserName: "まーうん",
		Content:  item.Title + "/" + url, //todo: 表示方法要検討
	}

	message := discordwebhook.Message{
		Username: &user.UserName,
		Content:  &user.Content,
	}

	err = discordwebhook.SendMessage(hook, message)
	if err != nil {
		return fmt.Errorf("failed to send message. %v", err)
	}
	return nil
}
