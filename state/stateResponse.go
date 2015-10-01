package state

import (
	"net"
)

type StateResponse struct {
	GameState
	Address *net.UDPAddr
}
