package entities

import (
	"math/rand"
	"sync"
)

// Resource can model any kind of resource needed by an entity
type Resource struct {
	sync.Mutex
	// Does it need an ID?
	Type    string
	Amount  float64
	Endless bool
}

// ResourceTypes includes the know resource types
var ResourceTypes = []string{
	"food_berries",
	"food_algae",
	"food_plankton",
}

func getRandomType() string {
	return ResourceTypes[rand.Int()%len(ResourceTypes)]
}

// GetRandomResource creates a random resource type based on given parameters
func GetRandomResource(minAmount float64, maxAmount float64, isEndless bool) *Resource {

	//TODO: use random seed?
	amount := minAmount + float64(rand.Int63n(int64(maxAmount-minAmount)))

	return &Resource{
		Type:    getRandomType(),
		Amount:  amount,
		Endless: isEndless,
	}
}

// NewResource creates a new resource instance
func NewResource(Type string, amount float64, endless bool) *Resource {
	return &Resource{
		Type:    Type,
		Amount:  amount,
		Endless: endless,
	}
}

// IsDepleted returns true if amount is smaller than 0
func (r *Resource) IsDepleted() bool {
	return r.Amount < 0.0
}

// Consume s a specified amount from the resource and returns a bool if it was successful and the amount it consumed
func (r *Resource) Consume(amount float64) (bool, float64) {

	r.Lock()
	switch {
	case r.Amount >= amount:
		r.Amount -= amount
		r.Unlock()
		return true, amount
	default:
		r.Unlock()
		return false, 0.0
	}
}
