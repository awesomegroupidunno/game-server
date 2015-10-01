package state

import "net"

type Ack struct {
	Uuid    string
	Address *net.UDPAddr
}
