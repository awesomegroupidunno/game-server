package game

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"net"
)

type CommandRouter struct {
	Acks        chan state.Ack
	GameManager *GameManager
}

// Routes the command to the proper place
// Currently just puts commands in the GameManager's commands list
func (r *CommandRouter) RouteCommand(c *cmd.GameCommand, address *net.UDPAddr) {
	(*c).Command().UserId = address.IP.String()
	commandType := (*c).Command().Type
	if commandType == cmd.Get {
	} else if commandType == cmd.Post {
		r.routePost(c, address)
	}
}

// Sends a Post Command to the right place
// Adds command to the GameManager's commands list
// Places an ack in the Ack channel
func (r *CommandRouter) routePost(c *cmd.GameCommand, address *net.UDPAddr) {
	r.GameManager.AddCommand((*c))
	r.Acks <- state.Ack{UUID: (*c).Command().UniqueId, Address: address}
}
