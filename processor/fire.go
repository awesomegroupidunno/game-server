package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
	"time"
)

type FireCommandProcessor struct {
	Physics   *Physics
	lastFired map[string]time.Time
}

func (t *FireCommandProcessor) Run(g *state.GameState, c cmd.GameCommand) {

	if t.lastFired == nil {
		t.lastFired = make(map[string]time.Time)
	}
	command := c.(*cmd.FireCommand)

	vehicle := g.GetVehicle(command.UserId)
	if vehicle == nil || !vehicle.IsAlive {
		return
	}

	last := t.lastFired[vehicle.Owner]
	diff := time.Now().Sub(last)

	if diff < t.Physics.BulletDelay {
		return
	}

	t.lastFired[vehicle.Owner] = time.Now()

	b := state.Bullet{
		X:        vehicle.X,
		Y:        vehicle.Y,
		Width:    t.Physics.BulletWidth,
		Height:   t.Physics.BulletWidth,
		Velocity: t.Physics.BulletVelocity,
		Angle:    vehicle.Angle,
		OwnerId:  vehicle.Owner}

	g.Bullets = append(g.Bullets, &b)

}
