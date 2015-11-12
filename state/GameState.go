package state

import (
	"math"
)

type GameState struct {
	Val              string
	Vehicles         []*Vehicle
	Bases            []*Base
	ShieldGenerators []*ShieldGenerator
	Bullets          []*Bullet
	Shields          []*Shield
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
	stateCopy.GameOver = g.GameOver

	for i := 0; i < len(g.Vehicles); i++ {
		var v Vehicle = *g.Vehicles[i]
		v.X = math.Floor(v.X)
		v.Y = math.Floor(v.Y)
		v.Angle = math.Floor(v.Angle)
		v.IsMe = false
		stateCopy.Vehicles = append(stateCopy.Vehicles, &v)
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
