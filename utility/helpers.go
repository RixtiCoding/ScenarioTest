package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mvdan/xurls"
	"log"
	rand "math/rand"
	"net/http"
	"strings"
	"time"
)

type ImageMessage struct {
	Body string
}

func ContainsUrl(message Message) bool {
	rxStrict := xurls.Strict
	//FindAllString performs a check on the message to see if it contains any valid urls
	sanitizedUrls := rxStrict.FindAllString(message.Body, -1)
	if sanitizedUrls != nil {
		return true
	}

	return false
}

//IsImage currently recognizes only .jpg images. in a real word scenario it would recognize more.
func IsImage(message Message) bool {
	rxStrict := xurls.Strict
	imageUrls := rxStrict.FindAllString(message.Body, -1)
	var IsImage bool
	for _, imageUrl := range imageUrls {
		if strings.HasSuffix(imageUrl, ".jpg") {
			IsImage = true
			StoreImage(imageUrl)
			// this will get printed for each image present in the message.
			fmt.Printf("\nMessage with UUID [%v] contains an image link. Stored for approval.\n", message.UID)
		}

	}

	return IsImage
}

func StoreImage(imageUrl string) {
	var img = ImageMessage{Body: imageUrl}

	j, err := json.Marshal(&img)
	if err != nil {
		log.Println("Could not marshal image json")
	}

	_, _ = http.Post("http://127.0.0.1:8080/images", "application/json", bytes.NewBuffer(j))

}

func RandomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
