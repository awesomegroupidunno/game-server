package game

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"net"
)

type CommandRouter struct {
	Responses   chan state.StateResponse
	Acks        chan state.Ack
	GameManager GameManager
}

func (r *CommandRouter) RouteCommand(c *cmd.GameCommand, address *net.UDPAddr) {

	commandType := (*c).Command().Type
	if commandType == "GET" {
		r.routeGet(c, address)
	} else if commandType == "POST" {
		r.routePost(c, address)
	}
}

func (r *CommandRouter) routeGet(c *cmd.GameCommand, address *net.UDPAddr) {
	currentState := r.GameManager.TakeState()
	r.Responses <- state.StateResponse{State: currentState, Address: address}
}
func (r *CommandRouter) routePost(c *cmd.GameCommand, address *net.UDPAddr) {
	r.Acks <- state.Ack{Uuid: (*c).Command().UniqueId, Address: address}
}
