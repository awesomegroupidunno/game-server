package cmd

const Turn = "TURN"

type TurnCommand struct {
	BaseCommand
	Value float32
}

func NewTurn(v float32) TurnCommand {
	tc := TurnCommand{BaseCommand{Type: Post, Subtype: Turn, UniqueId: ""}, v}
	return tc
}

func (b *TurnCommand) Command() *BaseCommand {
	return &(b.BaseCommand)
}

