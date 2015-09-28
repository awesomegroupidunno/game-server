package cmd

type AccelerationCommand struct {
	Command
	Value float32
}

func newAcceleration(v float32) AccelerationCommand {
	return AccelerationCommand{Command{Type: "POST", Subtype: "ACCELERATION", UniqueId: ""}, v}
}
