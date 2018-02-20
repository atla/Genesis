package entities

import (
	"fmt"
	"time"
)

// Behaviour ... base interface to execute all behaviours
// behaviour implementations can act on different type of entities
type Behaviour interface {
	Act(time time.Time, entity Entity)
}

// StarvingBehaviour ... starving bevahour is applied to all living entities as a base requirement
type StarvingBehaviour struct {
	repletion    int8
	StarvingRate int8
}

// NewStarvingBehaviour ... creates a new starving behaviour
func NewStarvingBehaviour() *StarvingBehaviour {
	return &StarvingBehaviour{
		repletion:    100,
		StarvingRate: 1,
	}
}

// Act ... act implementation of starving behaviour
func (s *StarvingBehaviour) Act(time time.Time, entity Entity) {

	// starving behaviour only works on living entities
	switch entity.(type) {
	case *LivingEntity:
		s.repletion -= s.StarvingRate

		if s.repletion <= 0 {
			fmt.Println(entity.ID(), " starved to dead.")
		}
		if s.repletion > 100 {
			fmt.Println(entity.ID(), " is well fed.")
		}

	}
}
