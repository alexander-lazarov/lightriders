package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func cleanBoard() *Board {
	b := InitialBoard()
	b.grid[0][7] = EmptyCell
	b.grid[15][7] = EmptyCell

	return b
}

func TestInitialBoardAdvanceNoWinner(t *testing.T) {
	b := InitialBoard()

	w, _ := b.Advance()

	assert.Equal(t, w, (Winner)(NoWinner))

	assert.Equal(t, b.grid[0][7], (Cell)(P1Tail))
	assert.Equal(t, b.grid[15][7], (Cell)(P2Tail))
	assert.Equal(t, b.grid[1][7], (Cell)(P1Head))
	assert.Equal(t, b.grid[14][7], (Cell)(P2Head))
}

func TestInitialBoardAdvanceDraw(t *testing.T) {
	b := InitialBoard()

	b.P1.direction = Left
	b.P2.direction = Right

	w, _ := b.Advance()

	assert.Equal(t, w, (Winner)(Draw))
}

func TestInitialBoardAdvanceP1Wins(t *testing.T) {
	b := InitialBoard()

	b.P1.direction = Right
	b.P2.direction = Right

	w, _ := b.Advance()

	assert.Equal(t, w, (Winner)(P1Wins))
}

func TestInitialBoardAdvanceP2Wins(t *testing.T) {
	b := InitialBoard()

	b.P1.direction = Left
	b.P2.direction = Left

	w, _ := b.Advance()

	assert.Equal(t, w, (Winner)(P2Wins))
}

func TestDrawWhenArriveAtTheSameCell(t *testing.T) {
	b := cleanBoard()
	b.grid[6][7] = P1Head
	b.grid[8][7] = P2Head

	w, _ := b.Advance()

	assert.Equal(t, w, (Winner)(Draw))
	assert.Equal(t, b.grid[7][7], (Cell)(Crash))
}
