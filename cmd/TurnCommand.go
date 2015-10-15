package cmd

const Turn = "TURN"

type TurnCommand struct {
	BaseCommand
	Value float64 `json:"NumValue"`
}

func NewTurn(v float64) TurnCommand {
	tc := TurnCommand{BaseCommand{Type: Post, Subtype: Turn, UniqueId: ""}, v}
	return tc
}

func (b *TurnCommand) Command() *BaseCommand {
	return &(b.BaseCommand)
}
