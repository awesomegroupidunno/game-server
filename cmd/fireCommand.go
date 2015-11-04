package cmd

const Fire = "FIRE"

type FireCommand struct {
	BaseCommand
}

func NewFire() FireCommand {
	return FireCommand{BaseCommand{Type: Post, Subtype: Fire, UniqueId: ""}}
}

func (b *FireCommand) Command() *BaseCommand {
	return &(b.BaseCommand)
}
