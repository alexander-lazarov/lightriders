package models

// Player represents a player in the game
type Player struct {
	name          string
	direction     Direction
	prevDirection Direction
}
