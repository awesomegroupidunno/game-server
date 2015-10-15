package cmd

const Connect = "CONNECT"

type ConnectCommand struct {
	BaseCommand
	Value string `json:"StrValue"`
}

func NewConnect(v string) ConnectCommand {
	tc := ConnectCommand{BaseCommand{Type: Post, Subtype: Connect, UniqueId: ""}, v}
	return tc
}

func (b *ConnectCommand) Command() *BaseCommand {
	return &(b.BaseCommand)
}
