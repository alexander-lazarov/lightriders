package models

// Position represents a position on the grid
type Position struct {
	X int
	Y int
}

// nextPosition in the grid
func (p Position) nextPosition(d Direction) (result Position) {
	result = p

	switch d {
	case Up:
		result.Y--
	case Down:
		result.Y++
	case Left:
		result.X--
	case Right:
		result.X++
	default:
		panic("Unknown position")
	}

	return
}
