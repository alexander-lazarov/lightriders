package models

const (
	// EmptyCell represents an empty cell
	EmptyCell = iota

	// P1Head cell that contains the player 1 position
	P1Head

	// P1Neck cell that contains the previous head position
	P1Neck

	// P1Tail cell - tail of player 1
	P1Tail

	// P2Head cell that contains the player 2 position
	P2Head

	// P2Neck cell that contains the previous head position
	P2Neck

	// P2Tail cell - tail of player 1
	P2Tail

	// Crash cell - when p1 or p2 hits the board edge or another player
	Crash
)

// Cell represents a cell of the board grid
// The enum above contains the possibble values
type Cell byte
