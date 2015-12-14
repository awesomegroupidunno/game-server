package processor

import (
	"github.com/awesomegroupidunno/game-server/collision"
	"github.com/awesomegroupidunno/game-server/state"
	"math"
	"math/rand"
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
	BaseOffset                  float64
	BaseWidth                   float64
	BaseHealth                  int
	ShieldGeneratorHealth       int
	ShieldWidth                 float64
	ShieldOffset                float64
	BulletDamage                int
	RocketDamage                int
	VehicleHealth               int
	VehicleRespawn              time.Duration
	MaxPowerups                 int
	GravityBullets              bool
	PowerupRespawn              time.Duration
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
		BulletDelay:                 100.0 * time.Millisecond,
		BaseOffset:                  45,
		BaseHealth:                  1000,
		BaseWidth:                   40,
		ShieldWidth:                 20,
		ShieldOffset:                30,
		ShieldGeneratorHealth:       400,
		BulletDamage:                8,
		RocketDamage:                20,
		VehicleHealth:               300,
		VehicleRespawn:              5 * time.Second,
		GravityBullets:              false,
		PowerupRespawn:              8 * time.Second,
		MaxPowerups:                 3,
	}
}

func (p *Physics) NewGameState() state.GameState {

	bases := []*state.Base{}

	b1 := state.Base{
		Point:         state.NewPoint(p.BaseOffset, p.WorldHeight-p.BaseOffset),
		Sized:         state.NewSized(p.BaseWidth, p.BaseWidth),
		CurrentHealth: p.BaseHealth,
		MaxHealth:     p.BaseHealth,
		TeamId:        0}

	b2 := state.Base{
		Point:         state.NewPoint(p.WorldWidth-p.BaseOffset, p.BaseOffset),
		Sized:         state.NewSized(p.BaseWidth, p.BaseWidth),
		CurrentHealth: p.BaseHealth,
		MaxHealth:     p.BaseHealth,
		TeamId:        1}

	bases = append(bases, &b1, &b2)

	shields := []*state.Shield{}

	s1 := state.Shield{
		Point:     state.NewPoint(p.WorldWidth-p.BaseOffset, p.BaseOffset),
		Sized:     state.NewSized(p.BaseWidth*1.5, p.BaseWidth*1.5),
		IsEnabled: true,
		TeamId:    1}

	s2 := state.Shield{
		Point:     state.NewPoint(p.BaseOffset, p.WorldHeight-p.BaseOffset),
		Sized:     state.NewSized(p.BaseWidth*1.5, p.BaseWidth*1.5),
		IsEnabled: true,
		TeamId:    0}

	shields = append(shields, &s1, &s2)

	generators := []*state.ShieldGenerator{}
	g1 := state.ShieldGenerator{
		Point:         state.NewPoint(p.WorldWidth-p.ShieldOffset, p.WorldHeight-p.BaseOffset),
		Sized:         state.NewSized(p.ShieldWidth, p.ShieldWidth),
		CurrentHealth: p.ShieldGeneratorHealth,
		MaxHealth:     p.ShieldGeneratorHealth,
		TeamId:        0,
		RespawnTime:   time.Now()}

	g1.Shield = &s2

	g2 := state.ShieldGenerator{
		Point:         state.NewPoint(p.ShieldOffset, p.ShieldOffset),
		Sized:         state.NewSized(p.ShieldWidth, p.ShieldWidth),
		CurrentHealth: p.ShieldGeneratorHealth,
		MaxHealth:     p.ShieldGeneratorHealth,
		TeamId:        1,
		RespawnTime:   time.Now()}

	g2.Shield = &s1
	generators = append(generators, &g1, &g2)

	state := state.GameState{
		Val:              "",
		Vehicles:         []*state.Vehicle{},
		Bases:            bases,
		ShieldGenerators: generators,
		GameOver:         -1,
		Bullets:          []*state.Bullet{},
		Shields:          shields,
		Rockets:          []*state.Rocket{},
		PowerUps:         []*state.Powerup{},
		SecToRestart:     -1,
	}
	return state
}

func (p *Physics) RespawnVehicle(v *state.Vehicle, g state.GameState) bool {
	if !v.IsAlive {
		if time.Now().After(v.TimeDestroyed.Add(p.VehicleRespawn)) {
			v.IsAlive = true
			loc := p.findSpace(v.Sized, g)
			v.Y = loc.Y
			v.X = loc.X
			v.Angle = 0
			v.CurrentHealth = v.MaxHealth
			v.ActivePowerup = NO_POWERUP
			v.StoredPowerup = NO_POWERUP
		}
	}
	return false
}

