package encoder

import (
	"encoding/json"
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
)

type JsonEncoderDecoder struct {
	Tag string
}

func (j *JsonEncoderDecoder) Encode(state state.GameState) ([]byte, error) {
	val, err := json.Marshal(state)
	return val, err
}

func (j *JsonEncoderDecoder) Decode(b []byte) (cmd.GameCommand, error) {
	c := cmd.BaseCommand{}

	error := json.Unmarshal(b, &c)

	if c.Type == cmd.Get {
		switch c.Subtype {
		case cmd.State:
			s := cmd.StateCommand{}
			e := json.Unmarshal(b, &s)
			return &s, e
		}
	}

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
		}

	}
	return &c, error
}
