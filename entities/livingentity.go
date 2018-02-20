package entities

import (
	"fmt"
	"time"
)

// LivingEntity ... standard entity used by eveyy unit in the system
type LivingEntity struct {
	*BaseEntity
}

// NewLivingEntity ... creates and returns a new base entity instance
func NewLivingEntity() *LivingEntity {
	return &LivingEntity{
		BaseEntity: NewBaseEntity(),
	}
}

// Process ... processes every entity
func (s *LivingEntity) Process(time time.Time) {
	fmt.Println("Processing living entity ID: ", s.id, "time: ", time)
}
