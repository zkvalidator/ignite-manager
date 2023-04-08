package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "m-example-chain/testutil/keeper"
	"m-example-chain/x/mexamplemodule/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MexamplemoduleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
