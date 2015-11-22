package state

import "github.com/awesomegroupidunno/game-server/collision"

var _ collision.Box2d = Powerup{}

type Powerup struct {
	Point
	Sized
	ShouldRemove bool `json:"-"`
	PowerupType  int
}

func (p Powerup) Position() (x float64, y float64) {
	return p.X, p.Y
}

func (p Powerup) Size() (width float64, height float64) {
	return p.Height, p.Width
}

func (p Powerup) AngleDegrees() float64 {
	return 0
}
