package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"time"
)

const (
	NO_POWERUP   = -1
	HEAL         = iota
	SPEEDUP      = iota
	ROCKET       = iota
	NUM_POWERUPS = iota - 1
)

type PowerupCommandProcessor struct {
	Physics *Physics
}

func (t *PowerupCommandProcessor) Run(g *state.GameState, c cmd.GameCommand) {
	command := c.(*cmd.PowerupCommand)

	vehicle := g.GetVehicle(command.UserId)
	if vehicle == nil || !vehicle.IsAlive {
		return
	}
	switch vehicle.StoredPowerup {

	case HEAL:
		t.healVehicle(vehicle)
		return
	case SPEEDUP:
		t.applySpeedPowerUp(vehicle)
		return
	case ROCKET:
		t.fireRocket(vehicle, g)
		return
	}

}

func (t *PowerupCommandProcessor) healVehicle(v *state.Vehicle) {
	v.CurrentHealth = v.MaxHealth
	v.StoredPowerup = NO_POWERUP
}

func (t *PowerupCommandProcessor) applySpeedPowerUp(v *state.Vehicle) {
	v.ActivePowerup = SPEEDUP
	v.StoredPowerup = NO_POWERUP
	v.OverRideSpeedTill = time.Now().Add(10 * time.Second)
}

func (t *PowerupCommandProcessor) fireRocket(v *state.Vehicle, g *state.GameState) {
	v.StoredPowerup = NO_POWERUP
	targetedVehicle := targetVehicle(v, g)
	r := state.Rocket{
		Point:    state.NewPoint(v.X, v.Y),
		Sized:    state.NewSized(t.Physics.BulletWidth*1.25, t.Physics.BulletWidth*1.25),
		Target:   targetedVehicle,
		Velocity: t.Physics.BulletVelocity * .75,
	}

	g.Rockets = append(g.Rockets, &r)
}

func targetVehicle(v *state.Vehicle, g *state.GameState) *state.Vehicle {
	for _, vehicle := range g.Vehicles {
		if vehicle.Owner != v.Owner && vehicle.TeamId != v.TeamId {
			return vehicle
		}
	}
	return nil
}
