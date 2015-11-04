package encoder

import (
	"encoding/json"
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
)

// Implements EncoderDecoder
type JsonEncoderDecoder struct {
	Tag string
}

// Converts a GameState to a byte array
// byte array is actually a formatted JSON string
func (j *JsonEncoderDecoder) Encode(state state.GameState) ([]byte, error) {
	val, err := json.Marshal(state)
	return val, err
}

// Converts a byte array to a command
// also converts the command to the proper implementation of the command when possible
// byte array is actually a formatted JSON string
func (j *JsonEncoderDecoder) Decode(b []byte) (cmd.GameCommand, error) {
	c := cmd.BaseCommand{}

	error := json.Unmarshal(b, &c)

	if c.Type == cmd.Post {
		switch c.Subtype {
		case cmd.Turn:
			s := cmd.TurnCommand{}
			e := json.Unmarshal(b, &s)
			return &s, e
		case cmd.Acceleration:
			s := cmd.AccelerationCommand{}
			e := json.Unmarshal(b, &s)
			return &s, e
		case cmd.Connect:
			s := cmd.ConnectCommand{}
			e := json.Unmarshal(b, &s)
			return &s, e
		case cmd.Fire:
			s := cmd.FireCommand{}
			e := json.Unmarshal(b, &s)
			return &s, e
		}

	}
	return &c, error
}
