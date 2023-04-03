package mymodule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "mychain/testutil/keeper"
	"mychain/testutil/nullify"
	"mychain/x/mymodule"
	"mychain/x/mymodule/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EntityList: []types.Entity{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		EntityCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MymoduleKeeper(t)
	mymodule.InitGenesis(ctx, *k, genesisState)
	got := mymodule.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EntityList, got.EntityList)
	require.Equal(t, genesisState.EntityCount, got.EntityCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
