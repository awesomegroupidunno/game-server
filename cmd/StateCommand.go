package cmd

const State = "STATE"

type StateCommand struct {
	BaseCommand
}

func NewState() StateCommand {
	return StateCommand{BaseCommand{Type: Get, Subtype: State}}
}

func (b *StateCommand) Command() BaseCommand {
	return BaseCommand{Type: b.Type, Subtype: b.Subtype, UniqueId: b.UniqueId}
}
