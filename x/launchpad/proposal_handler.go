package launchpad

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/nebula-labs/nebula/x/launchpad/keeper"
	"github.com/nebula-labs/nebula/x/launchpad/types"
)

func NewSetProjectVerifiedProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.SetProjectVerifiedProposal:
			return handleSetProjectVerifiedProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized launchpad proposal content type: %T", c)
		}
	}
}

func handleSetProjectVerifiedProposal(ctx sdk.Context, k keeper.Keeper, p *types.SetProjectVerifiedProposal) error {
	return keeper.HandleSetProjectVerifiedProposal(ctx, k, p)
}
