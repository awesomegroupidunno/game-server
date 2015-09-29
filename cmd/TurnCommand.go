package cmd

type TurnCommand struct {
	BaseCommand
	Value float32
}

func NewTurn(v float32) TurnCommand {
	tc := TurnCommand{BaseCommand{Type: "POST", Subtype: "TURN", UniqueId: ""}, v}
	return tc
}

func (b *TurnCommand) Command() BaseCommand {
	return BaseCommand{Type: b.Type, Subtype: b.Subtype, UniqueId: b.UniqueId}
}
