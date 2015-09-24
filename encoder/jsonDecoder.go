package encoder

import (
	"encoding/json"
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
)

type JsonDecoder struct {
}

func (j *JsonDecoder) Encode(state state.GameState) ([]byte, error) {
	val, err := json.Marshal(state)
	return val, err
}

func (j *JsonDecoder) Decode(b []byte) (cmd.Command, error) {
	return make([]byte, 1024)
}
