package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/nebula-labs/nebula/x/launchpad/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (server msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProjectRequest) (*types.MsgCreateProjectResponse, error) {
	// get ctx SDK context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get project owner
	projectOwner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	// invoke logic CreateProject
	project_id, err := server.Keeper.CreateProject(ctx, projectOwner, msg)
	if err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		// an event to signify a project created
		sdk.NewEvent(
			types.TypeProjectCreated,
			sdk.NewAttribute(types.AttributeProjectID, strconv.FormatUint(project_id, 10)),
		),
		// an event to signify the event comes from which module and which signer
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	// return result to gRPC server
	return &types.MsgCreateProjectResponse{
		ProjectId: project_id,
	}, nil
}

func (server msgServer) DeleteProject(goCtx context.Context, msg *types.MsgDeleteProjectRequest) (*types.MsgDeleteProjectResponse, error) {
	// get ctx SDK context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if owner address is valid
	projectOwner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	// invoke logic CreateProject
	if err := server.Keeper.DeleteProject(ctx, projectOwner, msg); err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		// an event to signify a project created
		sdk.NewEvent(
			types.TypeProjectDeleted,
			sdk.NewAttribute(types.AttributeProjectID, strconv.FormatUint(msg.ProjectId, 10)),
		),
		// an event to signify the event comes from which module and which signer
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	return &types.MsgDeleteProjectResponse{}, nil
}

func (server msgServer) WithdrawAllTokens(goCtx context.Context, msg *types.MsgWithdrawAllTokensRequest) (*types.MsgWithdrawAllTokensResponse, error) {
	// get ctx SDK context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if owner address is valid
	projectOwner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	// invoke logic CreateProject
	if err := server.Keeper.WithdrawTokens(ctx, projectOwner, msg); err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		// an event to signify a project created
		sdk.NewEvent(
			types.TypeWithdrawTokens,
			sdk.NewAttribute(types.AttributeProjectID, strconv.FormatUint(msg.ProjectId, 10)),
		),
		// an event to signify the event comes from which module and which signer
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	return &types.MsgWithdrawAllTokensResponse{}, nil
}
