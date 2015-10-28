package network_test

import (
	"bufio"
	"encoding/json"
	"github.com/awesomegroupidunno/game-server/cmd"
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
	state_channel := make(chan state.StateResponse)
	physics := processor.DefaultPhysics()
	factory := processor.NewFactory(&physics)
	manager := game.NewManager(state.NewGameState(), state_channel, &factory)
	router := game.CommandRouter{GameManager: &manager}

	reciever := network.UdpReceiver{
		Router:    router,
		Acks:      ack_channel,
		Responses: state_channel}

	reciever.Run()
	conn, err := net.DialTimeout("udp", "127.0.0.1:10001", 1*time.Second)
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
		reciever.Responses <- state.StateResponse{State: state.NewGameState()}
		response := make([]byte, 2048)

		// read response
		n, err := bufio.NewReader(conn).Read(response)
		So(err, ShouldEqual, nil)

		resp := state.StateResponse{}

		unmarshal_error := json.Unmarshal(response[:n], &resp)
		So(unmarshal_error, ShouldEqual, nil)

	})

}
