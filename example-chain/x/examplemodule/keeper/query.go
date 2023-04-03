package keeper

import (
	"example-chain/x/examplemodule/types"
)

var _ types.QueryServer = Keeper{}
