package main

import (
	"fmt"
	"github.com/awesomegroupidunno/game-server/encoder"
	"github.com/awesomegroupidunno/game-server/game"
	"github.com/awesomegroupidunno/game-server/network"
	"github.com/awesomegroupidunno/game-server/state"
)

func main() {

	decoder := encoder.JsonEncoderDecoder{}
	router := game.CommandRouter{}
	ack_channel := make(chan state.Ack, 100)
	state_channel := make(chan state.StateResponse, 100)
	router.Acks = ack_channel
	router.Responses = state_channel
	a := network.UdpReceiver{PortNumber: ":10001", MaxPacket: 8192, EncoderDecoder: &decoder, Router: router, Acks: ack_channel, Responses: state_channel}
	a.Run()

	i := 0
	fmt.Scanf("%i", i)
}
