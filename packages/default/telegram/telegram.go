package main

import (
	"fmt"
)

func Main(args map[string]interface{}) map[string]interface{} {
	fmt.Printf("%#v\n", args["__ow_path"])

	return args
}

// func sendMessage(chatID, message string) {
// 	authedURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", os.Getenv("TELEGRAM_API_KEY"))
// 	req, err := http.NewRequest(http.MethodGet, authedURL, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	q := req.URL.Query()
// 	q.Add("chat_id", chatID)
// 	q.Add("text", message)
// 	req.URL.RawQuery = q.Encode()

// 	_, err = http.DefaultClient.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// }
