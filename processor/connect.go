package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
)

type ConnectCommandProcessor struct {
}

func (t *ConnectCommandProcessor) Run(g *state.GameState, c cmd.GameCommand) {
	command := c.(*cmd.ConnectCommand)
	log.Println(command)

}
