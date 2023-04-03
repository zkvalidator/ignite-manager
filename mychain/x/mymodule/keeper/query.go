package keeper

import (
	"mychain/x/mymodule/types"
)

var _ types.QueryServer = Keeper{}
