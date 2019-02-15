package input

import (
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/hajimehoshi/ebiten"
)

type Keyset map[ebiten.Key](models.Direction)

var KeysetArrows = Keyset{
	ebiten.KeyUp:    models.Up,
	ebiten.KeyDown:  models.Down,
	ebiten.KeyLeft:  models.Left,
	ebiten.KeyRight: models.Right}

var KeysetWASD = Keyset{
	ebiten.KeyW: models.Up,
	ebiten.KeyS: models.Down,
	ebiten.KeyA: models.Left,
	ebiten.KeyD: models.Right}
