package state

type Base struct {
	Point
	Sized
	MaxHealth     int
	CurrentHealth int
	ShieldEnabled bool
	TeamId        int
}

func (v Base) Position() (x float64, y float64) {
	return v.X, v.Y
}

func (v Base) Size() (width float64, height float64) {
	return v.Width, v.Height
}

func (v Base) AngleDegrees() float64 {
	return 0.0
}
