package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func cleanBoard() *Board {
	b := InitialBoard()
	b.Grid[0][8] = EmptyCell
	b.Grid[15][8] = EmptyCell

	return b
}

func TestInitialBoardAdvanceNoWinner(t *testing.T) {
	b := InitialBoard()

	w, _ := b.Advance()

	assert.Equal(t, w, (Winner)(NoWinner))

	assert.Equal(t, b.Grid[0][8], (Cell)(P1Tail))
	assert.Equal(t, b.Grid[15][8], (Cell)(P2Tail))
	assert.Equal(t, b.Grid[1][8], (Cell)(P1Head))
	assert.Equal(t, b.Grid[14][8], (Cell)(P2Head))
}

func TestInitialBoardAdvanceDraw(t *testing.T) {
	b := InitialBoard()

	b.P1.Direction = Left
	b.P2.Direction = Right

	w, _ := b.Advance()

	assert.Equal(t, w, (Winner)(Draw))
}

func TestInitialBoardAdvanceP1Wins(t *testing.T) {
	b := InitialBoard()

	b.P1.Direction = Right
	b.P2.Direction = Right

	w, _ := b.Advance()

	assert.Equal(t, w, (Winner)(P1Wins))
}

func TestInitialBoardAdvanceP2Wins(t *testing.T) {
	b := InitialBoard()

	b.P1.Direction = Left
	b.P2.Direction = Left

	w, _ := b.Advance()

	assert.Equal(t, w, (Winner)(P2Wins))
}

func TestDrawWhenArriveAtTheSameCell(t *testing.T) {
	b := cleanBoard()
	b.Grid[6][8] = P1Head
	b.Grid[8][8] = P2Head

	w, _ := b.Advance()

	assert.Equal(t, w, (Winner)(Draw))
	assert.Equal(t, b.Grid[7][8], (Cell)(Crash))
}
