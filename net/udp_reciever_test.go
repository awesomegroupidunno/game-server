package net

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBasic(t *testing.T) {
	Convey("Basic Test", t, func() {

		So(3, ShouldEqual, 3)

	})
}
