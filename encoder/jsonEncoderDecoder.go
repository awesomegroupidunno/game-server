package encoder

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(string(b))

	if c.Type == "GET" {
		if c.Subtype == "STATE" {
			s := cmd.StateCommand{}
			e := json.Unmarshal(b, &s)
			return &s, e
		}
	}

	if c.Type == "POST" {

		if c.Subtype == "TURN" {
			s := cmd.TurnCommand{}
			e := json.Unmarshal(b, &s)
			return &s, e
		}

		if c.Subtype == "ACCELERATION" {
			s := cmd.AccelerationCommand{}
			e := json.Unmarshal(b, &s)
			return &s, e
		}
	}
	return &c, error
}
