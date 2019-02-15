package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOppositePanics(t *testing.T) {
	invalidDir := (Direction)(255)

	assert.Panics(t, func() { invalidDir.opposite() }, "Does not panic on invalid directions")
}

func TestOppositeDirs(t *testing.T) {
	a := map[Direction]Direction{
		Up:    Down,
		Down:  Up,
		Left:  Right,
		Right: Left,
	}

	for in, out := range a {
		assert.Equal(t, in.opposite(), out)
	}
}
