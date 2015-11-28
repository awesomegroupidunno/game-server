package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"time"
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

	size := state.NewSized(t.Physics.VehicleWidth, t.Physics.VehicleHeight)
	pt := t.Physics.findSpace(size, *g)

	newVehicle := state.Vehicle{
		Point:             pt,
		Sized:             size,
		Velocity:          0.0,
		Angle:             0,
		TeamId:            teamNum,
		MaxHealth:         t.Physics.VehicleHealth,
		CurrentHealth:     t.Physics.VehicleHealth,
		Owner:             command.UserId,
		Mass:              10,
		IsAlive:           true,
		ActivePowerup:     NO_POWERUP,
		StoredPowerup:     NO_POWERUP,
		OverRideSpeedTill: time.Now().Add(-5 * time.Second)}
	g.Vehicles = append(g.Vehicles, &newVehicle)

}
