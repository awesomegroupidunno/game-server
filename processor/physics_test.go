package processor

import (
	"github.com/awesomegroupidunno/game-server/state"
	. "github.com/smartystreets/goconvey/convey"
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
