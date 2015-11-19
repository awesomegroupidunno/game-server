package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
	"time"
)

const (
	NO_POWERUP = -1
	HEAL       = iota
	SPEEDUP    = iota
	ROCKET     = iota
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

	log.Println("powerup used!")

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
}

func (t *PowerupCommandProcessor) applySpeedPowerUp(v *state.Vehicle) {
	v.ActivePowerup = SPEEDUP
	v.StoredPowerup = NO_POWERUP
	v.OverRideSpeedTill = time.Now().Add(10 * time.Second)
}

func (t *PowerupCommandProcessor) fireRocket(v *state.Vehicle, g *state.GameState) {
	v.StoredPowerup = NO_POWERUP
	targetedVehicle := g.Vehicles[1]
	r := state.Rocket{
		X:        v.X,
		Y:        v.Y,
		Width:    t.Physics.BulletWidth * 1.25,
		Height:   t.Physics.BulletWidth * 1.25,
		Target:   targetedVehicle,
		Velocity: t.Physics.BulletVelocity * .75,
	}

	g.Rockets = append(g.Rockets, &r)
}
