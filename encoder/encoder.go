package encoder

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
)

type EncoderDecoder interface {
	Encode(state state.GameState) ([]byte, error)
	Decode(b []byte) (cmd.Command, error)
}
