package cmd

const Acceleration = "ACCELERATION"

type AccelerationCommand struct {
	BaseCommand
	Value float32
}

func NewAcceleration(v float32) AccelerationCommand {
	return AccelerationCommand{BaseCommand{Type: Post, Subtype: Acceleration, UniqueId: ""}, v}
}

func (b *AccelerationCommand) Command() *BaseCommand {
	return &(b.BaseCommand)
}
