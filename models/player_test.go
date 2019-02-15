package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDirectionNonOpposite(t *testing.T) {
	p := Player{direction: Up, prevDirection: Up}

	p.SetDirection(Left)
	assert.Equal(t, p.direction, (Direction)(Left))
}

func TestSetDirectionOpposite(t *testing.T) {
	p := Player{direction: Up, prevDirection: Up}

	p.SetDirection(Down)
	assert.Equal(t, p.direction, (Direction)(Up))
}
