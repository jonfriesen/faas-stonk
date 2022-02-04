package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type TelegramMessage struct {
	Text string `json:"text,omitempty"`
	Chat struct {
		ID string `json:"id,omitempty"`
	} `json:"chat,omitempty"`
}

func Main(args map[string]interface{}) map[string]interface{} {
	o, err := json.MarshalIndent(args, "", "	")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", o)

	if v, e := args["message"]; e {
		o, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		var msg TelegramMessage
		err = json.Unmarshal(o, &msg)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", msg)

		sendMessage(msg.Chat.ID, fmt.Sprintf("You said: %s", msg.Text))
	}

	return args
}

func sendMessage(chatID, message string) {
	authedURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", os.Getenv("TELEGRAM_API_KEY"))
	req, err := http.NewRequest(http.MethodGet, authedURL, nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("chat_id", chatID)
	q.Add("text", message)
	req.URL.RawQuery = q.Encode()

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
}
