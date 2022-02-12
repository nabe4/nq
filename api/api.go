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
	fmt.Println(token)
	params := url.Values{}
	params.Add("token", token)
	params.Add("channel", "#general")
	params.Add("text", text)
	resp, err := http.PostForm("https://slack.com/api/chat.postMessage", params)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	log.Printf("Request Success: %s", resp.Status)
	// fmt.Println(resp)
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// message := os.Getenv("SAMPLE_MESSAGE")
	// fmt.Println(message)
}
