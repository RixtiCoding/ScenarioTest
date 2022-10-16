package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Bannedlist struct {
	Words []string `json:"words"`
}

//FetchBannedWords gets the banned words from our Bannedwords.json
func FetchBannedWords() Bannedlist {
	var badwords Bannedlist

	file, err := os.Open("Bannedwords.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed opening bad words file: %s", err)
	}

	byteValue, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteValue, &badwords)
	if err != nil {
		log.Fatalf("Failed to unmarshal bad words: %s", err)
	}

	return badwords
}
