package main

import (
	"fmt"
	ent "genesis/entities"
	sim "genesis/simulation"
)

func main() {

	fmt.Println("Starting Genesis")

	var entity = ent.NewBaseEntity()
	var complexEntity = ent.NewComplexEntity()

	complexEntity.AddEntity(entity)

	simulation := &sim.Simulation{

		Params: &sim.Params{
			TickIntervalMilliseconds: 100,
		},
	}
	simulation.AddEntity(entity)
	simulation.AddEntity(complexEntity)

	simulation.Start()

}
