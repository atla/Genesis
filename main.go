package main

import (
	"genesis/entities"
	"genesis/simulation"
)

func main() {

	genesis := NewGenesis(entities.EntityGeneratorInstance, simulation.NewSimulation(simulation.NewParams(100)))
	genesis.Run()

}
