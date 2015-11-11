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

	ackChannel := make(chan state.Ack, 100)
	log.Println("Acking channel created")

	stateChannel := make(chan state.StateResponse)
	log.Println("State channel created")

	physics := processor.DefaultPhysics()
	log.Println("Default Physics created")

	factory := processor.NewFactory(&physics)
	log.Println("Command Processor Factory created")

	current_state := physics.NewGameState()

	gameManager := game.NewManager(current_state, stateChannel, &factory)
	log.Println("Gamemanager created")

	router := game.CommandRouter{Acks: ackChannel, GameManager: &gameManager}
	log.Println("Router created")

	receiver := network.UdpReceiver{PortNumber: ":10001",
		MaxPacket:      8192,
		EncoderDecoder: &decoder,
		Router:         router,
		Acks:           ackChannel,
		Responses:      stateChannel}

	log.Println("Udp reciever created")
	go gameManager.Start()
	log.Println("Gamemanager started")

	receiver.Run()
	log.Println("Udp reciever running, press ctr+c to shutdown")

	waiter.Wait()
	log.Println("Shutting down")

}
