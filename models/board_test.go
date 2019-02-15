package models

import (
	"testing"
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

	b.SetDirP1(Left)
	b.SetDirP2(Right)

	w, _ := b.Advance()

	t.Run("Winner", func(t *testing.T) {
		if w != Draw {
			t.Errorf("Expected game winner to be Draw")
		}
	})
}

func TestInitialBoardAdvanceP1Wins(t *testing.T) {
	b := InitialBoard()

	b.SetDirP1(Right)
	b.SetDirP2(Right)

	w, _ := b.Advance()

	t.Run("Winner", func(t *testing.T) {
		if w != P1Wins {
			t.Errorf("Expected game winner to be P1")
		}
	})
}

func TestInitialBoardAdvanceP2Wins(t *testing.T) {
	b := InitialBoard()

	b.SetDirP1(Left)
	b.SetDirP2(Left)

	w, _ := b.Advance()

	t.Run("Winner", func(t *testing.T) {
		if w != P2Wins {
			t.Errorf("Expected game winner to be P2")
		}
	})
}

func TestDrawWhenArriveAtTheSameCell(t *testing.T) {
	b := cleanBoard()
	b.grid[6][7] = P1Head
	b.grid[8][7] = P2Head

	w, _ := b.Advance()

	t.Run("Winner", func(t *testing.T) {
		if w != Draw {
			t.Errorf("Expected game winner to be Draw")
		}
	})

	t.Run("Crash cell", func(t *testing.T) {
		if b.grid[7][7] != Crash {
			t.Errorf("Expected the mid cell to be a crash cell")
		}
	})
}
