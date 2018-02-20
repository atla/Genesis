package simulation

import (
	"fmt"
	ent "genesis/entities"
	"time"
)

// Params ... define various global params for simulation
type Params struct {
	TickIntervalMilliseconds time.Duration
}

// Simulation ... server struct
type Simulation struct {
	ticker   *time.Ticker
	Params   *Params
	running  bool
	entities []ent.GenesisEntity
}

// Start ... starts the server with a ticking delay of 5 milliseconds
func (simulation *Simulation) Start() {
	simulation.ticker = time.NewTicker(time.Millisecond * simulation.Params.TickIntervalMilliseconds)
	simulation.running = true

	go func() {
		for t := range simulation.ticker.C {
			for _, ent := range simulation.entities {
				ent.Process(t)
			}
		}
	}()

	simulation.runFor(time.Minute * 1)
	simulation.loop()
}

// RunFor ... asdasd
func (simulation *Simulation) runFor(duration time.Duration) {
	go func() {
		time.Sleep(duration)
		fmt.Println("Stopping simulation...")
		simulation.running = false
	}()
}

// AddEntity ... add an entity to the simulation
func (simulation *Simulation) AddEntity(entity ent.GenesisEntity) {
	simulation.entities = append(simulation.entities, entity)
}

// Shutdown ... asdad
func (simulation *Simulation) Shutdown() {
	fmt.Println("Simulation shutdown")
}

// Loop ... main simulation loop that runs until simulation.running is set to false
func (simulation *Simulation) loop() {
	for {
		if !simulation.running {
			simulation.ticker.Stop()
			simulation.Shutdown()
			break
		}
	}
}
