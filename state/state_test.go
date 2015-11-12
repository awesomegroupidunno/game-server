package state_test

import (
	"github.com/awesomegroupidunno/game-server/processor"
	"github.com/awesomegroupidunno/game-server/state"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindVehicle(t *testing.T) {

	physics := processor.DefaultPhysics()
	Convey("Find Vehicle", t, func() {
		a := physics.NewGameState()
		So(a.Vehicles, ShouldBeEmpty)
		a.Vehicles = make([]*state.Vehicle, 4)
		a.Vehicles[0] = &(state.Vehicle{X: 1, Owner: "me"})
		a.Vehicles[1] = &(state.Vehicle{X: 2, Owner: "you"})
		a.Vehicles[2] = &(state.Vehicle{X: 3, Owner: "austin"})
		a.Vehicles[3] = &(state.Vehicle{X: 4, Owner: "abc"})
		So(a.GetVehicle("me").X, ShouldEqual, 1)
		So(a.GetVehicle("you").X, ShouldEqual, 2)
		So(a.GetVehicle("austin").X, ShouldEqual, 3)
		So(a.GetVehicle("abc").X, ShouldEqual, 4)
		So(a.GetVehicle("gz"), ShouldEqual, nil)

	})
}

func TestStateCopy(t *testing.T) {
	physics := processor.DefaultPhysics()
	Convey("Copy State", t, func() {
		a := physics.NewGameState()

		a.Vehicles = make([]*state.Vehicle, 4)
		a.Vehicles[0] = &(state.Vehicle{X: 1, Owner: "me"})
		a.Vehicles[1] = &(state.Vehicle{X: 2, Owner: "you"})
		a.Vehicles[2] = &(state.Vehicle{X: 3, Owner: "austin"})
		a.Vehicles[3] = &(state.Vehicle{X: 4, Owner: "abc"})

		a.Bullets = make([]*state.Bullet, 4)
		a.Bullets[0] = &(state.Bullet{X: 1, OwnerId: "me"})
		a.Bullets[1] = &(state.Bullet{X: 2, OwnerId: "you"})
		a.Bullets[2] = &(state.Bullet{X: 3, OwnerId: "austin"})
		a.Bullets[3] = &(state.Bullet{X: 4, OwnerId: "abc"})

		theCopy := a.Copy()
		So(len(theCopy.Vehicles), ShouldEqual, len(a.Vehicles))
		So(len(theCopy.Bases), ShouldEqual, len(a.Bases))
		So(len(theCopy.ShieldGenerators), ShouldEqual, len(a.ShieldGenerators))
		So(theCopy.GameOver, ShouldEqual, a.GameOver)
	})
}

func TestBox2dVehicle(t *testing.T) {
	vehicle := state.Vehicle{X: 10,
		Y:      20,
		Width:  15,
		Height: 40,
		Angle:  3}
	Convey("Proper Box Vehicle", t, func() {
		So(vehicle.AngleDegrees(), ShouldAlmostEqual, vehicle.Angle, .001)
		x, y := vehicle.Position()
		w, h := vehicle.Size()

		So(x, ShouldAlmostEqual, vehicle.X, .001)
		So(y, ShouldAlmostEqual, vehicle.Y, .001)
		So(w, ShouldAlmostEqual, vehicle.Height, .001)
		So(h, ShouldAlmostEqual, vehicle.Width, .001)

	})
}

func TestBox2dBullet(t *testing.T) {
	bullet := state.Bullet{X: 10,
		Y:      20,
		Width:  15,
		Height: 40,
		Angle:  3}
	Convey("Proper Box Bullet", t, func() {
		So(bullet.AngleDegrees(), ShouldAlmostEqual, bullet.Angle, .001)
		x, y := bullet.Position()
		w, h := bullet.Size()

		So(x, ShouldAlmostEqual, bullet.X, .001)
		So(y, ShouldAlmostEqual, bullet.Y, .001)
		So(w, ShouldAlmostEqual, bullet.Height, .001)
		So(h, ShouldAlmostEqual, bullet.Width, .001)

	})
}
