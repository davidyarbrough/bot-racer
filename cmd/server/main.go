// Command server provides the Bot Racer multiplayer server
package main

import (
	"log"
	"net/http"
)

func main() {
	// TODO: Initialize database connection
	// TODO: Set up WebSocket handler
	// TODO: Set up REST API routes

	log.Println("Starting Bot Racer server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
