package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "mychain/testutil/keeper"
	"mychain/testutil/nullify"
	"mychain/x/mymodule/keeper"
	"mychain/x/mymodule/types"
)

func createNEntity(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Entity {
	items := make([]types.Entity, n)
	for i := range items {
		items[i].Id = keeper.AppendEntity(ctx, items[i])
	}
	return items
}

func TestEntityGet(t *testing.T) {
	keeper, ctx := keepertest.MymoduleKeeper(t)
	items := createNEntity(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEntity(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEntityRemove(t *testing.T) {
	keeper, ctx := keepertest.MymoduleKeeper(t)
	items := createNEntity(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEntity(ctx, item.Id)
		_, found := keeper.GetEntity(ctx, item.Id)
		require.False(t, found)
	}
}

func TestEntityGetAll(t *testing.T) {
	keeper, ctx := keepertest.MymoduleKeeper(t)
	items := createNEntity(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEntity(ctx)),
	)
}

func TestEntityCount(t *testing.T) {
	keeper, ctx := keepertest.MymoduleKeeper(t)
	items := createNEntity(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEntityCount(ctx))
}
