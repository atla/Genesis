package main

import (
	"fmt"
	ent "genesis/entities"
	sim "genesis/simulation"
)

// Genesis ... main startup class for Genesis server
type Genesis struct {
	EntityGenerator *ent.EntityGenerator
	Simulation      *sim.Simulation
}

// NewGenesis ... creates a new genesis instance
func NewGenesis(eg *ent.EntityGenerator, sim *sim.Simulation) *Genesis {
	return &Genesis{eg, sim}
}

// Run ... runs the genesis instance
func (g *Genesis) Run() {

	fmt.Println("Starting Genesis")

	g.Simulation.AddEntity(g.EntityGenerator.Generate("area"))

	g.Simulation.AddEntity(g.EntityGenerator.Generate("microbe"))
	g.Simulation.AddEntity(g.EntityGenerator.Generate("microbe"))
	g.Simulation.AddEntity(g.EntityGenerator.Generate("microbe"))

	g.Simulation.Start()
}

func registerEntityTypes() {

}
