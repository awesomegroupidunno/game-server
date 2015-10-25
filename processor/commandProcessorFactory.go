package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
)

type CommandProcessor interface {
	Run(g *state.GameState, c cmd.GameCommand)
}

type CommandProcessorFactory struct {
	Physics  *Physics
	turnProc *TurnCommandProcessor
	accProc  *AccelerationCommandProcessor
	conProc  *ConnectCommandProcessor
}

func NewFactory(physics *Physics) CommandProcessorFactory {
	return CommandProcessorFactory{Physics: physics,
		turnProc: &TurnCommandProcessor{Physics: physics},
		accProc:  &AccelerationCommandProcessor{Physics: physics},
		conProc:  &ConnectCommandProcessor{Physics: physics}}
}

// Returns a command processor for the command passed
// if no processor is available nil is returned
func (f *CommandProcessorFactory) GetCommandProcessor(c *cmd.GameCommand) CommandProcessor {

	switch (*c).Command().Subtype {
	case cmd.Turn:
		return f.turnProc
	case cmd.Acceleration:
		return f.accProc
	case cmd.Connect:
		return f.conProc
	}

	return nil
}
