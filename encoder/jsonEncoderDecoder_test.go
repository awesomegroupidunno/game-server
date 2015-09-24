package encoder_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/awesomegroupidunno/game-server/encoder"
)

func TestInstantiation(t *testing.T) {
	Convey("Instantiation", t, func() {
		a := encoder.JsonEncoderDecoder{Tag: "test"}
		So(a.Tag, ShouldEqual, "test")
	})
}