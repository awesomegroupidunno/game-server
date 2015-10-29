package state

import "net"

type Ack struct {
	UUID    string
	Address *net.UDPAddr
}
