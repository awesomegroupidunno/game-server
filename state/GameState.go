package state

type GameState struct {
	Val             string
	Vehicles        []Vehicle
	Bases           []Base
	ShieldGenerators []ShieldGenerator
	GameOver        bool
}

func (g *GameState) Copy() GameState {
	stateCopy := GameState{}
	stateCopy.Val = g.Val
	stateCopy.GameOver = g.GameOver
	copy(stateCopy.Vehicles, g.Vehicles)
	copy(stateCopy.Bases, g.Bases)
	copy(stateCopy.ShieldGenerators, g.ShieldGenerators)

	return stateCopy
}

func NewGameState() GameState  {
	state := GameState{
		Val:"",
		Vehicles: []Vehicle{},
		Bases: []Base{},
		ShieldGenerators: []ShieldGenerator{},
		GameOver: true}
	return state
}
