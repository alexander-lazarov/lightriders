package models

import "fmt"

// BoardWidth in cells
const BoardWidth = 16

// BoardHeight in cells
const BoardHeight = 16

// Grid is a two-dimensional array of bytes, representing the grid state
// The lower left corner has coordinates 0, 0
type Grid [BoardWidth][BoardHeight]Cell

// Board rerpesents a game board:
// - P1, P2 - reference to players
// - grid - a two dimensional array, containing the board state
type Board struct {
	P1   *Player
	P2   *Player
	grid *Grid
}

const (
	// NoWinner if the game continues
	NoWinner = iota

	// Draw if reached end game and the game is draw
	Draw

	// P1Wins if player 1 wins
	P1Wins

	// P2Wins if player 2 wins
	P2Wins
)

// Winner representation, enum above shows possible values
type Winner byte

// GetCell the value of a board cell
func (b *Board) GetCell(p Position) Cell {
	return b.grid[p.X][p.Y]
}

func (b *Board) setCell(p Position, value Cell) {
	b.grid[p.X][p.Y] = value
}

func (b *Board) SetDirP1(d Direction) {
	if d != b.P1.prevDirection.opposite() {
		b.P1.direction = d
	}
}

func (b *Board) SetDirP2(d Direction) {
	if d != b.P2.prevDirection.opposite() {
		b.P2.direction = d
	}
}

// Advance the game 1 move
// Returns a Winner and a Board
// This is a comment
func (b *Board) Advance() (Winner, *Board) {
	b.P1.prevDirection = b.P1.direction
	b.P2.prevDirection = b.P2.direction

	// p1oldneck := b.findCell(P1Neck)
	p1neck := b.findCell(P1Head)
	p1head := p1neck.nextPosition(b.P1.direction)

	// p2oldneck := b.findCell(P2Neck)
	p2neck := b.findCell(P2Head)
	p2head := p2neck.nextPosition(b.P2.direction)

	p1crash := false
	p2crash := false

	if p1head == p2head {
		p1crash = true
		p2crash = true
		b.setCell(p1head, Crash)
	} else {
		if b.isOutside(p1head) || b.GetCell(p1head) != EmptyCell {
			p1crash = true
			b.setCell(p1neck, Crash)
		} else {
			// b.setCell(p1oldneck, P1Tail)
			b.setCell(p1neck, P1Neck)
			b.setCell(p1head, P1Head)
		}

		if b.isOutside(p2head) || b.GetCell(p2head) != EmptyCell {
			p2crash = true
			b.setCell(p2neck, Crash)
		} else {
			// b.setCell(p2oldneck, P2Tail)
			b.setCell(p2neck, P2Neck)
			b.setCell(p2head, P2Head)
		}
	}

	if p1crash && p2crash {
		return Draw, b
	} else if p1crash && !p2crash {
		return P2Wins, b
	} else if !p1crash && p2crash {
		return P1Wins, b
	} else {
		return NoWinner, b
	}
}

func (b *Board) isOutside(p Position) bool {
	return p.X < 0 || p.X >= BoardWidth || p.Y < 0 || p.Y >= BoardHeight
}

func (b *Board) findCell(c Cell) (p Position) {
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			if b.GetCell(Position{X: i, Y: j}) == c {
				p.X = i
				p.Y = j

				return
			}
		}
	}

	panic(fmt.Sprintf("Could not find cell %d", c))
}

// InitialBoard creates an initial board
func InitialBoard() *Board {
	P1 := Player{name: "Player 1", direction: Right, prevDirection: Right}
	P2 := Player{name: "Player 2", direction: Left, prevDirection: Left}
	grid := new(Grid)

	b := new(Board)
	b.P1 = &P1
	b.P2 = &P2
	b.grid = grid

	for i := range b.grid {
		for j := range b.grid[i] {
			b.grid[i][j] = EmptyCell
		}
	}

	b.grid[0][7] = P1Head
	b.grid[15][7] = P2Head

	return b
}
