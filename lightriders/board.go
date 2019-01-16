package lightriders

const boardWidth = 16
const boardHeight = 16

// Player represents a player in the game
type Player struct {
	name string
}

const (
	// EmptyCell represents an empty cell
	EmptyCell = iota

	// Player1Head cell that contains the player 1 position
	Player1Head

	// Player2Head cell that contains the player 2 position
	Player2Head

	// Player1Tail cell - tail of player 1
	Player1Tail

	// Player2Tail cell - tail of player 1
	Player2Tail
)

// Cell represents a cell of the board grid
// The enum above contains the possibble values
type Cell byte

// Direction representation
const (
	// Up
	Up = iota

	// Left
	Left

	// Down
	Down

	// Right
	Right
)

// Direction representation, above are the available values
type Direction byte

// Grid is a two-dimensional array of bytes, representing the grid state
// The lower left corner has coordinates 0, 0
type Grid [boardWidth][boardHeight]Cell

// Board rerpesents a game board:
// - player1, player 2 - reference to players
// - grid - a two dimensional array, containing the board state
type Board struct {
	player1 *Player
	player2 *Player
	grid    *Grid
}

const (
	// NoWinner if the game continues
	NoWinner = iota

	// Draw if reached end game and the game is draw
	Draw

	// Player1Wins if player 1 wins
	Player1Wins

	// Player2Wins if player 2 wins
	Player2Wins
)

// Winner representation, enum above shows possible values
type Winner byte

// Advance the game 1 move
// Returns a Winner and a Board
// This is a comment
func (b *Board) Advance(p1, p2 Direction) (Winner, *Board) {
	return Draw, nil
}

// AvailableMoves returns a slice of possibble moves for the given player
// If true is passed, retruns the moves for player 1, otherwise for player 2
// Available moves are all except for when trying to move back into the tail
func (b *Board) AvailableMoves(player1 bool) []Direction {
	r := make([]Direction, 0)

	return r
}
