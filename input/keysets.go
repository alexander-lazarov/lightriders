package input

import (
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/hajimehoshi/ebiten"
)

// Keyset is a map from input keys to output actions
type Keyset map[ebiten.Key](models.Direction)

// KeysetArrows for playing with the arrows
var KeysetArrows = Keyset{
	ebiten.KeyUp:    models.Up,
	ebiten.KeyDown:  models.Down,
	ebiten.KeyLeft:  models.Left,
	ebiten.KeyRight: models.Right}

// KeysetWASD for playing with WASD keys
var KeysetWASD = Keyset{
	ebiten.KeyW: models.Up,
	ebiten.KeyS: models.Down,
	ebiten.KeyA: models.Left,
	ebiten.KeyD: models.Right}
