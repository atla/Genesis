package entities

import "time"

// GenesisEntity ... every entity in the simulation shall be derived from SimEntity
type Entity interface {
	Process(time time.Time)
	ID() string
	IsAlive() bool
}