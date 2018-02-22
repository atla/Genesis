package entities

import (
	"fmt"
	"time"
)

// LivingEntity ... standard entity used by eveyy unit in the system
type LivingEntity struct {
	*BaseEntity
	Behaviours []Behaviour
}

// NewLivingEntity ... creates and returns a new base entity instance
func NewLivingEntity() *LivingEntity {
	return &LivingEntity{
		BaseEntity: NewBaseEntity(),
	}
}

// Process ... processes every entity
func (le *LivingEntity) Process(time time.Time) {
	fmt.Println("Processing living entity ID: ", le.id, "time: ", time)

	for _, Behaviour := range le.Behaviours {
		Behaviour.Act(time, le)
	}
}

// AddBehaviour adds a new behaviour
func (le *LivingEntity) AddBehaviour(b Behaviour) {

	le.Behaviours = append(le.Behaviours, b)

}
