package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/nebula-labs/nebula/x/ido/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{Keeper: k}
}

func (q Querier) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: q.GetParams(ctx)}, nil
}

func (q Querier) IDO(goCtx context.Context, req *types.IDORequest) (*types.IDOResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	ido, err := q.GetIDOByID(ctx, req.ProjectId)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no such ido related to project id = %d", req.ProjectId))
	}

	return &types.IDOResponse{
		Ido: &ido,
	}, nil
}

func (q Querier) TotalIDO(ctx context.Context, req *types.TotalIDORequest) (*types.TotalIDOResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	context := sdk.UnwrapSDKContext(ctx)
	idos := []types.IDO{}
	store := context.KVStore(q.Keeper.storeKey)
	valStore := prefix.NewStore(store, types.KeyPrefixIDO)

	pageRes, err := query.FilteredPaginate(valStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		ido := types.IDO{}
		err := ido.Unmarshal(value)
		if err != nil {
			return false, err
		}
		idos = append(idos, ido)

		return true, nil
	})

	if err != nil {
		return nil, err
	}
	return &types.TotalIDOResponse{Idos: idos, Pagination: pageRes}, nil
}
