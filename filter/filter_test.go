package filter

import (
	"MessageFilter/db"
	"MessageFilter/utility"
	"bytes"
	"github.com/go-redis/redis/v8"
	"testing"
)

func FakeNewDatabase() (*db.Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
	if err := client.Ping(db.Ctx).Err(); err != nil {
		return nil, err
	}
	return &db.Database{
		Client: client,
	}, nil
}

func TestFilterInput(t *testing.T) {
	dtb, _ := FakeNewDatabase() // non-existent database, created only because FilterInput requires it.

	// No bad words
	mockMessagesTEST1 := []utility.Message{
		{Body: "# TEXT MESSAGE \n\n This is a test message!", UID: 1},
		{Body: "# TEXT MESSAGE \n\n I love cheesecake.", UID: 2},
	}

	// Only bad words
	mockMessagesTEST2 := []utility.Message{
		{Body: "# TEXT MESSAGE \n\n Death is a strange thing.", UID: 1},
		{Body: "# TEXT MESSAGE \n\n We need to rescue the dog.", UID: 2},
	}

	// 1 internal link + 1 external link
	mockMessagesTEST3 := []utility.Message{
		{Body: "# INTERNAL LINK \n\n - [Section 1](#section1)", UID: 1},
		{Body: "# EXTERNAL LINK \n\n [The Eiffel Tower](https://imageswithtowers.com)", UID: 2},
	}

	tt := []struct {
		messages []utility.Message
		name     string
		expect   string
	}{
		{name: "No bad words", expect: "\n[1] # TEXT MESSAGE \n\n This is a test message!\n\n[2] # TEXT MESSAGE \n\n I love cheesecake.\n", messages: mockMessagesTEST1},
		{name: "Only bad words", expect: "\nMessage with UUID [1] rejected for sensitive content.\n\nMessage with UUID [2] rejected for sensitive content.\n", messages: mockMessagesTEST2},
		{name: "1 internal 1 external", expect: "\n[1] # INTERNAL LINK \n\n - [Section 1](#section1)\n\nMessage with UUID [2] rejected for external link presence.\n", messages: mockMessagesTEST3},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var output bytes.Buffer
			FilterInput(&output, tc.messages, dtb)
			if tc.expect != output.String() {
				t.Errorf("got %s but expected %s", output.String(), tc.expect)
			}
		})
	}

}
