package examplemodule_test

import (
	"testing"

	keepertest "example-chain/testutil/keeper"
	"example-chain/testutil/nullify"
	"example-chain/x/examplemodule"
	"example-chain/x/examplemodule/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExamplemoduleKeeper(t)
	examplemodule.InitGenesis(ctx, *k, genesisState)
	got := examplemodule.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
