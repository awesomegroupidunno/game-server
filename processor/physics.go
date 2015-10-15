package processor

type Physics struct {
	AccelerationCommandModifier float64
	TurnCommandModifier         float64
}

func DefaultPhysics() Physics {
	return Physics{
		AccelerationCommandModifier: 1.0,
		TurnCommandModifier:         1.0}
}
