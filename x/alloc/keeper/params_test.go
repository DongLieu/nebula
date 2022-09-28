package keeper_test

import (
	"testing"

	testkeeper "github.com/nebula-labs/nebula/testutil/keeper"
	"github.com/nebula-labs/nebula/x/alloc/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AllocKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
