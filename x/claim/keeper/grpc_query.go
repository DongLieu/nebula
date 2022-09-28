package keeper

import (
	"github.com/nebula-labs/nebula/x/claim/types"
)

var _ types.QueryServer = Keeper{}
