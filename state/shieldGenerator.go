package state

import "time"

type ShieldGenerator struct {
	TeamId        int
	MaxHealth     int
	CurrentHealth int
	RespawnTime   time.Time `json:"-"`
	X             int
	Y             int
	Width         int
}
