package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	bot_token := os.Getenv("SLACK_BOT_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")

	api := slack.New(bot_token)
	channelArr := []string{channel_id}
	fileArr := []string{""}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %s\t URL: %s\t", file.Name, file.URL)
	}

}
