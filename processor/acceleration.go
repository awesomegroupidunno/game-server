package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
)

type AccelerationCommandProcessor struct {
}

func (t *AccelerationCommandProcessor) Run(g *state.GameState, c cmd.GameCommand) {
	command := c.(*cmd.AccelerationCommand)
	log.Println(command)

}
