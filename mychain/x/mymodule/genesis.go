package mymodule

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mychain/x/mymodule/keeper"
	"mychain/x/mymodule/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the entity
	for _, elem := range genState.EntityList {
		k.SetEntity(ctx, elem)
	}

	// Set entity count
	k.SetEntityCount(ctx, genState.EntityCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.EntityList = k.GetAllEntity(ctx)
	genesis.EntityCount = k.GetEntityCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
