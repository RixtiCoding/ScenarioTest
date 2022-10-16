package utility

import (
	"testing"
)

func TestContainsUrl(t *testing.T) {
	msg := Message{Body: "https://noschemesite.com"}

	got := ContainsUrl(msg)
	if got != true {
		t.Error("ContainsUrl returned false, want true")
	}

}

func TestIsImage(t *testing.T) {
	msg := Message{Body: "https://fakesite/cat.jpg"}

	got := IsImage(msg)
	if got != true {
		t.Error("IsImage returned false, want true ")
	}
}

// nothing to test here
//func TestStoreImage(t *testing.T) {
//
//}
