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
		a.Vehicles = make([]*state.Vehicle, 4)
		a.Vehicles[0] = &(state.Vehicle{Endurance: 1, Owner: "me"})
		a.Vehicles[1] = &(state.Vehicle{Endurance: 2, Owner: "you"})
		a.Vehicles[2] = &(state.Vehicle{Endurance: 3, Owner: "austin"})
		a.Vehicles[3] = &(state.Vehicle{Endurance: 4, Owner: "abc"})
		So(a.GetVehicle("me").Endurance, ShouldEqual, 1)
		So(a.GetVehicle("you").Endurance, ShouldEqual, 2)
		So(a.GetVehicle("austin").Endurance, ShouldEqual, 3)
		So(a.GetVehicle("abc").Endurance, ShouldEqual, 4)
		So(a.GetVehicle("gz"), ShouldEqual, nil)

	})
}

func TestStateCopy(t *testing.T) {
	Convey("Copy State", t, func() {
		a := state.NewGameState()
		theCopy := a.Copy()
		So(len(theCopy.Vehicles), ShouldEqual, len(a.Vehicles))
		So(len(theCopy.Bases), ShouldEqual, len(a.Bases))
		So(len(theCopy.ShieldGenerators), ShouldEqual, len(a.ShieldGenerators))
		So(theCopy.GameOver, ShouldEqual, a.GameOver)
	})
}

func TestBox2d(t *testing.T) {
	vehicle := state.Vehicle{X: 10,
		Y:      20,
		Width:  15,
		Height: 40,
		Angle:  3}
	Convey("Proper Box", t, func() {
		So(vehicle.AngleDegrees(), ShouldAlmostEqual, vehicle.Angle, .001)
		x, y := vehicle.Position()
		w, h := vehicle.Size()

		So(x, ShouldAlmostEqual, vehicle.X, .001)
		So(y, ShouldAlmostEqual, vehicle.Y, .001)
		So(w, ShouldAlmostEqual, vehicle.Width, .001)
		So(h, ShouldAlmostEqual, vehicle.Height, .001)

	})
}
