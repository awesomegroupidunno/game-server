package net

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBasic(t *testing.T) {
	Convey("Basic Test", t, func() {
		a:=UdpReceiver{PortNumber:"1234"}
		So(a.PortNumber, ShouldEqual, "1234")

	})
}
