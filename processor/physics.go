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
	VehicleWidth                float64
	VehicleHeight               float64
	BulletVelocity              float64
	BulletWidth                 float64
}

func DefaultPhysics() Physics {
	return Physics{
		AccelerationCommandModifier: 5.0,
		TurnCommandModifier:         3.0,
		MaxVehicleVelocity:          150.0,
		FrictionSpeedLoss:           20.0,
		VehicleWidth:                50,
		VehicleHeight:               75,
		BulletVelocity:              200.0,
		BulletWidth:                 10,
	}
}

func (p *Physics) MoveVehicle(vehicle *state.Vehicle, duration time.Duration) {
	x, y := p.move2d(vehicle.X, vehicle.Y, vehicle.Angle, vehicle.Velocity, duration)

	vehicle.X = x
	vehicle.Y = y
}

func (p *Physics) MoveBullet(bullet *state.Bullet, duration time.Duration) {
	x, y := p.move2d(bullet.X, bullet.Y, bullet.Angle, bullet.Velocity, duration)

	bullet.X = x
	bullet.Y = y
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

func (p *Physics) VehicleCollisionPhysics(v1, v2 *state.Vehicle) {
	cos1, sin1 := splitComponent(v1.Angle)
	cos2, sin2 := splitComponent(v2.Angle)

	px1 := (cos1 * v1.Velocity) * v1.Mass
	py1 := (sin1 * v1.Velocity) * v1.Mass

	px2 := (cos2 * v2.Velocity) * v2.Mass
	py2 := (sin2 * v2.Velocity) * v2.Mass

	totalpx := (px1 + px2) / 4
	totalpy := (py1 + py2) / 4

	v1.Velocity = combineComponents(totalpx/3, totalpy/3*2)
	v2.Velocity = combineComponents(totalpx/3*2, totalpy/3)

	if math.Abs(v1.Velocity) >= p.MaxVehicleVelocity*2 {
		v1.Velocity = math.Abs(v1.Velocity) / v1.Velocity * p.MaxVehicleVelocity * 2
	}
	if math.Abs(v2.Velocity) >= p.MaxVehicleVelocity*2 {
		v2.Velocity = math.Abs(v2.Velocity) / v2.Velocity * p.MaxVehicleVelocity * 2
	}

	v1.Angle = math.Atan2(totalpy, totalpx) * RadToDeg
	v2.Angle = math.Atan2(totalpy, totalpx) * RadToDeg

}

func splitComponent(angle float64) (x, y float64) {
	rad := DegToRad * angle
	xAngle := math.Cos(rad)
	yAngle := math.Sin(rad)
	return xAngle, yAngle
}

func combineComponents(x, y float64) (resultant float64) {
	return math.Sqrt(x*x + y*y)
}

// calculates new x and y position for 2d motion
// 		returns x,y
func (p *Physics) move2d(x, y, angle, velocity float64, duration time.Duration) (float64, float64) {
	xAngle, yAngle := splitComponent(angle)
	x = x + (velocity * duration.Seconds() * xAngle)
	y = y + (velocity * duration.Seconds() * yAngle)
	return x, y

}
