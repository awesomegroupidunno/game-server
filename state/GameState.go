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

	return stateCopy
}

func NewGameState() GameState {
	state := GameState{
		Val:              "",
		Vehicles:         []*Vehicle{},
		Bases:            []*Base{},
		ShieldGenerators: []*ShieldGenerator{},
		GameOver:         false}
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
