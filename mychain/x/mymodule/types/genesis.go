package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EntityList: []Entity{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in entity
	entityIdMap := make(map[uint64]bool)
	entityCount := gs.GetEntityCount()
	for _, elem := range gs.EntityList {
		if _, ok := entityIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for entity")
		}
		if elem.Id >= entityCount {
			return fmt.Errorf("entity id should be lower or equal than the last id")
		}
		entityIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
