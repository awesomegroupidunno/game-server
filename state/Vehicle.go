package state

import (
	"github.com/awesomegroupidunno/game-server/collision"
	"time"
)

var _ collision.Box2d = Vehicle{}

type Vehicle struct {
	Point
	Sized
	Velocity          float64
	Angle             float64
	TeamId            int
	MaxHealth         int
	CurrentHealth     int
	Owner             string
	Mass              float64 `json:"-"`
	IsMe              bool
	IsAlive           bool
	TimeDestroyed     time.Time `json:"-"`
	OverRideSpeedTill time.Time `json:"-"`
	ActivePowerup     int
	StoredPowerup     int
}

func (v Vehicle) Position() (x float64, y float64) {
	return v.X, v.Y
}

func (v Vehicle) Size() (width float64, height float64) {
	return v.Height, v.Width
}

func (v Vehicle) AngleDegrees() float64 {
	return v.Angle
}
