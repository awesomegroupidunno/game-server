package state

import "time"

type ShieldGenerator struct {
	TeamId        int
	MaxHealth     int
	CurrentHealth int
	RespawnTime   time.Time `json:"-"`
	X             int
	Y             int
	Width         int
}

func (v ShieldGenerator) Position() (x float64, y float64) {
	return float64(v.X), float64(v.Y)
}

func (v ShieldGenerator) Size() (width float64, height float64) {
	return float64(v.Width), float64(v.Width)
}

func (v ShieldGenerator) AngleDegrees() float64 {
	return 0.0
}
