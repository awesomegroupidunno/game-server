package processor

import (
	"github.com/awesomegroupidunno/game-server/state"
	"math"
	"time"
)

type Physics struct {
	AccelerationCommandModifier float64
	TurnCommandModifier         float64
	MaxVehicleVelocity          float64
	FrictionSpeedLoss           float64
}

func DefaultPhysics() Physics {
	return Physics{
		AccelerationCommandModifier: 1.0,
		TurnCommandModifier:         1.0,
		MaxVehicleVelocity:          4.0,
		FrictionSpeedLoss:           1.00}
}

func (p *Physics) MoveVehicle(vehicle *state.Vehicle, duration time.Duration) {
	x_angle := math.Cos(vehicle.Angle)
	y_angle := math.Sin(vehicle.Angle)
	vehicle.X = vehicle.X + (vehicle.Velocity * duration.Seconds() * x_angle)
	vehicle.Y = vehicle.Y + (vehicle.Velocity * duration.Seconds() * y_angle)
}

func (p *Physics) VehicleFrictionSlow(vehicle *state.Vehicle, duration time.Duration) {
	speedLoss := p.FrictionSpeedLoss * duration.Seconds()
	if vehicle.Velocity > 0 {
		vehicle.Velocity = float64(vehicle.Velocity) - float64(speedLoss)
		if vehicle.Velocity < 0 {
			vehicle.Velocity = 0
		}
	} else {
		vehicle.Velocity = float64(vehicle.Velocity) + float64(speedLoss)
		if vehicle.Velocity > 0 {
			vehicle.Velocity = 0
		}
	}
}
