package cmd

const POWERUP = "POWERUP"

type PowerupCommand struct {
	BaseCommand
}

func NewPowerup() PowerupCommand {
	return PowerupCommand{BaseCommand{Type: Post, Subtype: POWERUP, UniqueId: ""}}
}

func (p *PowerupCommand) Command() *BaseCommand {
	return &(p.BaseCommand)
}
