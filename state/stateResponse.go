package state

import (
	"net"
)

type StateResponse struct {
	State   GameState
	Address *net.UDPAddr
}
