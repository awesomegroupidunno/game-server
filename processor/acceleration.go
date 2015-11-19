package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"time"
)

type AccelerationCommandProcessor struct {
	Physics *Physics
}

func (t *AccelerationCommandProcessor) Run(g *state.GameState, c cmd.GameCommand) {
	command := c.(*cmd.AccelerationCommand)

	vehicle := g.GetVehicle(command.UserId)
	if vehicle == nil || !vehicle.IsAlive {
		return
	}

	powerupOn := vehicle.ActivePowerup == SPEEDUP && vehicle.OverRideSpeedTill.After(time.Now())

	if powerupOn {
		t.Physics.AccelerationCommandModifier *= 2
		t.Physics.MaxVehicleVelocity *= 2
	}

	t.accelerateVehicle(vehicle, command)

	if powerupOn {
		t.Physics.AccelerationCommandModifier /= 2
		t.Physics.MaxVehicleVelocity /= 2
	}
}

func (t *AccelerationCommandProcessor) accelerateVehicle(v *state.Vehicle, c *cmd.AccelerationCommand) {
	v.Velocity = v.Velocity + (c.Value * t.Physics.AccelerationCommandModifier)

	if v.Velocity >= t.Physics.MaxVehicleVelocity {
		v.Velocity = t.Physics.MaxVehicleVelocity
	}

	if (v.Velocity) <= (-1 * t.Physics.MaxVehicleVelocity) {
		v.Velocity = -1 * t.Physics.MaxVehicleVelocity
	}
}
