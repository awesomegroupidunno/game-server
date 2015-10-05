package state

type GameState struct {
	Val             string
	Vehicles        []Vehicle
	Bases           []Base
	ShieldGenerator []ShieldGenerator
	GameOver        bool
}

func (g *GameState) Copy() GameState {
	n := GameState{}
	n.Val = g.Val
	n.GameOver = g.GameOver
	copy(n.Vehicles, g.Vehicles)
	copy(n.Bases, g.Bases)
	copy(n.ShieldGenerator, g.ShieldGenerator)

	return n
}
