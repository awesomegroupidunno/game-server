package collision_test

import (
	"github.com/awesomegroupidunno/game-server/collision"
	"github.com/awesomegroupidunno/game-server/state"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCollisions(t *testing.T) {
	Convey("Basic No Collision", t, func() {
		v1 := state.Vehicle{X: 10, Y: 10, Width: 10, Height: 10, Angle: 0}
		v2 := state.Vehicle{X: 30, Y: 30, Width: 10, Height: 10, Angle: 0}

		check := collision.Collides(v1, v2)

		So(check, ShouldEqual, false)
	})

	Convey("Basic Should Collide", t, func() {
		v1 := state.Vehicle{X: 10, Y: 10, Width: 10, Height: 10, Angle: 0}
		v2 := state.Vehicle{X: 5, Y: 5, Width: 10, Height: 10, Angle: 0}

		check := collision.Collides(v1, v2)

		So(check, ShouldEqual, true)
	})

	Convey("Basic Should Collide", t, func() {
		v1 := state.Vehicle{X: 10, Y: 10, Width: 10, Height: 10, Angle: 0}
		v2 := state.Vehicle{X: 7, Y: 7, Width: 10, Height: 10, Angle: 0}

		check := collision.Collides(v1, v2)

		So(check, ShouldEqual, true)
	})

	Convey("Basic Should Collide with Rotation", t, func() {
		v1 := state.Vehicle{X: 10, Y: 10, Width: 10, Height: 10, Angle: 90}
		v2 := state.Vehicle{X: 7, Y: 7, Width: 10, Height: 10, Angle: 0}

		check := collision.Collides(v1, v2)

		So(check, ShouldEqual, true)
	})

	Convey("Basic Should Collide with Rotation", t, func() {
		v1 := state.Vehicle{X: 10, Y: 10, Width: 10, Height: 10, Angle: 90}
		v2 := state.Vehicle{X: 7, Y: 7, Width: 10, Height: 10, Angle: 90}

		check := collision.Collides(v1, v2)

		So(check, ShouldEqual, true)
	})
}
