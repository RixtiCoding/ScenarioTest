package filter

import (
	"MessageFilter/ApprovalService"
	"MessageFilter/api/helpers"
	"MessageFilter/db"
	"MessageFilter/utility"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func FilterInput(w io.Writer, msgs []utility.Message, dtb *db.Database) {
	badwords := FetchBadWords()

outer:
	for _, msg := range msgs {
		for _, word := range badwords {
			if strings.Contains(strings.ToLower(msg.Body), word) {
				dtb.SetToRedis(db.Ctx, msg.Body, "LANGUAGE")
				fmt.Fprintf(w, "\nMessage with UUID [%v] rejected for sensitive content.\n", msg.UID)
				continue outer
			} else if utility.ContainsUrl(msg) {
				if utility.IsImage(msg) {
					ApprovalService.ApprovalService(os.Stdout, msg)
					continue outer
				}
				fmt.Fprintf(w, "\nMessage with UUID [%v] rejected for external link presence.\n", msg.UID)
				continue outer
			}

		}
		fmt.Fprintf(w, "\n[%v] %v\n", msg.UID, msg.Body)
	}

}

func FetchBadWords() []string {
	var wordlist helpers.Bannedlist
	resp, err := http.Get("http://127.0.0.1:8080/badwords")
	if err != nil {
		log.Println("Could not fetch banned words list")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Could not read response body")
	}

	err = json.Unmarshal(body, &wordlist.Words)
	if err != nil {
		log.Println("Could not unmarshal response body")
	}
	return wordlist.Words
}
