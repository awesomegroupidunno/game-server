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
		data := cmd.BaseCommand{Type: "GET", Subtype: "STATE", UniqueId: "ABC123"}
		buffer, error := json.Marshal(data)

		command, error := formatter.Decode(buffer)

		So(error, ShouldEqual, nil)
		So(command.Command().Type, ShouldEqual, "GET")
		So(command.Command().Subtype, ShouldEqual, "STATE")
		So(command.Command().UniqueId, ShouldEqual, "ABC123")
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

func TestDecodeAcceleration(t *testing.T) {
	Convey("Acceleration Decode", t, func() {
		formatter := encoder.JsonEncoderDecoder{Tag: "DecodeTest"}

		data := cmd.NewAcceleration(.2)
		data.UniqueId = "ABC123"
		buffer, error := json.Marshal(data)

		command, err := formatter.Decode(buffer)

		acceleration := command.(*cmd.AccelerationCommand)

		So(error, ShouldEqual, nil)
		So(acceleration.Command().Type, ShouldEqual, cmd.Post)
		So(acceleration.Command().Subtype, ShouldEqual, cmd.Acceleration)
		So(err, ShouldEqual, nil)
		So(acceleration.Value, ShouldAlmostEqual, .2, .000001)
		So(acceleration.Command().UniqueId, ShouldEqual, "ABC123")
	})
}

func TestDecodeTurn(t *testing.T) {
	Convey("Turn Decode", t, func() {
		formatter := encoder.JsonEncoderDecoder{Tag: "DecodeTest"}

		data := cmd.NewTurn(.2)
		data.UniqueId = "ABC123"
		buffer, error := json.Marshal(data)

		command, err := formatter.Decode(buffer)

		turn := command.(*cmd.TurnCommand)

		So(error, ShouldEqual, nil)
		So(turn.Command().Type, ShouldEqual, cmd.Post)
		So(turn.Command().Subtype, ShouldEqual, cmd.Turn)
		So(err, ShouldEqual, nil)
		So(turn.Value, ShouldAlmostEqual, .2, .000001)
		So(turn.Command().UniqueId, ShouldEqual, "ABC123")
	})
}
