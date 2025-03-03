// Command terminal provides a terminal-based interface for Bot Racer
package main

import (
	"log"

	"github.com/davidyarbrough/bot-racer/internal/game"
	"github.com/davidyarbrough/bot-racer/internal/track"
)

func main() {
	// Initialize track generator
	generator := track.NewGenerator()

	// Generate a test track
	track, err := generator.GenerateFromGUID("test-guid")
	if err != nil {
		log.Fatal(err)
	}

	// Create new game state
	state := game.New(track)

	// TODO: Initialize terminal UI and game loop
	_ = state // Temporary to avoid unused variable
}
