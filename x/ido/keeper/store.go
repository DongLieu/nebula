package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nebula-labs/nebula/x/ido/types"

	"github.com/gogo/protobuf/proto"
)

func (k Keeper) MarshalIDO(ido types.IDO) ([]byte, error) {
	return proto.Marshal(&ido)
}

func (k Keeper) UnmarshalIDO(bz []byte) (types.IDO, error) {
	var ido types.IDO
	if err := proto.Unmarshal(bz, &ido); err != nil {
		return ido, err
	}

	if ido.Entries == nil {
		ido.Entries = make(map[string]types.Entry)
	}

	return ido, nil
}

func (k Keeper) HasIDO(ctx sdk.Context, project_id uint64) bool {
	store := ctx.KVStore(k.storeKey)
	projectKey := types.GetKeyPrefixProject(project_id)
	return store.Has(projectKey)
}

func (k Keeper) SetIDO(ctx sdk.Context, ido types.IDO) error {
	bz, err := k.MarshalIDO(ido)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	projectKey := types.GetKeyPrefixProject(ido.ProjectId)
	store.Set(projectKey, bz)

	return nil
}

func (k Keeper) SetIDOAndRegisterLaunchpad(ctx sdk.Context, ido types.IDO) error {
	if err := k.SetIDO(ctx, ido); err != nil {
		return err
	}

	if err := k.launchpadKeeper.RegisterReleaseMechanismToProject(ctx, ido.ProjectId, &ido); err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetIDOByID(ctx sdk.Context, project_id uint64) (types.IDO, error) {
	store := ctx.KVStore(k.storeKey)
	projectKey := types.GetKeyPrefixProject(project_id)
	if !store.Has(projectKey) {
		return types.IDO{}, fmt.Errorf("IDO with ID %d does not exist", project_id)
	}

	IDO, err := k.UnmarshalIDO(store.Get(projectKey))
	if err != nil {
		return types.IDO{}, err
	}

	return IDO, nil
}

func (k Keeper) IterateIDO(ctx sdk.Context, fn func(index int64, ido types.IDO) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixIDO)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		ido := types.IDO{}
		err := proto.Unmarshal(iterator.Value(), &ido)
		if err != nil {
			panic(err)
		}
		stop := fn(i, ido)

		if stop {
			break
		}
		i++
	}
}

func (k Keeper) DeleteIDOById(ctx sdk.Context, projectId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetKeyPrefixProject(projectId))
}
