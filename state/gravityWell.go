package state

type GravityWell struct {
	Point
	Sized
	TeamId int
}

func (v GravityWell) Position() (x float64, y float64) {
	return v.X, v.Y
}

func (v GravityWell) Size() (width float64, height float64) {
	return v.Width, v.Height
}

func (v GravityWell) AngleDegrees() float64 {
	return 0.0
}
