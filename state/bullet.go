package state

import "github.com/awesomegroupidunno/game-server/collision"

var _ collision.Box2d = Bullet{}

type Bullet struct {
	Point
	Sized
	Angle        float64
	Velocity     float64 `json:"-"`
	OwnerId      string  `json:"-"`
	ShouldRemove bool    `json:"-"`
}

func (b Bullet) Position() (x float64, y float64) {
	return b.X, b.Y
}

func (b Bullet) Size() (width float64, height float64) {
	return b.Height, b.Width
}

func (b Bullet) AngleDegrees() float64 {
	return b.Angle
}
