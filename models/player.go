package models

// Player represents a player in the game
type Player struct {
	name          string
	direction     Direction
	prevDirection Direction
}

func (p *Player) SetDirection(d Direction) {
	if d != p.prevDirection.opposite() {
		p.direction = d
	}
}
