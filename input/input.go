package input

import (
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// HandleInputToPlayer checks if any key from the keyset has been pressed
// and updates the player direction
func HandleInputToPlayer(board *models.Player, keyset *Keyset) {
	for in, out := range *keyset {
		if inpututil.IsKeyJustPressed(in) {
			board.SetDirection(out)
		}
	}
}

// GetInputDir checks if any key from the keyset has been pressed during
// the last frame and returns the player direction in that case. If no key
// has been pressed the second argument is bool
func GetInputDir(keyset *Keyset) (models.Direction, bool) {
	for in, out := range *keyset {
		if inpututil.IsKeyJustPressed(in) {
			return out, true
		}
	}

	return 0, false
}
