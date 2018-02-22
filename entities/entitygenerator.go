package entities

import (
	"fmt"
)

// EntityGeneratorInstance ... public singleton for EntityGenerator
var EntityGeneratorInstance = NewEntityGenerator()

// EntityGenerator ... entity geneator
type EntityGenerator struct {
	entityMap map[string]Entity
}

// NewEntityGenerator ... creates a new entity generator
func NewEntityGenerator() *EntityGenerator {
	return &EntityGenerator{
		entityMap: map[string]Entity{},
	}
}

func init() {

	// setup some initial entity types
	var microbe1 = NewLivingEntity()
	microbe1.AddBehaviour(NewStarvingBehaviour(1.0))

	var area1 = NewAreaEntity()

	EntityGeneratorInstance.Register("microbe", microbe1)
	EntityGeneratorInstance.Register("area", area1)
}

// Register ... register new entity templates
func (eg *EntityGenerator) Register(key string, value Entity) {

	if _, ok := eg.entityMap[key]; ok {
		//do something here
		fmt.Println("Error, entity type already registered", key)
	}

	eg.entityMap[key] = value
}

// Generate ... asdasd
func (eg *EntityGenerator) Generate(key string) Entity {

	if val, ok := eg.entityMap[key]; ok {

		//TODO: create deep copy of entity
		return val
	}

	fmt.Println("Error, no such type registered " + key)
	return nil
}
