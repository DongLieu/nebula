package keeper

import (
	"github.com/nebula-labs/nebula/x/alloc/types"
)

var _ types.QueryServer = Keeper{}
