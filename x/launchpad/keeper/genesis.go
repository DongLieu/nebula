package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nebula-labs/nebula/x/launchpad/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	if genState.GlobalProjectIdStart != uint64(len(genState.Projects)) {
		panic(sdkerrors.Wrapf(
			types.ErrGenesisProjectIDDoesNotMatchProjectArray,
			"Check genesis again: attempting to init genesis with unequal project array length and GlobalProjectIdStart",
		))
	}
	k.SetNextProjectID(ctx, genState.GlobalProjectIdStart)

	// setup initial projects
	for _, any := range genState.Projects {
		if err := k.SetProject(ctx, *any); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	projects, err := k.GetAllProjects(ctx)
	if err != nil {
		panic(err)
	}
	genesis.Projects = projects
	genesis.GlobalProjectIdStart = uint64(len(projects))

	return genesis
}
