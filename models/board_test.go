package models

import (
	"testing"
)

func TestInitialBoardAdvanceNoWinner(t *testing.T) {
	b := InitialBoard()

	w, _ := b.Advance(Right, Left)

	t.Run("Winner", func(t *testing.T) {
		if w != NoWinner {
			t.Errorf("Expected game winner to be NoWinner %d", w)
		}
	})

	t.Run("Player 1 neck is set", func(t *testing.T) {
		if b.grid[0][7] != P1Neck {
			t.Errorf("Expected player 1 neck to be set")
		}
	})

	t.Run("Player 2 neck is set", func(t *testing.T) {
		if b.grid[15][7] != P2Neck {
			t.Errorf("Expected player 2 neck to be set")
		}
	})

	t.Run("Player 2 head is set", func(t *testing.T) {
		if b.grid[1][7] != P1Head {
			t.Errorf("Expected player 1 head to be set")
		}
	})

	t.Run("Player 2 neck is set", func(t *testing.T) {
		if b.grid[14][7] != P2Head {
			t.Errorf("Expected player 2 head to be set")
		}
	})
}

func TestInitialBoardAdvanceDraw(t *testing.T) {
	b := InitialBoard()

	w, _ := b.Advance(Left, Right)

	t.Run("Winner", func(t *testing.T) {
		if w != Draw {
			t.Errorf("Expected game winner to be Draw")
		}
	})

	// TODO - consider what cell type to use when a player crashes into a wall
}

func TestInitialBoardAdvanceP1Wins(t *testing.T) {
	b := InitialBoard()

	w, _ := b.Advance(Left, Right)

	t.Run("Winner", func(t *testing.T) {
		if w != P1Wins {
			t.Errorf("Expected game winner to be NoWinner")
		}
	})

	// TODO - consider what cell type to use when a player crashes into a wall
}

func TestInitialBoardAdvanceP2Wins(t *testing.T) {
	b := InitialBoard()

	w, _ := b.Advance(Left, Left)

	t.Run("Winner", func(t *testing.T) {
		if w != P1Wins {
			t.Errorf("Expected game winner to be NoWinner")
		}
	})

	// TODO - consider what cell type to use when a player crashes into a wall
}
