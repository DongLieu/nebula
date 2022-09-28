package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/nebula-labs/nebula/x/launchpad/types"
)

func (k Keeper) CreateProject(ctx sdk.Context, projectOwner sdk.AccAddress, msg *types.MsgCreateProjectRequest) (uint64, error) {
	// get project id
	projectID := k.GetNextProjectIDAndIncrement(ctx)

	// get project address
	projectAddress := k.NewProjectAddress(projectID)

	// create project
	project := types.Project{
		ProjectOwner:       msg.GetOwner(),
		ProjectTitle:       msg.GetProjectTitle(),
		ProjectId:          projectID,
		ProjectAddress:     projectAddress.String(),
		ProjectInformation: msg.GetProjectInformation(),
		ProjectStatus:      types.PROJECT_INIT,
		ProjectVerified:    false,
		RegisteredRm:       []*codectypes.Any{},
	}

	// save project module address to the account keeper
	acc := k.accountKeeper.NewAccount(
		ctx,
		authtypes.NewModuleAccount(
			authtypes.NewBaseAccountWithAddress(projectAddress),
			project.GetProjectAddress(),
		),
	)
	k.accountKeeper.SetAccount(ctx, acc)

	// save project to KV stores
	if err := k.SetProject(ctx, project); err != nil {
		return 0, err
	}

	return projectID, nil
}

func (k Keeper) DeleteProject(ctx sdk.Context, projectOwner sdk.AccAddress, msg *types.MsgDeleteProjectRequest) error {
	// get project id
	projectId := msg.GetProjectId()

	// get project by id
	project, err := k.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}

	// check if msg.Owner is current project owner
	if project.GetProjectOwner() != msg.GetOwner() {
		return types.ErrNotProjecOwner
	}

	if project.ProjectStatus != types.PROJECT_INIT {
		return types.ErrCannotDeleteProject
	}

	// check if all registered release mechanism is inactive
	for _, any := range project.RegisteredRm {
		rm, err := k.ParseAnyReleaseMechanism(any)
		if err != nil {
			return err
		}
		if rm.GetReleaseMechanismStatus() != types.RM_INIT {
			return types.ErrCannotDeleteProject
		}
	}

	projectAddress := sdk.AccAddress(project.ProjectAddress)

	// withdraw tokens to owner
	tokens := k.bankKeeper.GetAllBalances(ctx, projectAddress)
	err = k.bankKeeper.SendCoins(ctx, projectAddress, projectOwner, tokens)
	if err != nil {
		return err
	}

	// delete project module address from the account keeper
	k.accountKeeper.RemoveAccount(
		ctx,
		authtypes.NewModuleAccount(
			authtypes.NewBaseAccountWithAddress(projectAddress),
			project.GetProjectAddress(),
		),
	)

	// delete project from KV stores
	k.DeleteProjectById(ctx, projectId)

	// delete all release mechanisms through hooks
	k.hooks.AfterProjectDeteted(ctx, projectId)

	return nil
}

func (k Keeper) WithdrawTokens(ctx sdk.Context, projectOwner sdk.AccAddress, msg *types.MsgWithdrawAllTokensRequest) error {
	// get project by id
	project, err := k.GetProjectById(ctx, msg.ProjectId)
	if err != nil {
		return err
	}

	// check if msg.Owner is current project owner
	if project.GetProjectOwner() != msg.GetOwner() {
		return types.ErrNotProjecOwner
	}

	if project.ProjectStatus != types.PROJECT_ENDED {
		return types.ErrCannotWithdrawTokens
	}

	// check if all registered release mechanism has ended (status: 2)
	for _, any := range project.RegisteredRm {
		rm, err := k.ParseAnyReleaseMechanism(any)
		if err != nil {
			return err
		}
		if rm.GetReleaseMechanismStatus() != types.RM_ENDED {
			return types.ErrCannotWithdrawTokens
		}
	}

	projectModuleAddress := sdk.AccAddress(project.ProjectAddress)

	// withdraw all tokens to owner
	tokens := k.bankKeeper.GetAllBalances(ctx, projectModuleAddress)
	err = k.bankKeeper.SendCoins(ctx, projectModuleAddress, projectOwner, tokens)
	if err != nil {
		return err
	}

	k.hooks.AfterWithdrawTokens(ctx, project.ProjectId)

	return nil
}
