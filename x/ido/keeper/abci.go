package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nebula-labs/nebula/x/ido/types"
	launchpadtypes "github.com/nebula-labs/nebula/x/launchpad/types"
)

func (k Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// iterate through each ido and enable those that have reached time
	k.IterateIDO(ctx, func(index int64, ido types.IDO) (stop bool) {

		// pass ended project
		if ido.IdoStatus == launchpadtypes.RM_ENDED {
			return false
		}

		// check if a project time has come for IDO_INIT
		if ctx.BlockTime().After(ido.StartTime) && ido.IdoStatus == launchpadtypes.RM_INIT {
			ido.IdoStatus = launchpadtypes.RM_ACTIVE
			k.SetIDOAndRegisterLaunchpad(ctx, ido)
			k.launchpadKeeper.SetProjectActive(ctx, ido.ProjectId)
			return false
		}

		// check if an ido has reached end condition
		if canEnd(ido) {
			ido.IdoStatus = launchpadtypes.RM_ENDED
			k.SetIDOAndRegisterLaunchpad(ctx, ido)
			k.launchpadKeeper.SetProjectEndable(ctx, ido.ProjectId)
			return false
		}

		return false
	})
}

func canEnd(ido types.IDO) bool {

	if len(ido.TokenForDistribution) > 0 {
		return false
	}

	// ido status must be active
	if ido.IdoStatus != launchpadtypes.RM_ACTIVE {
		return false
	}

	return true
}
