package main

import (
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

// SimulationParams ... define various global params for simulation
type SimulationParams struct {
	tickIntervalMilliseconds time.Duration
}

// Simulation ... server struct
type Simulation struct {
	ticker   *time.Ticker
	params   *SimulationParams
	running  bool
	entities []GenEntity
}

// GenEntity ... every entity in the simulation shall be derived from SimEntity
type GenEntity interface {
	Process(time time.Time)
	ID() string
}

// BaseEntity ... standard entity used by eveyy unit in the system
type BaseEntity struct {
	id    string
	kind  string
	alive bool
}

// ComplexEntity ... is itself a entity and contains a set of child entities
type ComplexEntity struct {
	entities []*BaseEntity
	*BaseEntity
}

// NewBaseEntity ... creates and returns a new base entity instance
func NewBaseEntity() *BaseEntity {
	return &BaseEntity{
		id:    ksuid.New().String(),
		alive: true,
	}
}

// NewComplexEntity ... create a new complex entitiy
func NewComplexEntity() *ComplexEntity {
	return &ComplexEntity{
		BaseEntity: NewBaseEntity(),
	}
}

// Process ... processes every entity
func (s *BaseEntity) Process(time time.Time) {
	fmt.Println("Processing entity ID: ", s.id, "time: ", time)
}

// ID ... returns the id of the entity
func (s *BaseEntity) ID() string {
	return s.id
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

func main() {

	fmt.Println("Starting Genesis")

	var entity = NewBaseEntity()
	var complexEntity = NewComplexEntity()

	complexEntity.AddEntity(entity)

	simulation := &Simulation{
		params: &SimulationParams{
			tickIntervalMilliseconds: 100,
		},
	}
	simulation.AddEntity(entity)
	simulation.AddEntity(complexEntity)

	simulation.Start()

}

// Start ... starts the server with a ticking delay of 5 milliseconds
func (simulation *Simulation) Start() {
	simulation.ticker = time.NewTicker(time.Millisecond * simulation.params.tickIntervalMilliseconds)
	simulation.running = true

	go func() {
		for t := range simulation.ticker.C {
			for _, ent := range simulation.entities {
				ent.Process(t)
			}
		}
	}()

	simulation.RunFor(time.Minute * 1)
	simulation.Loop()
}

// RunFor ... asdasd
func (simulation *Simulation) RunFor(duration time.Duration) {
	go func() {
		time.Sleep(duration)
		fmt.Println("Stopping simulation...")
		simulation.running = false
	}()
}

// AddEntity ... add an entity to the simulation
func (simulation *Simulation) AddEntity(entity GenEntity) {
	simulation.entities = append(simulation.entities, entity)
}

// Shutdown ... asdad
func (simulation *Simulation) Shutdown() {
	fmt.Println("Simulation shutdown")
}

// Loop ... main simulation loop that runs until simulation.running is set to false
func (simulation *Simulation) Loop() {
	for {
		if !simulation.running {
			simulation.ticker.Stop()
			simulation.Shutdown()
			break
		}
	}
}
