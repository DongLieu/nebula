package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/nebula-labs/nebula/x/launchpad/types"
)

func HandleSetProjectVerifiedProposal(ctx sdk.Context, k Keeper, p *types.SetProjectVerifiedProposal) error {
	// check if project is correct owner
	project, err := k.GetProjectById(ctx, p.ProjectId)
	if err != nil {
		return err
	}

	if project.GetProjectOwner() != p.ProjectOwner {
		return types.ErrNotProjecOwner
	}

	project.ProjectVerified = true

	err = k.SetProject(ctx, project)
	if err != nil {
		return err
	}

	event := sdk.NewEvent(
		types.TypeSetProjectVerified,
		sdk.NewAttribute(types.AttributeValueCategory, types.ModuleName),
		sdk.NewAttribute(types.AttributeProjectID, strconv.FormatUint(p.ProjectId, 10)),
	)
	ctx.EventManager().EmitEvent(event)

	return nil
}
