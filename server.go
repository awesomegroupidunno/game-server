package main

import (
	"github.com/awesomegroupidunno/game-server/encoder"
	"github.com/awesomegroupidunno/game-server/game"
	"github.com/awesomegroupidunno/game-server/network"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
	"sync"
)

func main() {

	var waiter sync.WaitGroup
	waiter.Add(1)
	log.Println("Server starting up")
	decoder := encoder.JsonEncoderDecoder{}
	log.Println("Encoder created")

	manager := game.GameManager{}
	log.Println("Gamemanager created")

	ack_channel := make(chan state.Ack, 100)
	log.Println("Acking channel created")

	state_channel := make(chan state.StateResponse, 100)
	log.Println("State channel created")

	router := game.CommandRouter{Acks: ack_channel, Responses: state_channel, GameManager: &manager}
	log.Println("Router created")

	a := network.UdpReceiver{PortNumber: ":10001", MaxPacket: 8192, EncoderDecoder: &decoder, Router: router, Acks: ack_channel, Responses: state_channel}
	log.Println("Udp reciever created")
	go manager.Start()
	log.Println("Gamemanager started")

	a.Run()
	log.Println("Udp reciever running, press ctr+c to shutdown")

	waiter.Wait()
	log.Println("Shutting down")

}
