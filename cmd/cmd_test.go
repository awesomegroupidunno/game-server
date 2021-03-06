package cmd_test

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBase(t *testing.T) {
	Convey("Base Command", t, func() {
		base := cmd.BaseCommand{Type: "type", Subtype: "subtype", UniqueId: "uniqueid"}
		base.Command().UserId = "theUser"
		So(base.Command().UserId, ShouldEqual, "theUser")
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

func TestConnect(t *testing.T) {
	Convey("Connect Command", t, func() {
		connect := cmd.NewConnect("myname")
		So(connect.UniqueId, ShouldEqual, "")
		So(connect.Subtype, ShouldEqual, cmd.Connect)
		So(connect.Type, ShouldEqual, cmd.Post)
		So(connect.Value, ShouldEqual, "myname")
		So(connect.Type, ShouldEqual, connect.Command().Type)
		So(connect.Subtype, ShouldEqual, connect.Command().Subtype)
		So(connect.UniqueId, ShouldEqual, connect.Command().UniqueId)
	})
}

func TestFire(t *testing.T) {
	Convey("Fire Command", t, func() {
		fire := cmd.NewFire()
		So(fire.UniqueId, ShouldEqual, "")
		So(fire.Subtype, ShouldEqual, cmd.Fire)
		So(fire.Type, ShouldEqual, cmd.Post)
		So(fire.Type, ShouldEqual, fire.Command().Type)
		So(fire.Subtype, ShouldEqual, fire.Command().Subtype)
		So(fire.UniqueId, ShouldEqual, fire.Command().UniqueId)
	})
}

func TestPowerup(t *testing.T) {
	Convey("powerup Command", t, func() {
		fire := cmd.NewPowerup()
		So(fire.UniqueId, ShouldEqual, "")
		So(fire.Subtype, ShouldEqual, cmd.POWERUP)
		So(fire.Type, ShouldEqual, cmd.Post)
		So(fire.Type, ShouldEqual, fire.Command().Type)
		So(fire.Subtype, ShouldEqual, fire.Command().Subtype)
		So(fire.UniqueId, ShouldEqual, fire.Command().UniqueId)
	})
}
