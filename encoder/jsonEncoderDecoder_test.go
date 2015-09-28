package encoder_test

import (
	"encoding/json"
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/encoder"
	"github.com/awesomegroupidunno/game-server/state"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInstantiation(t *testing.T) {
	Convey("Instantiation", t, func() {
		a := encoder.JsonEncoderDecoder{Tag: "test"}
		So(a.Tag, ShouldEqual, "test")
	})
}

func TestDecode(t *testing.T) {
	Convey("Decode", t, func() {
		formatter := encoder.JsonEncoderDecoder{Tag: "DecodeTest"}

		data := cmd.Command{Type: "GET", Subtype: "STATE", UniqueId: "ABC123"}
		buffer, error := json.Marshal(data)

		command, error := formatter.Decode(buffer)

		So(error, ShouldEqual, nil)
		So(command.Type, ShouldEqual, "GET")
		So(command.Subtype, ShouldEqual, "STATE")
		So(command.UniqueId, ShouldEqual, "ABC123")
	})
}

func TestEncode(t *testing.T) {
	Convey("Encode", t, func() {
		formatter := encoder.JsonEncoderDecoder{Tag: "EncodeTest"}
		data := state.GameState{Val: "test Val"}

		decoded := state.GameState{}

		buffer, error := formatter.Encode(data)

		unmarshal_error := json.Unmarshal(buffer, &decoded)

		So(error, ShouldEqual, nil)
		So(unmarshal_error, ShouldEqual, nil)
		So(decoded.Val, ShouldEqual, "test Val")

	})
}
