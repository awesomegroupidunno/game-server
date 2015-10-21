package processor

import (
	"github.com/awesomegroupidunno/game-server/state"
	"math"
	"time"
)

const (
	RadToDeg = 180 / math.Pi
	DegToRad = math.Pi / 180
)

type Physics struct {
	AccelerationCommandModifier float64
	TurnCommandModifier         float64
	MaxVehicleVelocity          float64
	FrictionSpeedLoss           float64
}

func DefaultPhysics() Physics {
	return Physics{
		AccelerationCommandModifier: 5.0,
		TurnCommandModifier:         3.0,
		MaxVehicleVelocity:          150.0,
		FrictionSpeedLoss:           0.25}
}

func (p *Physics) MoveVehicle(vehicle *state.Vehicle, duration time.Duration) {
	rad := DegToRad * vehicle.Angle
	x_angle := math.Cos(rad)
	y_angle := math.Sin(rad)
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
