package keeper_test

import (
	"context"
	"testing"

	keepertest "example-chain/testutil/keeper"
	"example-chain/x/examplemodule/keeper"
	"example-chain/x/examplemodule/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ExamplemoduleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
