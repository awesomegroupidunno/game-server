package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"math/rand"
)

type ConnectCommandProcessor struct {
}

func (t *ConnectCommandProcessor) Run(g *state.GameState, c cmd.GameCommand) {
	command := c.(*cmd.ConnectCommand)

	// if the user already has a vehicle, ignore
	vehicle := g.GetVehicle(c.Command().UserId)
	if vehicle != nil {
		return
	}

	// For now, randomly join team 0 or 1
	teamNum := rand.Intn(2)

	newVehicle := state.Vehicle{
		X:             10,
		Y:             10,
		Velocity:      0.0,
		Angle:         0.0,
		Endurance:     100,
		Team_id:       teamNum,
		Max_health:    100,
		Current_heath: 100,
		Owner:         command.UserId}
	g.Vehicles = append(g.Vehicles, newVehicle)

}
