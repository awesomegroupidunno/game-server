package state

type Shield struct {
	Point
	Sized
	TeamId    int
	IsEnabled bool
}

func (v Shield) Position() (x float64, y float64) {
	return float64(v.X), float64(v.Y)
}

func (v Shield) Size() (width float64, height float64) {
	return float64(v.Width), float64(v.Width)
}

func (v Shield) AngleDegrees() float64 {
	return 0.0
}
