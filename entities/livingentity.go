package entities

import (
	"fmt"
	"time"
)

// LivingEntity ... standard entity used by eveyy unit in the system
type LivingEntity struct {
	*BaseEntity
	behaviours []Behaviour
}

// NewLivingEntity ... creates and returns a new base entity instance
func NewLivingEntity(behaviours []Behaviour) *LivingEntity {
	return &LivingEntity{
		BaseEntity: NewBaseEntity(),
		behaviours: behaviours,
	}
}

// Process ... processes every entity
func (le *LivingEntity) Process(time time.Time) {
	fmt.Println("Processing living entity ID: ", le.id, "time: ", time)

	for _, behaviour := range le.behaviours {
		behaviour.Act(time, le)
	}
}
