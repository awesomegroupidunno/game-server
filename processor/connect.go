package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
)

type ConnectCommandProcessor struct {
	Physics *Physics
}

func (t *ConnectCommandProcessor) Run(g *state.GameState, c cmd.GameCommand) {
	command := c.(*cmd.ConnectCommand)

	// if the user already has a vehicle, ignore
	vehicle := g.GetVehicle(command.UserId)
	if vehicle != nil {
		return
	}

	// For now, randomly join team 0 or 1
	teamNum := len(g.Vehicles) % 2

	newVehicle := state.Vehicle{
		X:             300,
		Y:             300,
		Velocity:      0.0,
		Angle:         0.0,
		Endurance:     100,
		TeamId:        teamNum,
		Max_health:    100,
		CurrentHealth: 100,
		Width:         t.Physics.VehicleWidth,
		Height:        t.Physics.VehicleHeight,
		Owner:         command.UserId}
	g.Vehicles = append(g.Vehicles, &newVehicle)

}
