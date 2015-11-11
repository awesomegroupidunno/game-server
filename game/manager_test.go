package game

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/processor"
	"github.com/awesomegroupidunno/game-server/state"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestControl(t *testing.T) {
	physics := processor.DefaultPhysics()
	current_state := physics.NewGameState()
	response_channel := make(chan state.StateResponse, 100)

	physics.TurnCommandModifier = 1.0
	physics.AccelerationCommandModifier = 1.0
	factory := processor.NewFactory(&physics)

	manager := NewManager(current_state, response_channel, &factory)
	go manager.Start()

	Convey("Pause", t, func() {
		manager.Pause()
		time.Sleep(5 * time.Millisecond)
		So(manager.isPaused, ShouldEqual, true)
	})

	Convey("Resume", t, func() {
		manager.Resume()
		time.Sleep(5 * time.Millisecond)
		So(manager.isPaused, ShouldEqual, false)
	})

}

func TestCommandConsumption(t *testing.T) {
	physics := processor.DefaultPhysics()
	current_state := physics.NewGameState()
	response_channel := make(chan state.StateResponse, 100)

	physics.TurnCommandModifier = 1.0
	physics.AccelerationCommandModifier = 1.0
	factory := processor.NewFactory(&physics)

	manager := NewManager(current_state, response_channel, &factory)

	Convey("Pause", t, func() {
		manager.Pause()
		So(manager.isPaused, ShouldEqual, true)
		time.Sleep(5 * time.Millisecond)
	})

	Convey("AddCommand", t, func() {
		So(len(manager.commandsToProcess), ShouldEqual, 0)
		connect := cmd.NewConnect("tester")
		gamecommand := cmd.GameCommand(&connect)
		manager.AddCommand(gamecommand)
		So(len(manager.commandsToProcess), ShouldEqual, 1)
	})

	Convey("Removed On tick", t, func() {
		So(len(manager.commandsToProcess), ShouldEqual, 1)
		manager.tick()
		So(len(manager.commandsToProcess), ShouldEqual, 0)
	})

}
