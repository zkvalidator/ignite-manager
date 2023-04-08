package keeper_test

import (
	"testing"

	testkeeper "example-chain/testutil/keeper"
	"example-chain/x/examplemodule/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ExamplemoduleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
