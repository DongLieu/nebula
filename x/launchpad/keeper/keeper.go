package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/nebula-labs/nebula/x/launchpad/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		paramstore paramtypes.Subspace
		hooks      types.LaunchpadHooks

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetModuleAccountAddress gets the address of module account
func (k Keeper) GetModuleAccountAddress(ctx sdk.Context) sdk.AccAddress {
	return k.accountKeeper.GetModuleAddress(types.ModuleName)
}

// ============ Project Helper Logic

// Get new project address
func (k Keeper) NewProjectAddress(projectID uint64) sdk.AccAddress {
	key := append([]byte("project"), sdk.Uint64ToBigEndian(projectID)...)
	return address.Module(types.ModuleName, key)
}

func (k Keeper) SetProjectActive(ctx sdk.Context, projectId uint64) error {
	project, err := k.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}

	if project.ProjectStatus == types.PROJECT_ACTIVE {
		return nil
	}

	if project.ProjectStatus != types.PROJECT_INIT {
		return types.ErrCannotModifyProject
	}

	project.ProjectStatus = types.PROJECT_ACTIVE
	err = k.SetProject(ctx, project)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) SetProjectEndable(ctx sdk.Context, projectId uint64) error {
	project, err := k.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}

	if project.ProjectStatus != types.PROJECT_ACTIVE {
		return types.ErrCannotModifyProject
	}

	// check if all rm has ended
	for _, any := range project.RegisteredRm {
		rmItem, err := k.ParseAnyReleaseMechanism(any)
		if err != nil {
			return err
		}

		if rmItem.GetReleaseMechanismStatus() != types.RM_ENDED {
			return types.ErrCannotModifyProject
		}
	}

	project.ProjectStatus = types.PROJECT_ENDED
	err = k.SetProject(ctx, project)
	if err != nil {
		return err
	}

	return nil
}

// ============ Hooks

// Set the gamm hooks.
func (k *Keeper) SetHooks(gh types.LaunchpadHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set gamm hooks twice")
	}

	k.hooks = gh

	return k
}
