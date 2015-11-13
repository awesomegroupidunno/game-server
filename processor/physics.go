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
	WorldWidth                  float64
	WorldHeight                 float64
	AccelerationCommandModifier float64
	TurnCommandModifier         float64
	MaxVehicleVelocity          float64
	FrictionSpeedLoss           float64
	VehicleWidth                float64
	VehicleHeight               float64
	BulletVelocity              float64
	BulletWidth                 float64
	BulletDelay                 time.Duration
	BaseOffset                  int
	BaseWidth                   int
	BaseHealth                  int
	ShieldHealth                int
	ShieldWidth                 int
	ShieldOffset                int
}

func DefaultPhysics() Physics {
	return Physics{
		WorldWidth:                  640.0,
		WorldHeight:                 480.0,
		AccelerationCommandModifier: 5.0,
		TurnCommandModifier:         3.0,
		MaxVehicleVelocity:          150.0,
		FrictionSpeedLoss:           20.0,
		VehicleWidth:                25,
		VehicleHeight:               37,
		BulletVelocity:              250.0,
		BulletWidth:                 7,
		BulletDelay:                 300.0 * time.Millisecond,
		BaseOffset:                  45,
		BaseHealth:                  1000,
		BaseWidth:                   40,
		ShieldWidth:                 20,
		ShieldOffset:                30,
		ShieldHealth:                1000,
	}
}

func (p *Physics) NewGameState() state.GameState {

	bases := []*state.Base{}

	b1 := state.Base{X: int(p.BaseOffset),
		Y:             int(p.WorldHeight) - p.BaseOffset,
		CurrentHealth: p.BaseHealth,
		MaxHealth:     p.BaseHealth,
		Width:         p.BaseWidth,
		TeamId:        0}

	b2 := state.Base{X: int(p.WorldWidth) - p.BaseOffset,
		Y:             int(p.BaseOffset),
		CurrentHealth: p.BaseHealth,
		MaxHealth:     p.BaseHealth,
		Width:         p.BaseWidth,
		TeamId:        1}

	bases = append(bases, &b1, &b2)

	shields := []*state.Shield{}

	s1 := state.Shield{X: int(p.WorldWidth) - p.BaseOffset,
		IsEnabled: true,
		Y:         int(p.BaseOffset),
		Width:     int(float64(p.BaseWidth) * 1.5),
		TeamId:    1}

	s2 := state.Shield{X: int(p.BaseOffset),
		IsEnabled: true,
		Y:         int(p.WorldHeight) - p.BaseOffset,
		Width:     int(float64(p.BaseWidth) * 1.5),
		TeamId:    0}

	shields = append(shields, &s1, &s2)

	generators := []*state.ShieldGenerator{}
	g1 := state.ShieldGenerator{X: int(p.WorldWidth) - p.ShieldOffset,
		Y:             int(p.WorldHeight) - p.BaseOffset,
		CurrentHealth: p.ShieldHealth,
		MaxHealth:     p.ShieldHealth,
		Width:         p.ShieldWidth,
		TeamId:        0,
		RespawnTime:   time.Now()}

	g2 := state.ShieldGenerator{X: p.ShieldOffset,
		Y:             p.ShieldOffset,
		CurrentHealth: p.ShieldHealth,
		MaxHealth:     p.ShieldHealth,
		Width:         p.ShieldWidth,
		TeamId:        1,
		RespawnTime:   time.Now()}
	generators = append(generators, &g1, &g2)

	state := state.GameState{
		Val:              "",
		Vehicles:         []*state.Vehicle{},
		Bases:            bases,
		ShieldGenerators: generators,
		GameOver:         false,
		Bullets:          []*state.Bullet{},
		Shields:          shields}
	return state
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
	//TODO:
}

func (p *Physics) VehicleBounding(v *state.Vehicle) {
	shouldStop := v.X < 0

	if shouldStop {
		v.X = 2
	}

	if v.Y < 0 {
		shouldStop = true
		v.Y = 2
	}

	if v.Y > p.WorldHeight {
		shouldStop = true
		v.Y = p.WorldHeight - 2
	}
	if v.X > p.WorldWidth {
		shouldStop = true
		v.X = p.WorldWidth - 2
	}
	if shouldStop {
		v.Velocity = 0
	}

}

func (p *Physics) CleanUpBullets(bullets []*state.Bullet) []*state.Bullet {
	toRemove := []bool{}
	for i := 0; i < len(bullets); i++ {
		bullet := bullets[i]
		shouldRemove := bullet.X < 0
		shouldRemove = shouldRemove || bullet.Y < 0
		shouldRemove = shouldRemove || bullet.X > p.WorldWidth
		shouldRemove = shouldRemove || bullet.Y > p.WorldHeight

		toRemove = append(toRemove, shouldRemove)
	}
	bullets = cleanupBullets(bullets, toRemove)
	return bullets
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

func cleanupBullets(data []*state.Bullet, ids []bool) []*state.Bullet {
	if len(ids) == 0 {
		return data
	}
	retData := []*state.Bullet{}
	for i := 0; i < len(data); i++ {
		if !ids[i] {
			retData = append(retData, data[i])
		}
	}
	return retData
}
