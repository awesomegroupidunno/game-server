package main

import (
	"fmt"
	"github.com/awesomegroupidunno/game-server/net"
)

func main() {
	a := net.UdpReceiver{PortNumber: ":10001"}
	a.Run()

	i:=0
	fmt.Scanf("%i", i)
}
