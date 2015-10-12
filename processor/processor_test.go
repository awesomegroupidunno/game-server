package processor_test

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/processor"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/awesomegroupidunno/game-server/state"
)

func TestProcessorFactory(t *testing.T) {
	Convey("ProcessorFactory", t, func() {
		factory := processor.CommandProcessorFactory{}

		t := cmd.NewTurn(1)
		t_comm := cmd.GameCommand(&t)
		turn_cp := factory.GetCommandProcessor(&t_comm)
		So(turn_cp, ShouldHaveSameTypeAs, &processor.TurnCommandProcessor{})

		a := cmd.NewAcceleration(1)
		a_comm := cmd.GameCommand(&a)
		acceleration_cp := factory.GetCommandProcessor(&a_comm)
		So(acceleration_cp, ShouldHaveSameTypeAs, &processor.AccelerationCommandProcessor{})

		c := cmd.NewConnect("test")
		c_comm := cmd.GameCommand(&c)
		connect_cp := factory.GetCommandProcessor(&c_comm)
		So(connect_cp, ShouldHaveSameTypeAs, &processor.ConnectCommandProcessor{})

		b := cmd.BaseCommand{}
		b_comm := cmd.GameCommand(&b)
		base_cp := factory.GetCommandProcessor(&b_comm)
		So(base_cp, ShouldEqual, nil)
	})
}

func TestConnectionProcessor(t *testing.T) {
	Convey("Connection Processor", t, func() {
		conn_processor := processor.ConnectCommandProcessor{}

		c := cmd.NewConnect("abc123")
		c2 := cmd.NewConnect("abc")
		c.UserId = "abc123"
		c2.UserId = "abc"

		connect := cmd.GameCommand(&c)
		connect2 := cmd.GameCommand(&c2)

		game_state := state.NewGameState()

		So(len(game_state.Vehicles), ShouldEqual, 0)

		conn_processor.Run(&game_state, connect)
		So(len(game_state.Vehicles), ShouldEqual, 1)

		conn_processor.Run(&game_state, connect)
		So(len(game_state.Vehicles), ShouldEqual, 1)

		conn_processor.Run(&game_state, connect2)
		So(len(game_state.Vehicles), ShouldEqual, 2)
	})
}

func TestAccelerationProcessor(t *testing.T) {
	Convey("Acceleration Processor", t, func() {
		conn_processor := processor.AccelerationCommandProcessor{}

		c := cmd.NewAcceleration(.5)
		c2 := cmd.NewAcceleration(.2)
		c.UserId = "abc123"
		c2.UserId = "abc123"

		accelerate := cmd.GameCommand(&c)
		accelerate2 := cmd.GameCommand(&c2)

		game_state := state.NewGameState()

		game_state.Vehicles = append(game_state.Vehicles, &(state.Vehicle{Owner:"abc123", Velocity:0}))

		So(len(game_state.Vehicles), ShouldEqual, 1)
		So(game_state.Vehicles[0].Velocity, ShouldAlmostEqual, 0, .0001)

		conn_processor.Run(&game_state, accelerate)
		So(game_state.Vehicles[0].Velocity, ShouldAlmostEqual, .5, .0001)

		conn_processor.Run(&game_state, accelerate2)
		So(game_state.Vehicles[0].Velocity, ShouldAlmostEqual, .7, .0001)

		accelerate2.Command().UserId = "blach"
		conn_processor.Run(&game_state, accelerate2)
		So(game_state.Vehicles[0].Velocity, ShouldAlmostEqual, .7, .0001)

	})
}
