package processor

import (
	"github.com/awesomegroupidunno/game-server/state"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
	"time"
)

func TestMove(t *testing.T) {

	physics := DefaultPhysics()

	Convey("Move Right", t, func() {

		before_x := 10.0
		before_y := 10.0
		before_angle := 0.0
		before_speed := 1.0
		x, y := physics.move2d(before_x, before_y, before_angle, before_speed, time.Second)
		So(y, ShouldAlmostEqual, before_y, 0.001)
		So(x, ShouldAlmostEqual, before_y+1, 0.001)

	})

	Convey("Move Left", t, func() {

		before_x := 10.0
		before_y := 10.0
		before_angle := 0.0
		before_speed := -1.0
		x, y := physics.move2d(before_x, before_y, before_angle, before_speed, time.Second)
		So(y, ShouldAlmostEqual, before_y, 0.001)
		So(x, ShouldAlmostEqual, before_x-1, 0.001)

	})

	Convey("Move Up", t, func() {

		before_x := 10.0
		before_y := 10.0
		before_angle := 90.0
		before_speed := 1.0
		x, y := physics.move2d(before_x, before_y, before_angle, before_speed, time.Second)
		So(y, ShouldAlmostEqual, before_y+1, 0.001)
		So(x, ShouldAlmostEqual, before_x, 0.001)

	})

	Convey("Move Down", t, func() {

		before_x := 10.0
		before_y := 10.0
		before_angle := 270.0
		before_speed := 1.0
		x, y := physics.move2d(before_x, before_y, before_angle, before_speed, time.Second)
		So(y, ShouldAlmostEqual, before_y-1, 0.001)
		So(x, ShouldAlmostEqual, before_x, 0.001)

	})
}

func TestMoveVehicle(t *testing.T) {

	physics := DefaultPhysics()

	Convey("Move Right", t, func() {
		before_x := 10.0
		before_y := 10.0
		before_angle := 270.0
		before_speed := 1.0

		vehicle := state.Vehicle{X: before_x,
			Y:        before_y,
			Angle:    before_angle,
			Velocity: before_speed}

		physics.MoveVehicle(&vehicle, time.Second)

		So(vehicle.Y, ShouldAlmostEqual, before_y-1, .001)
		So(vehicle.X, ShouldAlmostEqual, before_x, .001)
	})
}

func TestSlowVehicle(t *testing.T) {

	physics := DefaultPhysics()

	Convey("Slow to stop Vehicle", t, func() {
		before_x := 10.0
		before_y := 10.0
		before_angle := 270.0
		before_speed := 1.0

		vehicle := state.Vehicle{X: before_x,
			Y:        before_y,
			Angle:    before_angle,
			Velocity: before_speed}

		physics.VehicleFrictionSlow(&vehicle, time.Second)

		So(vehicle.Velocity, ShouldAlmostEqual, 0, .001)

	})

	Convey("Slow to stop Vehicle, reverse", t, func() {
		before_x := 10.0
		before_y := 10.0
		before_angle := 270.0
		before_speed := -1.0

		vehicle := state.Vehicle{X: before_x,
			Y:        before_y,
			Angle:    before_angle,
			Velocity: before_speed}

		physics.VehicleFrictionSlow(&vehicle, time.Second)

		So(vehicle.Velocity, ShouldAlmostEqual, 0, .001)

	})
}

func TestCombineComponents(t *testing.T) {

	Convey("Simple Components", t, func() {
		a := combineComponents(4, 3)
		So(a, ShouldAlmostEqual, 5, .001)
	})

	Convey("Complex Components", t, func() {
		a := combineComponents(4, 4)
		So(a, ShouldAlmostEqual, math.Sqrt(32), .001)
	})

	Convey("Complex Components", t, func() {
		a := combineComponents(123.5, 4.99)
		So(a, ShouldAlmostEqual, math.Sqrt(math.Pow(123.5, 2)+math.Pow(4.99, 2)), .001)
	})
}

func TestVehicleBounding(t *testing.T) {
	physics := DefaultPhysics()

	Convey("Should Be in Bounds", t, func() {

		v := state.Vehicle{X: 20, Y: 30}
		physics.VehicleBounding(&v)

		So(v.X, ShouldAlmostEqual, 20, .001)
		So(v.Y, ShouldAlmostEqual, 30, .001)
	})

	Convey("Should Be reBound from negative", t, func() {

		v := state.Vehicle{X: -20, Y: -30}
		physics.VehicleBounding(&v)

		So(v.X, ShouldBeGreaterThan, 0)
		So(v.Y, ShouldBeGreaterThan, 0)
	})

	Convey("Should Be reBound from max", t, func() {

		v := state.Vehicle{X: math.MaxFloat64, Y: math.MaxFloat64}
		physics.VehicleBounding(&v)

		So(v.X, ShouldBeLessThan, physics.WorldWidth)
		So(v.Y, ShouldBeLessThan, physics.WorldHeight)
	})
}

func TestMoveBullet(t *testing.T) {

	physics := DefaultPhysics()

	Convey("Move Right", t, func() {
		before_x := 10.0
		before_y := 10.0
		before_angle := 270.0
		before_speed := 1.0

		bullet := state.Bullet{X: before_x,
			Y:        before_y,
			Angle:    before_angle,
			Velocity: before_speed}

		physics.MoveBullet(&bullet, time.Second)

		So(bullet.Y, ShouldAlmostEqual, before_y-1, .001)
		So(bullet.X, ShouldAlmostEqual, before_x, .001)
	})
}

func TestCleanupBullet(t *testing.T) {
	physics := DefaultPhysics()
	bullets := []*state.Bullet{}

	bullets = append(bullets, &state.Bullet{X: -10, Y: -30})
	bullets = append(bullets, &state.Bullet{X: 10, Y: 30})
	bullets = append(bullets, &state.Bullet{X: 50, Y: 80})
	bullets = append(bullets, &state.Bullet{X: 70, Y: -30})
	bullets = append(bullets, &state.Bullet{X: 50, Y: 90})

	Convey("Test Cleanup offscreen bullets", t, func() {
		So(len(bullets), ShouldEqual, 5)
		newList := physics.CleanUpBullets(bullets)

		So(len(newList), ShouldEqual, 3)
	})

}
