package mexamplemodule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "m-example-chain/testutil/keeper"
	"m-example-chain/testutil/nullify"
	"m-example-chain/x/mexamplemodule"
	"m-example-chain/x/mexamplemodule/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EntityNameList: []types.EntityName{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		EntityNameCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MexamplemoduleKeeper(t)
	mexamplemodule.InitGenesis(ctx, *k, genesisState)
	got := mexamplemodule.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EntityNameList, got.EntityNameList)
	require.Equal(t, genesisState.EntityNameCount, got.EntityNameCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
