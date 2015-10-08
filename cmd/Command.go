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
	Command() BaseCommand
	OwnerId() string
	SetOwnerId(s string)
}

func (b *BaseCommand) Command() BaseCommand {
	return *b
}

func (b *BaseCommand) OwnerId() string {
	return b.UserId
}

func (b *BaseCommand) SetOwnerId(s string) {
	b.UserId = s
}
