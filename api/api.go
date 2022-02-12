package api

import (
	"log"
	"net/http"
	"net/url"
)

func PostSlack(text string) {
	params := url.Values{}
	params.Add("token", "xoxb-2102527046113-3079034186087-tCqKymQmPnA98LxI8z3Ev6Y7")
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
