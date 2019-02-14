package main

import (
	"github.com/alexander-lazarov/lightriders/graphics"
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	b := models.InitialBoard()
	i := graphics.DrawBoard(b)

	screen.DrawImage(i, &ebiten.DrawImageOptions{})

	return nil
}

func main() {
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
