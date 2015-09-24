package main

import (
	"fmt"
	"github.com/awesomegroupidunno/game-server/encoder"
	"github.com/awesomegroupidunno/game-server/net"
)

func main() {

	decoder := encoder.JsonDecoder{}
	a := net.UdpReceiver{PortNumber: ":10001", MaxPacket: 8192, EncoderDecoder: &decoder}
	a.Run()

	i := 0
	fmt.Scanf("%i", i)
}
