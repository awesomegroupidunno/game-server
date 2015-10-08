package cmd_test

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBase(t *testing.T) {
	Convey("Base Command", t, func() {
		base := cmd.BaseCommand{Type: "type", Subtype: "subtype", UniqueId: "uniqueid"}
		So(base.UniqueId, ShouldEqual, "uniqueid")
		So(base.Subtype, ShouldEqual, "subtype")
		So(base.Type, ShouldEqual, "type")
		So(base.Type, ShouldEqual, base.Command().Type)
		So(base.Subtype, ShouldEqual, base.Command().Subtype)
		So(base.UniqueId, ShouldEqual, base.Command().UniqueId)
	})
}
func TestAcceleration(t *testing.T) {
	Convey("Acceleration Command", t, func() {
		acceleration := cmd.NewAcceleration(.4)
		So(acceleration.UniqueId, ShouldEqual, "")
		So(acceleration.Subtype, ShouldEqual, cmd.Acceleration)
		So(acceleration.Type, ShouldEqual, cmd.Post)
		So(acceleration.Value, ShouldAlmostEqual, .4, .00001)
		So(acceleration.Type, ShouldEqual, acceleration.Command().Type)
		So(acceleration.Subtype, ShouldEqual, acceleration.Command().Subtype)
		So(acceleration.UniqueId, ShouldEqual, acceleration.Command().UniqueId)
	})
}
func TestTurn(t *testing.T) {
	Convey("Turn Command", t, func() {
		turn := cmd.NewTurn(.3)
		So(turn.UniqueId, ShouldEqual, "")
		So(turn.Subtype, ShouldEqual, cmd.Turn)
		So(turn.Type, ShouldEqual, cmd.Post)
		So(turn.Value, ShouldAlmostEqual, .3, .00001)
		So(turn.Type, ShouldEqual, turn.Command().Type)
		So(turn.Subtype, ShouldEqual, turn.Command().Subtype)
		So(turn.UniqueId, ShouldEqual, turn.Command().UniqueId)
	})
}
