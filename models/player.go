package models

// Player represents a player in the game
type Player struct {
	Name          string
	Direction     Direction
	PrevDirection Direction
}

func (p *Player) SetDirection(d Direction) {
	if d != p.PrevDirection.opposite() {
		p.Direction = d
	}
}
