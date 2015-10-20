package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"math"
)

type TurnCommandProcessor struct {
	Physics *Physics
}

func (t *TurnCommandProcessor) Run(g *state.GameState, c cmd.GameCommand) {
	command := c.(*cmd.TurnCommand)

	vehicle := g.GetVehicle(command.UserId)
	if vehicle == nil {
		return
	}
	temp := vehicle

	temp.Angle = math.Mod(temp.Angle+(command.Value*t.Physics.TurnCommandModifier), 2*math.Pi)
	vehicle = temp
}
