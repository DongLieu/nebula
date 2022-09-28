package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/nebula-labs/nebula/testutil/keeper"
	"github.com/nebula-labs/nebula/x/alloc/keeper"
	"github.com/nebula-labs/nebula/x/alloc/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AllocKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
