package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/nebula-labs/nebula/x/launchpad/types"
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

func (q Querier) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	context := sdk.UnwrapSDKContext(ctx)
	return &types.QueryParamsResponse{Params: q.Keeper.GetParams(context)}, nil
}

func (q Querier) Project(ctx context.Context, req *types.ProjectRequest) (*types.ProjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	context := sdk.UnwrapSDKContext(ctx)
	project, err := q.Keeper.GetProjectById(context, req.ProjectId)
	if err != nil {
		return nil, err
	}
	return &types.ProjectResponse{Project: &project}, nil
}

func (q Querier) TotalProject(ctx context.Context, req *types.TotalProjectRequest) (*types.TotalProjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	context := sdk.UnwrapSDKContext(ctx)
	projects := []types.Project{}
	store := context.KVStore(q.Keeper.storeKey)
	valStore := prefix.NewStore(store, types.KeyPrefixProject)

	pageRes, err := query.FilteredPaginate(valStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		project := types.Project{}
		err := project.Unmarshal(value)
		if err != nil {
			return false, err
		}
		projects = append(projects, project)

		return true, nil
	})

	if err != nil {
		return nil, err
	}
	return &types.TotalProjectResponse{Projects: projects, Pagination: pageRes}, nil
}

func (q Querier) ProjectBalances(ctx context.Context, req *types.ProjectBalancesRequest) (*types.ProjectBalancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	context := sdk.UnwrapSDKContext(ctx)

	// get project
	project, err := q.Keeper.GetProjectById(context, req.ProjectId)
	if err != nil {
		return nil, err
	}

	balances := q.bankKeeper.GetAllBalances(context, sdk.AccAddress(project.ProjectAddress))

	return &types.ProjectBalancesResponse{
		Balances: balances,
	}, nil
}
