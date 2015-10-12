package processor_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/awesomegroupidunno/game-server/processor"
	"github.com/awesomegroupidunno/game-server/cmd"
)

func TestProcessorFactory(t *testing.T) {
	Convey("ProcessorFactory", t, func() {
		a := processor.CommandProcessorFactory{}

		t := cmd.NewTurn(1)
		t_comm := cmd.GameCommand(&t)

		ac := cmd.NewAcceleration(1)
		a_comm := cmd.GameCommand(&ac)
		connect := cmd.NewConnect("test")
		c_comm := cmd.GameCommand(&connect)

		turn_cp := a.GetCommandProcessor(&t_comm)
		acceleration_cp := a.GetCommandProcessor(&a_comm)
		connect_cp := a.GetCommandProcessor(&c_comm)

		So(turn_cp, ShouldHaveSameTypeAs, &processor.TurnCommandProcessor{} )
		So(acceleration_cp, ShouldHaveSameTypeAs, &processor.AccelerationCommandProcessor{} )
		So(connect_cp, ShouldHaveSameTypeAs, &processor.ConnectCommandProcessor{} )
	})
}