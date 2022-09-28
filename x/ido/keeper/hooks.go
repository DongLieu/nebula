package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	launchpadtypes "github.com/nebula-labs/nebula/x/launchpad/types"
)

func (k Keeper) AfterProjectDeteted(ctx sdk.Context, projectId uint64) {
	k.DeleteIDOById(ctx, projectId)
}

func (k Keeper) AfterWithdrawTokens(ctx sdk.Context, projectId uint64) {

}

// ___________________________________________________________________________________________________

// Hooks wrapper struct for incentives keeper.
type Hooks struct {
	k Keeper
}

var _ launchpadtypes.LaunchpadHooks = Hooks{}

// Return the wrapper struct.
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// hooks.
func (h Hooks) AfterProjectDeteted(ctx sdk.Context, projectId uint64) {
	h.k.AfterProjectDeteted(ctx, projectId)
}

func (h Hooks) AfterWithdrawTokens(ctx sdk.Context, projectId uint64) {
	h.k.AfterWithdrawTokens(ctx, projectId)
}
