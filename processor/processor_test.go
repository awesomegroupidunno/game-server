package processor_test

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/processor"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
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
