package state

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

type Sized struct {
	Width  float64
	Height float64
}

func NewSized(w, h float64) Sized {
	return Sized{Height: h, Width: w}
}

type GameState struct {
	Val              string
	Vehicles         []*Vehicle
	Bases            []*Base
	ShieldGenerators []*ShieldGenerator
	Bullets          []*Bullet
	Shields          []*Shield
	Rockets          []*Rocket
	PowerUps         []*Powerup
	GameOver         bool
}

func (g *GameState) Copy() GameState {
	stateCopy := GameState{}
	stateCopy.Val = g.Val
	stateCopy.GameOver = g.GameOver
	stateCopy.Bullets = []*Bullet{}
	stateCopy.Vehicles = []*Vehicle{}
	stateCopy.Bases = []*Base{}
	stateCopy.ShieldGenerators = []*ShieldGenerator{}
	stateCopy.Shields = []*Shield{}
	stateCopy.Rockets = []*Rocket{}
	stateCopy.GameOver = g.GameOver

	for i := 0; i < len(g.Vehicles); i++ {
		var v Vehicle = *g.Vehicles[i]
		v.X = math.Floor(v.X)
		v.Y = math.Floor(v.Y)
		v.Angle = math.Floor(v.Angle)
		v.ActivePowerup = g.Vehicles[i].ActivePowerup
		v.IsMe = false
		if v.IsAlive {
			stateCopy.Vehicles = append(stateCopy.Vehicles, &v)
		}
	}
	for i := 0; i < len(g.Bullets); i++ {
		var b Bullet = *g.Bullets[i]
		b.X = math.Floor(b.X)
		b.Y = math.Floor(b.Y)
		b.Angle = math.Floor(b.Angle)
		stateCopy.Bullets = append(stateCopy.Bullets, &b)
	}
	for i := 0; i < len(g.Bases); i++ {
		var b Base = *g.Bases[i]
		stateCopy.Bases = append(stateCopy.Bases, &b)
	}

	for i := 0; i < len(g.ShieldGenerators); i++ {
		var b ShieldGenerator = *g.ShieldGenerators[i]
		stateCopy.ShieldGenerators = append(stateCopy.ShieldGenerators, &b)
	}
	for i := 0; i < len(g.Shields); i++ {
		var b Shield = *g.Shields[i]
		stateCopy.Shields = append(stateCopy.Shields, &b)
	}
	for i := 0; i < len(g.Rockets); i++ {
		var b Rocket = *g.Rockets[i]
		stateCopy.Rockets = append(stateCopy.Rockets, &b)
	}

	return stateCopy
}

// returns a pointer to the vehicle with the owner's string id
// returns nil if no vehicle is found
func (g *GameState) GetVehicle(owner string) *Vehicle {
	for _, vehicle := range g.Vehicles {
		if vehicle.Owner == owner {
			return vehicle
		}
	}

	return nil
}
