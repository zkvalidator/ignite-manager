package examplemodule

import (
	"example-chain/x/examplemodule/keeper"
	"example-chain/x/examplemodule/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the entityName
	for _, elem := range genState.EntityNameList {
		k.SetEntityName(ctx, elem)
	}

	// Set entityName count
	k.SetEntityNameCount(ctx, genState.EntityNameCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.EntityNameList = k.GetAllEntityName(ctx)
	genesis.EntityNameCount = k.GetEntityNameCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
