package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
)

type FireCommandProcessor struct {
	Physics *Physics
}

func (t *FireCommandProcessor) Run(g *state.GameState, c cmd.GameCommand) {
	command := c.(*cmd.FireCommand)

	vehicle := g.GetVehicle(command.UserId)
	if vehicle == nil {
		return
	}

	log.Println("Fire from:" + command.UserId)

}
