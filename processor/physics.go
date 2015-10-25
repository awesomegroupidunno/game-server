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
		FrictionSpeedLoss:           20.0}
}

func (p *Physics) MoveVehicle(vehicle *state.Vehicle, duration time.Duration) {
	x, y := p.move2d(vehicle.X, vehicle.Y, vehicle.Angle, vehicle.Velocity, duration)

	vehicle.X = x
	vehicle.Y = y
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

// calculates new x and y position for 2d motion
// 		returns x,y
func (p *Physics) move2d(x, y, angle, velocity float64, duration time.Duration) (float64, float64) {
	rad := DegToRad * angle
	x_angle := math.Cos(rad)
	y_angle := math.Sin(rad)
	x = x + (velocity * duration.Seconds() * x_angle)
	y = y + (velocity * duration.Seconds() * y_angle)
	return x, y

}
