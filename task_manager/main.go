package main

import (
	"log"
	"task_manager/router"
)

func main() {
	err := router.Run()
	if err != nil {
		log.Printf("Failed to start the server: %s", err.Error())
	}
}

