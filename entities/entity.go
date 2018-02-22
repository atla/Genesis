package entities

import "time"

// Vector represents a point in ℝ³.
type Vector struct {
	X, Y, Z float64
}

// Entity is a base interface for all simulated objects
type Entity interface {
	Process(time time.Time)
	ID() string
	IsAlive() bool
	Position() Vector
	Description() string
}
