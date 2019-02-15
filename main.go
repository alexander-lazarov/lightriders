package main

import (
	"time"

	"github.com/alexander-lazarov/lightriders/graphics"
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	graphics.UpdateBoardImage(&board)
	screen.DrawImage(graphics.GetBoardImage(), &ebiten.DrawImageOptions{})

	return nil
}

var board models.Board

func main() {
	board = *models.InitialBoard()
	graphics.InitBoardImage()

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			board.Advance(models.Right, models.Left)
		}
	}()

	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
