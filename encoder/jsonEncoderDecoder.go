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

func (j *JsonEncoderDecoder) Decode(b []byte) (cmd.Command, error) {
	return cmd.Command{Type:"GET"}, nil
}
