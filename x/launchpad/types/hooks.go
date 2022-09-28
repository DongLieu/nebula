package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type LaunchpadHooks interface {
	AfterProjectDeteted(ctx sdk.Context, projectId uint64)
	AfterWithdrawTokens(ctx sdk.Context, projectId uint64)
}

var _ LaunchpadHooks = MultiLaunchpadHooks{}

// combine multiple launchpad hooks, all hook functions are run in array sequence.
type MultiLaunchpadHooks []LaunchpadHooks

// Creates hooks for the Launchpad Module.
func NewMultiProjectHooks(hooks ...LaunchpadHooks) MultiLaunchpadHooks {
	return hooks
}

func (h MultiLaunchpadHooks) AfterProjectDeteted(ctx sdk.Context, projectID uint64) {
	for i := range h {
		h[i].AfterProjectDeteted(ctx, projectID)
	}
}

func (h MultiLaunchpadHooks) AfterWithdrawTokens(ctx sdk.Context, projectID uint64) {
	for i := range h {
		h[i].AfterWithdrawTokens(ctx, projectID)
	}
}
