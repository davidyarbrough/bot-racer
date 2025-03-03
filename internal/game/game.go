// Package game provides core game mechanics and state management
package game

import (
	"github.com/davidyarbrough/bot-racer/internal/track"
)

// State represents the current state of a race
type State struct {
	Track      *track.Track
	PlayerLane int
	Speed      float64
	Position   float64
	Lap        int
}

// New creates a new game state with initial values
func New(track *track.Track) *State {
	return &State{
		Track:      track,
		PlayerLane: track.Width / 2, // Start in middle lane
		Speed:      0,
		Position:   0,
		Lap:        1,
	}
}

// Update advances the game state based on input and time delta
func (s *State) Update(delta float64, input Input) {
	// TODO: Implement game state update logic
}

// Input represents player input for a game tick
type Input struct {
	Left  bool
	Right bool
}
