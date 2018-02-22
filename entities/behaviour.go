package entities

import (
	"time"
)

// Behaviour ... base interface to execute all behaviours
// behaviour implementations can act on different type of entities
type Behaviour interface {
	Act(time time.Time, entity interface{})
}
