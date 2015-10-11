package cmd

const Connect = "CONNECT"

type ConnectCommand struct {
	BaseCommand
	Value string
}

func NewConnect(v string) ConnectCommand {
	tc := ConnectCommand{BaseCommand{Type: Post, Subtype: Connect, UniqueId: ""}, v}
	return tc
}

func (b *ConnectCommand) Command() BaseCommand {
	return b.BaseCommand
}

func (b *ConnectCommand) OwnerId() string {
	return b.UserId
}

func (b *ConnectCommand) SetOwnerId(s string) {
	b.UserId = s
}
