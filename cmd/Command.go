package cmd

const Post = "POST"
const Get = "GET"

type BaseCommand struct {
	Type     string `json:"Type"`
	Subtype  string `json:"Subtype"`
	UniqueId string `json:"UniqueId"`
	UserId   string `json:"-"`
}

type GameCommand interface {
	Command() *BaseCommand
}

func (b *BaseCommand) Command() *BaseCommand {
	return b
}

