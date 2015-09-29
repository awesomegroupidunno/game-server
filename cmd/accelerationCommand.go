package cmd

type AccelerationCommand struct {
	BaseCommand
	Value float32
}

func NewAcceleration(v float32) AccelerationCommand {
	return AccelerationCommand{BaseCommand{Type: "POST", Subtype: "ACCELERATION", UniqueId: ""}, v}
}

func (b *AccelerationCommand) Command() BaseCommand {
	return BaseCommand{Type: b.Type, Subtype: b.Subtype, UniqueId: b.UniqueId}
}
