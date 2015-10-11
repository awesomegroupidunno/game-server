package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
)

type CommandProcessor interface {
	Run(g *state.GameState, c cmd.GameCommand)
}

type CommandProcessorFactory struct {
}

func (f *CommandProcessorFactory) GetCommandProcessor(c cmd.GameCommand) CommandProcessor {

	switch c.Command().Subtype {
	case cmd.Turn:
		return &TurnCommandProcessor{}
	case cmd.Acceleration:
		return &AccelerationCommandProcessor{}
	case cmd.Connect:
		return &ConnectCommandProcessor{}
	}
	return &TurnCommandProcessor{}
}
