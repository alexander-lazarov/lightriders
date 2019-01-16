package lightriders

import (
	"testing"
)

func initialBoard() *Board {
	player1 := Player{name: "Player 1"}
	player2 := Player{name: "Player 2"}
	grid := new(Grid)

	b := new(Board)
	b.player1 = &player1
	b.player2 = &player2
	b.grid = grid

	for i := range b.grid {
		for j := range b.grid[i] {
			b.grid[i][j] = EmptyCell
		}
	}

	b.grid[0][7] = Player1Head
	b.grid[15][7] = Player2Head

	return b
}

func TestInitialBoardAdvanceNoWinner(t *testing.T) {
	b := initialBoard()

	w, _ := b.Advance(Up, Down)

	t.Run("Winner", func(t *testing.T) {
		if w != NoWinner {
			t.Errorf("Expected game winner to be NoWinner")
		}
	})

	t.Run("Player 1 tail is set", func(t *testing.T) {
		if b.grid[0][7] != Player1Tail {
			t.Errorf("Expected player 1 tail to be set")
		}
	})

	t.Run("Player 2 tail is set", func(t *testing.T) {
		if b.grid[15][7] != Player2Tail {
			t.Errorf("Expected player 2 tail to be set")
		}
	})

	t.Run("Player 2 head is set", func(t *testing.T) {
		if b.grid[14][7] != Player1Head {
			t.Errorf("Expected player 1 tail to be set")
		}
	})

	t.Run("Player 1 head is set", func(t *testing.T) {
		if b.grid[14][7] != Player2Head {
			t.Errorf("Expected player 2 tail to be set")
		}
	})
}

func TestInitialBoardAdvanceDraw(t *testing.T) {
	b := initialBoard()

	w, _ := b.Advance(Down, Up)

	t.Run("Winner", func(t *testing.T) {
		if w != Draw {
			t.Errorf("Expected game winner to be NoWinner")
		}
	})

	// TODO - consider what cell type to use when a player crashes into a wall
}

func TestInitialBoardAdvanceP1Wins(t *testing.T) {
	b := initialBoard()

	w, _ := b.Advance(Up, Up)

	t.Run("Winner", func(t *testing.T) {
		if w != Player1Wins {
			t.Errorf("Expected game winner to be NoWinner")
		}
	})

	// TODO - consider what cell type to use when a player crashes into a wall
}

func TestInitialBoardAdvanceP2Wins(t *testing.T) {
	b := initialBoard()

	w, _ := b.Advance(Down, Down)

	t.Run("Winner", func(t *testing.T) {
		if w != Player1Wins {
			t.Errorf("Expected game winner to be NoWinner")
		}
	})

	// TODO - consider what cell type to use when a player crashes into a wall
}
