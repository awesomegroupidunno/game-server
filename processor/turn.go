package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
)

type TurnCommandProcessor struct {
}

func (t *TurnCommandProcessor) Run(g *state.GameState, c cmd.GameCommand) {
	command := c.(*cmd.TurnCommand)
	log.Println(command)

}
