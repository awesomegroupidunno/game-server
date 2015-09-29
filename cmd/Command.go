package cmd

type BaseCommand struct {
	Type     string `json:"Type"`
	Subtype  string `json:"Subtype"`
	UniqueId string `json:"UniqueId"`
}

type GameCommand interface {
	Command() BaseCommand
}

func (b *BaseCommand) Command() BaseCommand {
	return *b
}
