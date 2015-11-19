package network_test

import (
	"bufio"
	"encoding/json"
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/encoder"
	"github.com/awesomegroupidunno/game-server/game"
	"github.com/awesomegroupidunno/game-server/network"
	"github.com/awesomegroupidunno/game-server/processor"
	"github.com/awesomegroupidunno/game-server/state"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {

	ack_channel := make(chan state.Ack, 100)
	state_channel := make(chan state.StateResponse, 5)
	physics := processor.DefaultPhysics()
	factory := processor.NewFactory(&physics)
	manager := game.NewManager(physics.NewGameState(), state_channel, &factory)
	router := game.CommandRouter{GameManager: &manager, Acks: ack_channel}
	go manager.Start()

	reciever := network.UdpReceiver{
		Router:     router,
		Acks:       ack_channel,
		Responses:  state_channel,
		PortNumber: ":10002"}

	reciever.Run()
	conn, err := net.DialTimeout("udp", "127.0.0.1:10002", 1*time.Second)
	Convey("Test connect", t, func() {
		So(err, ShouldEqual, nil)
		So(conn, ShouldNotEqual, nil)
	})

	Convey("Test Connect command", t, func() {
		a := cmd.NewConnect("Dan")
		a.UniqueId = "abc123"
		message, err := json.Marshal(a)

		So(err, ShouldEqual, nil)

		i, error := conn.Write(message)

		So(i, ShouldBeGreaterThan, 0)
		So(error, ShouldEqual, nil)
	})

	Convey("Test State Responder", t, func() {
		reciever.Responses <- state.StateResponse{State: physics.NewGameState()}
		response := make([]byte, 2048)
		// read response
		conn.SetDeadline(time.Now().Add(2 * time.Second))
		n, err := bufio.NewReader(conn).Read(response)
		So(err, ShouldEqual, nil)
		So(string(response[:n]), ShouldEqual, "abc123")

		conn.SetDeadline(time.Now().Add(2 * time.Second))
		n, err = bufio.NewReader(conn).Read(response)
		So(err, ShouldEqual, nil)

		encoder := encoder.JsonEncoderDecoder{}

		_, encode_error := encoder.Decode(response[:n])
		So(encode_error, ShouldEqual, nil)

	})

}
