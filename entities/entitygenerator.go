package entities

import (
	"fmt"
)

// EntityGenerator ... entity geneator
type EntityGenerator struct {
	entityMap map[string]Entity
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

	fmt.Println("Error, no such type registed " + key)
	return nil
}
