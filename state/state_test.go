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

		a.Rockets
		So(a.Vehicles, ShouldBeEmpty)
		a.Vehicles = make([]*state.Vehicle, 4)
		a.Vehicles[0] = &(state.Vehicle{
			Point: state.NewPoint(1, 0),
			Owner: "me"})
		a.Vehicles[1] = &(state.Vehicle{
			Point: state.NewPoint(2, 0),
			Owner: "you"})
		a.Vehicles[2] = &(state.Vehicle{
			Point: state.NewPoint(3, 0),
			Owner: "austin"})
		a.Vehicles[3] = &(state.Vehicle{
			Point: state.NewPoint(4, 0),
			Owner: "abc"})
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
		a.Vehicles[0] = &(state.Vehicle{
			Point:   state.NewPoint(1, 0),
			Owner:   "me",
			IsAlive: true})

		a.Vehicles[1] = &(state.Vehicle{
			Point:   state.NewPoint(1, 0),
			Owner:   "you",
			IsAlive: true})

		a.Vehicles[2] = &(state.Vehicle{
			Point:   state.NewPoint(1, 0),
			Owner:   "austin",
			IsAlive: true})
		a.Vehicles[3] = &(state.Vehicle{
			Point:   state.NewPoint(1, 0),
			Owner:   "abc",
			IsAlive: true})

		a.Bullets = make([]*state.Bullet, 4)
		a.Bullets[0] = &(state.Bullet{})
		a.Bullets[1] = &(state.Bullet{})
		a.Bullets[2] = &(state.Bullet{})
		a.Bullets[3] = &(state.Bullet{})

		theCopy := a.Copy()
		So(len(theCopy.Vehicles), ShouldEqual, len(a.Vehicles))
		So(len(theCopy.Bases), ShouldEqual, len(a.Bases))
		So(len(theCopy.ShieldGenerators), ShouldEqual, len(a.ShieldGenerators))
		So(theCopy.GameOver, ShouldEqual, a.GameOver)
	})
}

func TestBox2dVehicle(t *testing.T) {
	vehicle := state.Vehicle{
		Point: state.NewPoint(10, 20),
		Sized: state.NewSized(15, 40),
		Angle: 3}
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
	bullet := state.Bullet{
		Point: state.Point{X: 10,
			Y: 20},
		Sized: state.Sized{Width: 15,
			Height: 40},
		Angle: 3}
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
