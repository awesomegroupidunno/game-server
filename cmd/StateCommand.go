package cmd

type StateCommand struct {
	Command
}

func newState() StateCommand {
	return StateCommand{Command{Type: "GET", Subtype: "STATE"}}
}