func (p *Physics) MoveVehicle(vehicle *state.Vehicle, duration time.Duration) {
	x, y := p.move2d(vehicle.X, vehicle.Y, vehicle.Angle, vehicle.Velocity, duration)

	vehicle.X = x
	vehicle.Y = y
}

func (p *Physics) MoveRocket(rocket *state.Rocket, duration time.Duration) {

	t := rocket.Target
	if t == nil {
		rocket.ShouldRemove = true
		return
	}

	dx := t.X - rocket.X
	dy := t.Y - rocket.Y
	rocket.Angle = math.Atan2(dy, dx) * RadToDeg
	x, y := p.move2d(rocket.X, rocket.Y, rocket.Angle, rocket.Velocity, duration)

	rocket.X = x
	rocket.Y = y

	if collision.Collides(rocket, t) {
		rocket.ShouldRemove = true
		p.damageVehicle(t, p.BulletDamage*20)
	}
}

func (p *Physics) MoveBullet(bullet *state.Bullet, duration time.Duration) {
	x, y := p.move2d(bullet.X, bullet.Y, bullet.Angle, bullet.Velocity, duration)

	bullet.X = x
	bullet.Y = y
}

func (p *Physics) ApplyGravity(state *state.GameState) {
	for _, well := range state.GravityWells {
		for _, v := range state.Vehicles {
			if well.Owner != v.Owner {

				delta := distance(well.Point, v.Point)
				dist := combineComponents(delta.X, delta.Y)

				pull := 1 / (dist * dist)

				xPull := delta.X * pull * 100
				yPull := delta.Y * pull * 100
				if xPull >= 5 {
					xPull = 5
				}
				if xPull <= -5 {
					xPull = -5
				}
				if yPull >= 5 {
					yPull = 5
				}
				if yPull <= -5 {
					yPull = 5
				}
				v.X += xPull
				v.Y += yPull
			}
		}
	}
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

func (p *Physics) VehicleCollision(v1, v2 *state.Vehicle) {
	if v1.IsAlive && v2.IsAlive {
		bounciness := 1.5
		p.VehicleKnockback(v1, v1.Angle+180, v1.Velocity*bounciness)
		p.VehicleKnockback(v2, v1.Angle, v1.Velocity)
	}
}

func (p *Physics) PickupPowerUp(v1 *state.Vehicle, power *state.Powerup) {
	if v1.IsAlive && !power.ShouldRemove {
		power.ShouldRemove = true
		v1.StoredPowerup = power.PowerupType
	}
}

// Creates knockback for a vehicle
func (p *Physics) VehicleKnockback(vehicle *state.Vehicle, kbAngle, kbVelocity float64) {
	// Get vehicle velocity vectors
	vehAngleX, vehAngleY := splitComponent(vehicle.Angle)
	vehVectorX := vehAngleX * vehicle.Velocity
	vehVectorY := vehAngleY * vehicle.Velocity

	// Get knockback velocity vectors
	kbAngleX, kbAngleY := splitComponent(kbAngle)
	kbVectorX := kbAngleX * kbVelocity
	kbVectorY := kbAngleY * kbVelocity

	// Combine vectors
	vectorX := vehVectorX + kbVectorX
	vectorY := vehVectorY + kbVectorY
	vehVelocity := combineComponents(vectorX, vectorY)

	// Calculate angle perpendicularity as a percent
	angleFactor := math.Mod(math.Abs(vehicle.Angle-kbAngle+90), 180) / 90.0

	// Set vehicle velocity
	if math.Signbit(vehicle.Velocity) == math.Signbit(vehVelocity) {
		vehicle.Velocity = -vehVelocity * angleFactor
	} else {
		vehicle.Velocity = vehVelocity * angleFactor
	}
}

func (p *Physics) DamageVehicle(v *state.Vehicle, b *state.Bullet) bool {
	b.ShouldRemove = true
	return p.damageVehicle(v, p.BulletDamage)
}

func (p *Physics) damageVehicle(v *state.Vehicle, amount int) bool {
	if v.IsAlive {
		v.CurrentHealth -= amount
		if v.CurrentHealth <= 0 {
			v.CurrentHealth = 0
			v.IsAlive = false
			v.TimeDestroyed = time.Now()
			return true
		}
	}
	return false
}

func (p *Physics) DamageShieldGenerator(b *state.Bullet, s *state.ShieldGenerator) bool {
	s.CurrentHealth -= p.BulletDamage
	b.ShouldRemove = true
	if s.CurrentHealth <= 0 {
		s.Shield.IsEnabled = false
		s.CurrentHealth = 0
		return true
	}
	return false

}

func (p *Physics) DamageBase(b *state.Bullet, base *state.Base) bool {
	base.CurrentHealth -= p.BulletDamage
	b.ShouldRemove = true
	if base.CurrentHealth <= 0 {
		base.CurrentHealth = 0
		return true
	}
	return false
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

var powerupsSpawned int = 0

func (p *Physics) SpawnPowerup(g *state.GameState) {
	powerupType := rand.Intn(NUM_POWERUPS) + 1
	powerupType = (powerupsSpawned % NUM_POWERUPS) + 1
	powerupsSpawned++

	size := state.Sized{30, 30}

	newPowerup := state.Powerup{
		Point:        p.findSpace(size, *g),
		Sized:        size,
		ShouldRemove: false,
		PowerupType:  powerupType,
	}

	g.PowerUps = append(g.PowerUps, &newPowerup)
}

func (p *Physics) BoundBullet(b *state.Bullet) {
	if !p.inBounds(b.X, b.Y) {
		b.ShouldRemove = true
	}
}

func (p *Physics) inBounds(x, y float64) bool {
	oob := x < 0
	oob = oob || y < 0
	oob = oob || x > p.WorldWidth
	oob = oob || y > p.WorldHeight
	return !oob
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

func distance(p1, p2 state.Point) state.Point {
	return state.NewPoint(p1.X-p2.X, p1.Y-p2.Y)
}

// calculates new x and y position for 2d motion
// 		returns x,y
func (p *Physics) move2d(x, y, angle, velocity float64, duration time.Duration) (float64, float64) {
	xAngle, yAngle := splitComponent(angle)
	x = x + (velocity * duration.Seconds() * xAngle)
	y = y + (velocity * duration.Seconds() * yAngle)
	return x, y

}

func CleanupBullets(data []*state.Bullet) []*state.Bullet {
	retData := []*state.Bullet{}
	for i := 0; i < len(data); i++ {
		if !data[i].ShouldRemove {
			retData = append(retData, data[i])
		}
	}
	return retData
}

func CleanupPowerups(data []*state.Powerup) []*state.Powerup {
	retData := []*state.Powerup{}
	for i := 0; i < len(data); i++ {
		if !data[i].ShouldRemove {
			retData = append(retData, data[i])
		}
	}
	return retData
}
func CleanupRockets(data []*state.Rocket) []*state.Rocket {
	retData := []*state.Rocket{}
	for i := 0; i < len(data); i++ {
		if !data[i].ShouldRemove {
			retData = append(retData, data[i])
		}
	}
	return retData
}
func CleanupWells(data []*state.GravityWell) []*state.GravityWell {
	retData := []*state.GravityWell{}
	for i := 0; i < len(data); i++ {
		if !data[i].ShouldRemove {
			retData = append(retData, data[i])
		}
	}
	return retData
}

func (p *Physics) findSpace(size state.Sized, gamestate state.GameState) state.Point {
	pt := state.Point{
		X: float64(rand.Intn(int(p.WorldWidth * 4 / 5))),
		Y: float64(rand.Intn(int(p.WorldHeight * 4 / 5))),
	}
	if p.isValidPlace(pt, size, gamestate) {
		return pt
	}

	return p.findSpace(size, gamestate)
}

func (p *Physics) isValidPlace(pt state.Point, size state.Sized, gamestate state.GameState) bool {
	hasCollision := false

	powerup := state.Powerup{
		Sized: size,
		Point: pt,
	}
	for _, b := range gamestate.Bases {
		if collision.Collides(powerup, b) {
			hasCollision = true
		}
	}
	for _, v := range gamestate.Vehicles {
		if collision.Collides(powerup, v) {
			hasCollision = true
		}
	}
	for _, b := range gamestate.Shields {
		if collision.Collides(powerup, b) {
			hasCollision = true
		}
	}
	for _, s := range gamestate.ShieldGenerators {
		if collision.Collides(powerup, s) {
			hasCollision = true
		}
	}
	for _, p := range gamestate.PowerUps {
		if collision.Collides(powerup, p) {
			hasCollision = true
		}
	}

	return !hasCollision
}
