package processor

import (
	"github.com/awesomegroupidunno/game-server/state"
	"math"
	"time"
)

type Physics struct {
	AccelerationCommandModifier float64
	TurnCommandModifier         float64
}

func DefaultPhysics() Physics {
	return Physics{
		AccelerationCommandModifier: 1.0,
		TurnCommandModifier:         1.0}
}

func (p *Physics) MoveVehicle(vehicle *state.Vehicle, duration time.Duration) {
	x_angle := math.Cos(vehicle.Angle)
	y_angle := math.Sin(vehicle.Angle)
	vehicle.X = vehicle.X + (vehicle.Velocity * duration.Seconds() * x_angle)
	vehicle.Y = vehicle.Y + (vehicle.Velocity * duration.Seconds() * y_angle)
}
