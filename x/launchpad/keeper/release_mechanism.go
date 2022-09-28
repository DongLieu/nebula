package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nebula-labs/nebula/x/launchpad/types"
)

func (k Keeper) MarshalReleaseMechanism(rm types.ReleaseMechanismI) ([]byte, error) {
	return k.cdc.MarshalInterface(rm)
}

func (k Keeper) UnmarshalReleaseMechanism(bz []byte) (types.ReleaseMechanismI, error) {
	var res types.ReleaseMechanismI
	return res, k.cdc.UnmarshalInterface(bz, &res)
}

func (k Keeper) ParseAnyReleaseMechanism(any *codectypes.Any) (types.ReleaseMechanismI, error) {
	var rm types.ReleaseMechanismI
	if err := k.cdc.UnpackAny(any, &rm); err != nil {
		return nil, err
	}

	return rm, nil
}

func (k Keeper) RegisterReleaseMechanismToProject(ctx sdk.Context, projectId uint64, rm types.ReleaseMechanismI) error {
	project, err := k.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}

	// check if rm is in RegisteredRm, if yes then replace value. Return
	for i, any := range project.RegisteredRm {
		rmItem, err := k.ParseAnyReleaseMechanism(any)
		if err != nil {
			return err
		}

		if rmItem.GetType() == rm.GetType() {
			anyItem, err := codectypes.NewAnyWithValue(rm)
			if err != nil {
				return err
			}
			project.RegisteredRm[i] = anyItem

			if err := k.SetProject(ctx, project); err != nil {
				return err
			}
			return nil
		}
	}

	// check if allowed to register
	if project.ProjectStatus != types.PROJECT_INIT {
		return types.ErrCannotRegisterRm
	}

	any, err := codectypes.NewAnyWithValue(rm)
	if err != nil {
		return err
	}

	project.RegisteredRm = append(project.RegisteredRm, any)
	if err := k.SetProject(ctx, project); err != nil {
		return err
	}

	return nil
}

func (k Keeper) UnregisterReleaseMechanismFromProject(ctx sdk.Context, projectId uint64, rm_type string) error {
	project, err := k.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}

	// determine the location of rm in project.RegisteredRm and replace it with last element of array then return array with (len-1)
	for i, any := range project.RegisteredRm {
		rmItem, err := k.ParseAnyReleaseMechanism(any)
		if err != nil {
			return err
		}

		if rmItem.GetType() == rm_type {
			length := len(project.RegisteredRm)
			project.RegisteredRm[i] = project.RegisteredRm[length-1]
			project.RegisteredRm = project.RegisteredRm[:length-1]
			break
		}
	}

	if err := k.SetProject(ctx, project); err != nil {
		return err
	}

	return nil
}
