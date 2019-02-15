package models

import (
	"testing"
)

func TestInitialNextPosition(t *testing.T) {

	t.Run("Up", func(t *testing.T) {
		actual := Position{X: 0, Y: 0}.nextPosition(Up)
		expected := Position{X: 0, Y: -1}

		if actual != expected {
			t.Errorf("Expected (%d, %d), got (%d, %d)", expected.X, expected.Y, actual.X, actual.Y)
		}
	})
}
