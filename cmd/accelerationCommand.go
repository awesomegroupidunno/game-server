package cmd

const Acceleration = "ACCELERATION"

type AccelerationCommand struct {
	BaseCommand
	Value float64 `json:"NumValue"`
}

func NewAcceleration(v float64) AccelerationCommand {
	return AccelerationCommand{BaseCommand{Type: Post, Subtype: Acceleration, UniqueId: ""}, v}
}

func (b *AccelerationCommand) Command() *BaseCommand {
	return &(b.BaseCommand)
}
