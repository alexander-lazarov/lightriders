package input

import (
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func HandleInput(board *models.Player, keyset *Keyset) {
	for in, out := range *keyset {
		if inpututil.IsKeyJustPressed(in) {
			board.SetDirection(out)
		}
	}
}
