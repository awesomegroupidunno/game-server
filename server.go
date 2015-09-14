package main

import (
	"fmt"
	"github.com/awesomegroupidunno/game-server/net"
)

func main() {
	fmt.Println("hello world")
	a:= net.UdpReceiver{PortNumber:":10001"}
	fmt.Println(a)
	a.Start();
}
