package cmd

type Acceleration struct {
	value FloatType
	Command
}

func newAcceleration(v FloatType) Acceleration {
	return TurnCommand{Command{Type: "POST", Subtype: "ACCELERATION"}, value: v}
}
