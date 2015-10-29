package collision

type Box2d interface {
	Position() (float64, float64)
	Size() (float64, float64)
	AngleDegrees() float64
}
