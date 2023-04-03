package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EntityNameList: []EntityName{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in entityName
	entityNameIdMap := make(map[uint64]bool)
	entityNameCount := gs.GetEntityNameCount()
	for _, elem := range gs.EntityNameList {
		if _, ok := entityNameIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for entityName")
		}
		if elem.Id >= entityNameCount {
			return fmt.Errorf("entityName id should be lower or equal than the last id")
		}
		entityNameIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
