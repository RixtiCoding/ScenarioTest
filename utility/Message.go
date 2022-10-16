package utility

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Message struct {
	UID  int    `json:"UID"`
	Body string `json:"Body"`
}

// FetchMessages gets the hard-coded messages from Messages.json
func FetchMessages() []Message {
	var Messages []Message

	file, err := os.Open("Messages.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed opening Messages file: %s", err)
	}

	byteValue, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteValue, &Messages)
	if err != nil {
		log.Fatalf("Failed to unmarshal messages: %s", err)
	}

	return Messages
}
