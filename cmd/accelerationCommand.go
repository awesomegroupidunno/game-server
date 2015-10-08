package cmd

const Acceleration = "ACCELERATION"

type AccelerationCommand struct {
	BaseCommand
	Value float32
}

func NewAcceleration(v float32) AccelerationCommand {
	return AccelerationCommand{BaseCommand{Type: Post, Subtype: Acceleration, UniqueId: ""}, v}
}

func (b *AccelerationCommand) Command() BaseCommand {
	return BaseCommand{Type: b.Type, Subtype: b.Subtype, UniqueId: b.UniqueId, UserId: b.UserId}
}

func (b *AccelerationCommand) OwnerId() string {
	return b.UserId
}

func (b *AccelerationCommand) SetOwnerId(s string) {
	b.UserId = s
}
