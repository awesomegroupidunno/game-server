package cmd

type TurnCommand struct {
	value FloatType
	Command
}

func newTurn(v FloatType) TurnCommand {
	return TurnCommand{Command{Type: "POST", Subtype: "TURN"}, value: v}
}
