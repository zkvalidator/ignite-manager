package keeper

import (
	"m-example-chain/x/mexamplemodule/types"
)

var _ types.QueryServer = Keeper{}
