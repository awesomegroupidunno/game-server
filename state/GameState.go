package state

type GameState struct {
	Val              string
	Vehicles         []*Vehicle
	Bases            []*Base
	ShieldGenerators []*ShieldGenerator
	Bullets          []*Bullet
	GameOver         bool
}

func (g *GameState) Copy() GameState {
	stateCopy := GameState{}
	stateCopy.Val = g.Val
	stateCopy.GameOver = g.GameOver
	copy(stateCopy.Vehicles, g.Vehicles)
	copy(stateCopy.Bases, g.Bases)
	copy(stateCopy.ShieldGenerators, g.ShieldGenerators)
	copy(stateCopy.Bullets, g.Bullets)

	for i := 0; i < len(stateCopy.Vehicles); i++ {
		stateCopy.Vehicles[i].IsMe = false
	}

	return stateCopy
}

func NewGameState() GameState {

	bases := []*Base{}
	b1 := Base{X: 30, Y: 30, CurrentHealth: 1000, MaxHealth: 1000, Width: 20, TeamId: 0}
	b2 := Base{X: 300, Y: 400, CurrentHealth: 1000, MaxHealth: 1000, Width: 20, TeamId: 1}
	bases = append(bases, &b1, &b2)

	generators := []*ShieldGenerator{}
	g1 := ShieldGenerator{X: 300, Y: 30, CurrentHealth: 1000, MaxHealth: 1000, Width: 25, TeamId: 0}
	g2 := ShieldGenerator{X: 300, Y: 40, CurrentHealth: 1000, MaxHealth: 1000, Width: 25, TeamId: 1}
	generators = append(generators, &g1, &g2)

	state := GameState{
		Val:              "",
		Vehicles:         []*Vehicle{},
		Bases:            bases,
		ShieldGenerators: generators,
		GameOver:         false,
		Bullets:          []*Bullet{}}
	return state
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
