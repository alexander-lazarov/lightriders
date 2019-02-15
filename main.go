package main

import (
	"time"

	"github.com/alexander-lazarov/lightriders/graphics"
	"github.com/alexander-lazarov/lightriders/input"
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	input.HandleInput(board.P1, &input.KeysetWASD)
	input.HandleInput(board.P2, &input.KeysetArrows)
	graphics.UpdateBoardImage(&board)
	screen.DrawImage(graphics.GetBoardImage(), &ebiten.DrawImageOptions{})

	return nil
}

func boardLoop() {
	var winner models.Winner = models.NoWinner

	for {
		time.Sleep(1000 * time.Millisecond)

		if winner == models.NoWinner {
			winner, _ = board.Advance()
		}
	}
}

var board models.Board

func main() {
	board = *models.InitialBoard()
	graphics.InitBoardImage()

	go boardLoop()

	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
