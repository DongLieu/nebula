package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nebula-labs/nebula/x/ido/types"
	launchpadtypes "github.com/nebula-labs/nebula/x/launchpad/types"
)

func (k Keeper) validateEnableIDO(ctx sdk.Context, owner sdk.AccAddress, project launchpadtypes.Project, msg *types.MsgEnableIDORequest) error {
	// check if owner is true owner of project
	projectOwner, err := sdk.AccAddressFromBech32(project.GetProjectOwner())
	if err != nil {
		return err
	}

	if !owner.Equals(projectOwner) {
		return launchpadtypes.ErrNotProjecOwner
	}

	// check if project is in init phase
	if project.ProjectStatus != launchpadtypes.PROJECT_INIT {
		return launchpadtypes.ErrCannotRegisterRm
	}

	// check if there is ido of the same id
	if k.HasIDO(ctx, project.ProjectId) {
		return types.ErrIDOAlreadyExist
	}

	// startTime has to be greater than now
	if msg.StartTime.Before(time.Now()) {
		return types.ErrStartTimeSmallerThanNow
	}

	// check if user address has enough tokens to add to ido. For now, only 1 token is supported
	for _, coin := range msg.TokenForDistribution {
		if !k.bankKeeper.HasBalance(ctx, owner, coin) {
			return types.ErrNotEnoughFunds
		}
	}

	return nil
}

func (k Keeper) validateCommitParticipation(ctx sdk.Context, participant sdk.AccAddress, tokenBought *sdk.Coins, entry *types.Entry, ido types.IDO, project launchpadtypes.Project, msg *types.MsgCommitParticipationRequest) error {
	// validate if IDO is active
	if ido.IdoStatus != launchpadtypes.RM_ACTIVE {
		return types.ErrIDOIsNotActive
	}

	// validate if participant has enough funds
	if !k.bankKeeper.HasBalance(ctx, participant, msg.TokenCommit[0]) {
		return types.ErrNotEnoughFunds
	}

	// validate if participant is not project owner
	if participant.String() == project.ProjectOwner {
		return types.ErrProjectOwnerNotAllowedToCommit
	}

	// validate if TokenCommit denom is allowed in IDO
	if !msg.TokenCommit.DenomsSubsetOf(ido.TokenListingPrice) {
		return types.ErrNoSupportedDenoms
	}

	// validate if TokenCommit is allowed
	*entry = types.Entry{
		Participant:  participant.String(),
		CommitAmount: msg.TokenCommit,
	}

	if val, ok := ido.Entries[participant.String()]; ok {
		entry = &val
		entry.CommitAmount = entry.CommitAmount.Add(msg.TokenCommit...)
	}

	for _, limit := range ido.AllocationLimit {
		if limit.Denom == msg.TokenCommit[0].Denom {
			if entry.CommitAmount[0].IsLT(limit.LowerLimit) {
				return types.ErrOutOfBoundPurchase
			}

			if limit.UpperLimit.IsLT(entry.CommitAmount[0]) {
				return types.ErrOutOfBoundPurchase
			}

			break
		}
	}

	// validate if token bought is enough in TokenForDistribution
	*tokenBought = k.CalculateDistributionTokens(msg.TokenCommit, ido.TokenListingPrice, ido.TokenForDistribution.GetDenomByIndex(0))

	if ido.TokenForDistribution[0].IsLT((*tokenBought)[0]) {
		return types.ErrNotEnoughIDOTokens
	}

	return nil
}

func (k Keeper) EnableIDO(ctx sdk.Context, owner sdk.AccAddress, msg *types.MsgEnableIDORequest) error {
	// get project
	project, err := k.launchpadKeeper.GetProjectById(ctx, msg.ProjectId)
	if err != nil {
		return err
	}

	// validate Msg
	if err := k.validateEnableIDO(ctx, owner, project, msg); err != nil {
		return err
	}

	// create IDO struct
	totalDistributedAmount := make(sdk.Coins, len(msg.TokenForDistribution))
	for i, coin := range msg.TokenForDistribution {
		totalDistributedAmount[i] = sdk.NewCoin(coin.Denom, sdk.ZeroInt())
	}

	ido := types.NewIDO(
		msg.ProjectId,
		msg.TokenForDistribution,
		totalDistributedAmount,
		msg.TokenListingPrice,
		launchpadtypes.RM_INIT,
		msg.AllocationLimit,
		msg.StartTime,
		make(map[string]types.Entry),
	)

	// transfer tokens from user wallet to project wallet
	k.bankKeeper.SendCoins(ctx, owner, sdk.AccAddress(project.GetProjectAddress()), msg.TokenForDistribution)

	// save IDO to KV stores and register to launchpad
	if err := k.SetIDOAndRegisterLaunchpad(ctx, ido); err != nil {
		return err
	}

	return nil
}

func (k Keeper) CommitParticipation(ctx sdk.Context, participant sdk.AccAddress, msg *types.MsgCommitParticipationRequest) error {
	// get IDO
	ido, err := k.GetIDOByID(ctx, msg.ProjectId)
	if err != nil {
		return err
	}

	// get project
	project, err := k.launchpadKeeper.GetProjectById(ctx, msg.ProjectId)
	if err != nil {
		return err
	}

	// validate
	var tokenBought sdk.Coins
	var entry types.Entry

	if err := k.validateCommitParticipation(ctx, participant, &tokenBought, &entry, ido, project, msg); err != nil {
		return err
	}

	// transfer funds from user wallet to project and vice versa
	k.bankKeeper.SendCoins(ctx, sdk.AccAddress(project.GetProjectAddress()), participant, tokenBought)
	k.bankKeeper.SendCoins(ctx, participant, sdk.AccAddress(project.GetProjectAddress()), msg.GetTokenCommit())

	ido.TokenForDistribution = ido.TokenForDistribution.Sub(tokenBought)
	ido.TotalDistributedAmount = ido.TotalDistributedAmount.Add(tokenBought...)

	// update entries
	ido.Entries[entry.Participant] = entry

	// save IDO to KV stores and register to launchpad
	if err := k.SetIDOAndRegisterLaunchpad(ctx, ido); err != nil {
		return err
	}

	return nil
}
