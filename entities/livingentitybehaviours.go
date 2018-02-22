package entities

import (
	"fmt"
	"time"
)

// StarvingBehaviour ... starving bevahour is applied to all living entities as a base requirement
type StarvingBehaviour struct {
	repletion    float32
	StarvingRate float32
}

// NewStarvingBehaviour ... creates a new starving behaviour
func NewStarvingBehaviour(starvingRate float32) *StarvingBehaviour {
	return &StarvingBehaviour{
		repletion:    100,
		StarvingRate: starvingRate,
	}
}

// Act ... act implementation of starving behaviour
func (s *StarvingBehaviour) Act(time time.Time, entity interface{}) {

	// starving behaviour only works on living entities
	switch entity.(type) {
	case LivingEntity:
		var livingEntity = entity.(LivingEntity)
		s.repletion -= s.StarvingRate

		if s.repletion <= 0 {
			fmt.Println(livingEntity.ID(), " starved to dead.")
		}
		if s.repletion > 100 {
			fmt.Println(livingEntity.ID(), " is well fed.")
		}
	default:
		fmt.Println("Error: StarvingBehaviour is not supported on entity type ", entity)
	}
}
