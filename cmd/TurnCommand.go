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

func (b *TurnCommand) Command() BaseCommand {
	return BaseCommand{Type: b.Type, Subtype: b.Subtype, UniqueId: b.UniqueId, UserId: b.UserId}
}

func (b *TurnCommand) OwnerId() string {
	return b.UserId
}

func (b *TurnCommand) SetOwnerId(s string) {
	b.UserId = s
}
