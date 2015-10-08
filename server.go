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

	ack_channel := make(chan state.Ack, 100)
	log.Println("Acking channel created")

	state_channel := make(chan state.StateResponse)
	log.Println("State channel created")

	gameManager := game.NewManager(state.NewGameState(), state_channel)
	log.Println("Gamemanager created")

	router := game.CommandRouter{Acks: ack_channel, GameManager: &gameManager}
	log.Println("Router created")

	reciever := network.UdpReceiver{PortNumber: ":10001",
		MaxPacket:      8192,
		EncoderDecoder: &decoder,
		Router:         router,
		Acks:           ack_channel,
		Responses:      state_channel}

	log.Println("Udp reciever created")
	go gameManager.Start()
	log.Println("Gamemanager started")

	reciever.Run()
	log.Println("Udp reciever running, press ctr+c to shutdown")

	waiter.Wait()
	log.Println("Shutting down")

}
