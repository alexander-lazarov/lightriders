package models

const (
	// EmptyCell represents an empty cell
	EmptyCell = iota

	// P1Head cell that contains the player 1 position
	P1Head

	// P2Head cell that contains the player 2 position
	P2Head

	// P1Tail cell - tail of player 1
	P1Tail

	// P2Tail cell - tail of player 2
	P2Tail

	// Crash cell - when p1 or p2 hits the board edge or another player
	Crash
)

// Cell represents a cell of the board grid
// The enum above contains the possibble values
type Cell byte
