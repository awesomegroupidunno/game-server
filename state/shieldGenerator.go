package state

import "time"

type ShieldGenerator struct {
	Team_id        int
	Max_health     int
	Current_health int
	RespawnTime    time.Time
}
