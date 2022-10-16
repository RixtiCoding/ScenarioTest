package main

import (
	"MessageFilter/db"
	"MessageFilter/filter"
	"MessageFilter/utility"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	dtb, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	fmt.Println("Hello! Fetching all messages...")
	time.Sleep(time.Second * 1)

	messages := utility.FetchMessages()
	filter.FilterInput(os.Stdout, messages, dtb)
	//fmt.Println(dtb.GetAllKeyValues(db.Ctx))

}
