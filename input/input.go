package input

import (
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func HandleInput(board *models.Board) {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		board.SetDirP1(models.Up)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		board.SetDirP1(models.Down)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		board.SetDirP1(models.Right)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		board.SetDirP1(models.Left)
	}
}
