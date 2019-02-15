package main

import (
	"sync"
	"time"

	"github.com/alexander-lazarov/lightriders/graphics"
	"github.com/alexander-lazarov/lightriders/input"
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/hajimehoshi/ebiten"
)

var board models.Board
var boardLock sync.RWMutex

func initGame() {
	board = *models.InitialBoard()
	graphics.InitBoardImage()

	go boardLoop()

	ebiten.Run(update, graphics.TotalWidth, graphics.TotalHeight, 2, "Lightriders!")
}

func update(screen *ebiten.Image) error {
	boardLock.Lock()
	input.HandleInput(board.P1, &input.KeysetWASD)
	input.HandleInput(board.P2, &input.KeysetArrows)
	boardLock.Unlock()

	boardLock.RLock()
	graphics.UpdateBoardImage(&board)
	screen.DrawImage(graphics.GetBoardImage(), &ebiten.DrawImageOptions{})
	boardLock.RUnlock()

	return nil
}

func boardLoop() {
	var winner models.Winner = models.NoWinner

	for {
		time.Sleep(1000 * time.Millisecond)

		boardLock.Lock()
		if winner == models.NoWinner {
			winner, _ = board.Advance()
		}
		boardLock.Unlock()
	}
}

func main() {
	initGame()
}
