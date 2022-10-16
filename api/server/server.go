package server

import "log"

func Init() {
	r := InitRouter()
	err := r.Run("localhost:8080")
	if err != nil {
		log.Printf("Could not initialize gin server: %v", err)
	}
}
