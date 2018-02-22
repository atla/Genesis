package entities

import (
	"fmt"
	"time"
)

// AreaEntity ... is itself a entity and contains a set of child entities
type AreaEntity struct {
	*ComplexEntity
}

// NewAreaEntity ... create a new AreaEntity entitiy
func NewAreaEntity() *AreaEntity {
	return &AreaEntity{
		ComplexEntity: NewComplexEntity(),
	}
}

// Process ... processes every entity
func (s *AreaEntity) Process(time time.Time) {
	fmt.Println("Processing area entity ID: ", s.id, " time: ", time)
	//do not process child entities?

	// run behaviours

}
