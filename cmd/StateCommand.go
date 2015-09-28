package cmd

type StateCommand struct {
	Command
}

func newState(v FloatType) StateCommand {
	return StateCommand{Command{Type: "GET", Subtype: "STATE"}}
}
