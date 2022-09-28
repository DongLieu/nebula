package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nebula-labs/nebula/x/launchpad/types"

	"github.com/gogo/protobuf/proto"
	gogotypes "github.com/gogo/protobuf/types"
)

func (k Keeper) MarshalProject(project types.Project) ([]byte, error) {
	return proto.Marshal(&project)
}

func (k Keeper) UnmarshalProject(bz []byte) (types.Project, error) {
	var acc types.Project
	return acc, proto.Unmarshal(bz, &acc)
}

func (k Keeper) SetProject(ctx sdk.Context, project types.Project) error {
	bz, err := k.MarshalProject(project)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	projectKey := types.GetKeyPrefixProject(project.GetProjectId())
	store.Set(projectKey, bz)

	return nil
}

func (k Keeper) DeleteProjectById(ctx sdk.Context, projectId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetKeyPrefixProject(projectId))
}

// GetNextProjectIDAndIncrement returns the next project id, and increments the corresponding state entry.
func (k Keeper) GetNextProjectIDAndIncrement(ctx sdk.Context) uint64 {
	var projectID uint64
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyNextGlobalProjectID)
	if bz == nil {
		panic(fmt.Errorf("project has not been initialized -- Should have been done in InitGenesis"))
	} else {
		val := gogotypes.UInt64Value{}

		err := k.cdc.Unmarshal(bz, &val)
		if err != nil {
			panic(err)
		}

		projectID = val.GetValue()
	}

	k.SetNextProjectID(ctx, projectID+1)
	return projectID + 1
}

func (k Keeper) GetProjectById(ctx sdk.Context, projectId uint64) (types.Project, error) {
	store := ctx.KVStore(k.storeKey)
	projectKey := types.GetKeyPrefixProject(projectId)
	if !store.Has(projectKey) {
		return types.Project{}, fmt.Errorf("project with ID %d does not exist", projectId)
	}
	project, err := k.UnmarshalProject(store.Get(projectKey))
	if err != nil {
		return types.Project{}, err
	}

	return project, nil
}

func (k Keeper) GetAllProjects(ctx sdk.Context) (res []*types.Project, err error) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.KeyPrefixProject)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		bz := iter.Value()

		project, err := k.UnmarshalProject(bz)
		if err != nil {
			return nil, err
		}

		res = append(res, &project)
	}
	return res, nil
}

// SetNextProjectID sets next project ID.
func (k Keeper) SetNextProjectID(ctx sdk.Context, projectID uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: projectID})
	store.Set(types.KeyNextGlobalProjectID, bz)
}

func (k Keeper) IterateProjectInfo(ctx sdk.Context, fn func(index uint64, projectInfo types.Project) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixProject)
	defer iterator.Close()

	i := uint64(0)

	for ; iterator.Valid(); iterator.Next() {
		project := types.Project{}
		err := proto.Unmarshal(iterator.Value(), &project)
		if err != nil {
			panic(err)
		}
		stop := fn(i, project)

		if stop {
			break
		}
		i++
	}
}
