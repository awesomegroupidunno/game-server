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

	if c.Type == "GET" {
		switch c.Subtype {
		case "STATE":
			s := cmd.StateCommand{}
			e := json.Unmarshal(b, &s)
			return &s, e
		}
	}

	if c.Type == "POST" {
		switch c.Subtype {
		case "TURN":
			s := cmd.TurnCommand{}
			e := json.Unmarshal(b, &s)
			return &s, e
		case "ACCELERATION":
			s := cmd.AccelerationCommand{}
			e := json.Unmarshal(b, &s)
			return &s, e
		}

	}
	return &c, error
}
