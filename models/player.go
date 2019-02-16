package models

// Player represents a player in the game
type Player struct {
	Name          string
	Direction     Direction
	PrevDirection Direction
}

// SetDirection checks if the change in direction is a valid one
// and updates the player if so
func (p *Player) SetDirection(d Direction) {
	if d != p.PrevDirection.opposite() {
		p.Direction = d
	}
}
