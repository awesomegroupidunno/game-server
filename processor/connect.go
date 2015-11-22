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

	newVehicle := state.Vehicle{
		Point: state.Point{
			X: 300,
			Y: 300},
		Sized: state.Sized{
			Width:  t.Physics.VehicleWidth,
			Height: t.Physics.VehicleHeight,
		},
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
