package main

import (
	"fmt"
	"github.com/awesomegroupidunno/game-server/net"
)

func main() {
	a := net.UdpReceiver{PortNumber: ":10001", MaxPacket: 8192}
	a.Run()

	i := 0
	fmt.Scanf("%i", i)
}
