package state

import "github.com/awesomegroupidunno/game-server/collision"

var _ collision.Box2d = Bullet{}

type Bullet struct {
	X        float64
	Y        float64
	Width    float64
	Height   float64
	Angle    float64
	Velocity float64 `json:"-"`
	OwnerId  string
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
