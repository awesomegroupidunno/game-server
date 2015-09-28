package cmd

type TurnCommand struct {
	Command
	Value float32
}

func newTurn(v float32) TurnCommand {
	tc := TurnCommand{Command{Type: "POST", Subtype: "TURN", UniqueId: ""}, v}
	return tc
}
