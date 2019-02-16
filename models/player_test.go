package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDirectionNonOpposite(t *testing.T) {
	p := Player{Direction: Up, PrevDirection: Up}

	p.SetDirection(Left)
	assert.Equal(t, p.Direction, (Direction)(Left))
}

func TestSetDirectionOpposite(t *testing.T) {
	p := Player{Direction: Up, PrevDirection: Up}

	p.SetDirection(Down)
	assert.Equal(t, p.Direction, (Direction)(Up))
}
