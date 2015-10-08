package encoder

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
)

type EncoderDecoder interface {
	// Converts a GameState to a byte array
	Encode(state state.GameState) ([]byte, error)

	// Converts a byte array to a command
	// also converts the command to the proper implementation of the command when possible
	Decode(b []byte) (cmd.GameCommand, error)
}
