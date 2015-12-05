package state

import "github.com/awesomegroupidunno/game-server/collision"

var _ collision.Box2d = Bullet{}

type Rocket struct {
	Point
	Sized
	Velocity     float64
	Angle        float64
	Target       *Vehicle
	ShouldRemove bool `json:"-"`
}

func (b Rocket) Position() (x float64, y float64) {
	return b.X, b.Y
}

func (b Rocket) Size() (width float64, height float64) {
	return b.Height, b.Width
}

func (b Rocket) AngleDegrees() float64 {
	return b.Angle
}
