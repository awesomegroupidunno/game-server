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

	vehicle := g.GetVehicle(command.UserId)
	if vehicle == nil {
		return
	}
	temp := vehicle

	temp.Velocity = temp.Velocity + command.Value

	vehicle = temp

	log.Printf("%v", vehicle)

}
