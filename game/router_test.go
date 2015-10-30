package game

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/processor"
	"github.com/awesomegroupidunno/game-server/state"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"testing"
)

func TestRoutePost(t *testing.T) {
	ack_channel := make(chan state.Ack, 100)
	state_channel := make(chan state.StateResponse, 100)

	physics := processor.Physics{}
	factory := processor.CommandProcessorFactory{Physics: &physics}

	gameManager := NewManager(state.NewGameState(), state_channel, &factory)

	router := CommandRouter{Acks: ack_channel, GameManager: &gameManager}
	address, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")

	Convey("Route Post", t, func() {
		So(err, ShouldEqual, nil)
		turn := cmd.NewTurn(.4)
		cmd := cmd.GameCommand(&turn)

		So(len(gameManager.commandsToProcess), ShouldEqual, 0)

		router.routePost(&cmd, address)
		resp := <-ack_channel

		So(resp.Address, ShouldEqual, address)
		So(len(gameManager.commandsToProcess), ShouldEqual, 1)
	})
}

func TestRoute(t *testing.T) {
	ack_channel := make(chan state.Ack, 100)
	state_channel := make(chan state.StateResponse, 100)

	physics := processor.Physics{}
	factory := processor.CommandProcessorFactory{Physics: &physics}

	gameManager := NewManager(state.NewGameState(), state_channel, &factory)

	router := CommandRouter{Acks: ack_channel, GameManager: &gameManager}
	address, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")

	Convey("Route", t, func() {
		So(err, ShouldEqual, nil)
		turn := cmd.NewConnect("test")
		cmd := cmd.GameCommand(&turn)

		So(len(gameManager.commandsToProcess), ShouldEqual, 0)

		router.RouteCommand(&cmd, address)
		resp := <-ack_channel

		So(resp.Address, ShouldEqual, address)
		So(len(gameManager.commandsToProcess), ShouldEqual, 1)
	})
}
