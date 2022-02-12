package main

import (
	"main/api"
	"main/command"
)

func main() {
	text := command.RunNetworkQualityOneLine()

	// テキスト整形
	text = "```" + text + "```"
	// Slack API実行
	api.PostSlack(text)
}
