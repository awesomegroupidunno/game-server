package main

import (
	"github.com/awesomegroupidunno/game-server/encoder"
	"github.com/awesomegroupidunno/game-server/game"
	"github.com/awesomegroupidunno/game-server/network"
	"github.com/awesomegroupidunno/game-server/processor"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
	"sync"
)

func main() {

	/*
		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			//f.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println("This is a test log entry")
	*/
	var waiter sync.WaitGroup
	waiter.Add(1)
	log.Println("Server starting up")
	decoder := encoder.JsonEncoderDecoder{}
	log.Println("Encoder created")

	ack_channel := make(chan state.Ack, 100)
	log.Println("Acking channel created")

	state_channel := make(chan state.StateResponse)
	log.Println("State channel created")

	physics := processor.DefaultPhysics()
	log.Println("Default Physics created")

	factory := processor.NewFactory(&physics)
	log.Println("Command Processor Factory created")

	gameManager := game.NewManager(state.NewGameState(), state_channel, &factory)
	log.Println("Gamemanager created")

	router := game.CommandRouter{Acks: ack_channel, GameManager: &gameManager}
	log.Println("Router created")

	receiver := network.UdpReceiver{PortNumber: ":10001",
		MaxPacket:      8192,
		EncoderDecoder: &decoder,
		Router:         router,
		Acks:           ack_channel,
		Responses:      state_channel}

	log.Println("Udp reciever created")
	go gameManager.Start()
	log.Println("Gamemanager started")

	receiver.Run()
	log.Println("Udp reciever running, press ctr+c to shutdown")

	waiter.Wait()
	log.Println("Shutting down")

}
