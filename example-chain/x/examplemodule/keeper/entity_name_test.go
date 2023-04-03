package keeper_test

import (
	"testing"

	keepertest "example-chain/testutil/keeper"
	"example-chain/testutil/nullify"
	"example-chain/x/examplemodule/keeper"
	"example-chain/x/examplemodule/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNEntityName(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EntityName {
	items := make([]types.EntityName, n)
	for i := range items {
		items[i].Id = keeper.AppendEntityName(ctx, items[i])
	}
	return items
}

func TestEntityNameGet(t *testing.T) {
	keeper, ctx := keepertest.ExamplemoduleKeeper(t)
	items := createNEntityName(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEntityName(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEntityNameRemove(t *testing.T) {
	keeper, ctx := keepertest.ExamplemoduleKeeper(t)
	items := createNEntityName(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEntityName(ctx, item.Id)
		_, found := keeper.GetEntityName(ctx, item.Id)
		require.False(t, found)
	}
}

func TestEntityNameGetAll(t *testing.T) {
	keeper, ctx := keepertest.ExamplemoduleKeeper(t)
	items := createNEntityName(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEntityName(ctx)),
	)
}

func TestEntityNameCount(t *testing.T) {
	keeper, ctx := keepertest.ExamplemoduleKeeper(t)
	items := createNEntityName(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEntityNameCount(ctx))
}
