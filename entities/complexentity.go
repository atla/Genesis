package entities

import (
	"fmt"
	"time"
)

// ComplexEntity ... is itself a entity and contains a set of child entities
type ComplexEntity struct {
	entities []*BaseEntity
	*BaseEntity
}

// NewComplexEntity ... create a new complex entitiy
func NewComplexEntity() *ComplexEntity {
	return &ComplexEntity{
		BaseEntity: NewBaseEntity(),
	}
}

// Process ... processes every entity
func (s *ComplexEntity) Process(time time.Time) {
	fmt.Println("Processing complex entity ID: ", s.id, " time: ", time)
	//do not process child entities?
}

// AddEntity ... asd
func (s *ComplexEntity) AddEntity(entity *BaseEntity) {
	s.entities = append(s.entities, entity)
}
