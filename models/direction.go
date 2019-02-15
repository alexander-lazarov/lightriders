package models

// Direction representation
const (
	// Up
	Up = iota

	// Left
	Left

	// Down
	Down

	// Right
	Right
)

// Direction representation, above are the available values
type Direction byte

func (d Direction) opposite() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	default:
		panic("Unknown direction")
	}
}
