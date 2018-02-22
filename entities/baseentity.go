package entities

import (
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

// BaseEntity ... standard entity used by eveyy unit in the system
type BaseEntity struct {
	id          string
	description string
	kind        string
	alive       bool
	position    Vector
}

// NewBaseEntity ... creates and returns a new base entity instance
func NewBaseEntity() *BaseEntity {
	return &BaseEntity{
		id:    ksuid.New().String(),
		alive: true,
	}
}

// Process ... processes every entity
func (be *BaseEntity) Process(time time.Time) {
	fmt.Println("Processing entity ID: ", be.id, "time: ", time)
}

// ID ... returns the id of the entity
func (be *BaseEntity) ID() string {
	return be.id
}

// IsAlive returns a alive/dead state of the entity
func (be *BaseEntity) IsAlive() bool {
	return be.alive
}

// Position returns the entitites position
func (be *BaseEntity) Position() Vector {
	return be.position
}

func (be *BaseEntity) Description() string {
	return be.description
}
