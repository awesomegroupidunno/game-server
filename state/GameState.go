package state

type GameState struct {
	Val             string
	Vehicles        []Vehicle
	Bases           []Base
	ShieldGenerator []ShieldGenerator
}
