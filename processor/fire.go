package processor

import (
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/awesomegroupidunno/game-server/state"
)

type FireCommandProcessor struct {
	Physics *Physics
}

func (t *FireCommandProcessor) Run(g *state.GameState, c cmd.GameCommand) {
	command := c.(*cmd.FireCommand)

	vehicle := g.GetVehicle(command.UserId)
	if vehicle == nil {
		return
	}

	b := state.Bullet{
		X:        vehicle.X,
		Y:        vehicle.Y,
		Width:    t.Physics.BulletWidth,
		Height:   t.Physics.BulletWidth,
		Velocity: t.Physics.BulletVelocity,
		Angle:    vehicle.Velocity}

	g.Bullets = append(g.Bullets, &b)

}
