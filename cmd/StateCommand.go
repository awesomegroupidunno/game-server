package cmd

type StateCommand struct {
	BaseCommand
}

func NewState() StateCommand {
	return StateCommand{BaseCommand{Type: "GET", Subtype: "STATE"}}
}

func (b *StateCommand) Command() BaseCommand {
	return BaseCommand{Type: b.Type, Subtype: b.Subtype, UniqueId: b.UniqueId}
}
