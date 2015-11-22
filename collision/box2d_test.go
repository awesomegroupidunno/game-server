package collision_test

import (
	"github.com/awesomegroupidunno/game-server/collision"
	"github.com/awesomegroupidunno/game-server/state"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCollisions(t *testing.T) {
	Convey("Basic No Collision", t, func() {
		v1 := state.Vehicle{
			Point: state.NewPoint(10, 10),
			Sized: state.NewSized(10, 10),
			Angle: 0}
		v2 := state.Vehicle{
			Point: state.NewPoint(30, 30),
			Sized: state.NewSized(10, 10),
			Angle: 0}

		check := collision.Collides(v1, v2)

		So(check, ShouldEqual, false)
	})

	Convey("Basic Should Collide", t, func() {
		v1 := state.Vehicle{
			Point: state.NewPoint(10, 10),
			Sized: state.NewSized(10, 10),
			Angle: 0}

		v2 := state.Vehicle{
			Point: state.NewPoint(5, 5),
			Sized: state.NewSized(10, 10),
			Angle: 0}
		check := collision.Collides(v1, v2)

		So(check, ShouldEqual, true)
	})

	Convey("Basic Should Collide", t, func() {
		v1 := state.Vehicle{
			Point: state.NewPoint(10, 10),
			Sized: state.NewSized(10, 10),
			Angle: 0}
		v2 := state.Vehicle{
			Point: state.NewPoint(7, 7),
			Sized: state.NewSized(10, 10),
			Angle: 0}
		check := collision.Collides(v1, v2)

		So(check, ShouldEqual, true)
	})

	Convey("Basic Should Collide with Rotation", t, func() {
		v1 := state.Vehicle{
			Point: state.NewPoint(10, 10),
			Sized: state.NewSized(10, 10),
			Angle: 90}
		v2 := state.Vehicle{
			Point: state.NewPoint(7, 7),
			Sized: state.NewSized(10, 10),
			Angle: 0}
		check := collision.Collides(v1, v2)

		So(check, ShouldEqual, true)
	})

	Convey("Basic Should Collide with Rotation", t, func() {
		v1 := state.Vehicle{
			Point: state.NewPoint(10, 10),
			Sized: state.NewSized(10, 10),
			Angle: 90}
		v2 := state.Vehicle{
			Point: state.NewPoint(7, 7),
			Sized: state.NewSized(10, 10),
			Angle: 90}
		check := collision.Collides(v1, v2)

		So(check, ShouldEqual, true)
	})
}
