package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func PostSlack(text string) {
	loadEnv()
	token := os.Getenv("SLACK_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")
	slack_url := os.Getenv("SLACK_URL")
	params := url.Values{
		"token":   {token},
		"channel": {channel},
		"text":    {text},
	}
	resp, err := http.PostForm(slack_url, params)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	log.Printf("Request Success: %s", resp.Status)
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}
