package state_test

import (
	"github.com/awesomegroupidunno/game-server/state"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindVehicle(t *testing.T) {
	Convey("Find Vehicle", t, func() {
		a := state.NewGameState()
		So(a.Vehicles, ShouldBeEmpty)
		a.Vehicles = append(a.Vehicles, state.Vehicle{Endurance: 1, Owner: "me"})
		a.Vehicles = append(a.Vehicles, state.Vehicle{Endurance: 2, Owner: "you"})
		a.Vehicles = append(a.Vehicles, state.Vehicle{Endurance: 3, Owner: "austin"})
		a.Vehicles = append(a.Vehicles, state.Vehicle{Endurance: 4, Owner: "abc"})
		So(a.GetVehicle("me").Endurance, ShouldEqual, 1)
		So(a.GetVehicle("you").Endurance, ShouldEqual, 2)
		So(a.GetVehicle("austin").Endurance, ShouldEqual, 3)
		So(a.GetVehicle("abc").Endurance, ShouldEqual, 4)
		So(a.GetVehicle("gz"), ShouldEqual, nil)

	})
}
