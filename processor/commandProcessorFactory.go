package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
)

type CommandProcessor interface {
	Run(g *state.GameState, c cmd.GameCommand)
}

type CommandProcessorFactory struct {
	Physics *Physics
}

// Returns a command processor for the command passed
// if no processor is available nil is returned
func (f *CommandProcessorFactory) GetCommandProcessor(c *cmd.GameCommand) CommandProcessor {

	switch (*c).Command().Subtype {
	case cmd.Turn:
		return &TurnCommandProcessor{Physics: f.Physics}
	case cmd.Acceleration:
		return &AccelerationCommandProcessor{Physics: f.Physics}
	case cmd.Connect:
		return &ConnectCommandProcessor{Physics: f.Physics}
	}

	return nil
}
