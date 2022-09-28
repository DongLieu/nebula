package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nebula-labs/nebula/x/ido/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: *keeper}
}

var _ types.MsgServer = msgServer{}

func (server msgServer) EnableIDO(goCtx context.Context, msg *types.MsgEnableIDORequest) (*types.MsgEnableIDOResponse, error) {
	// get ctx SDK context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get project owner
	project_owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	// invoke logic EnableIDO
	err = server.Keeper.EnableIDO(ctx, project_owner, msg)
	if err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		// an event to signify a project created
		sdk.NewEvent(
			types.TypeDistributionTokenAdded,
			sdk.NewAttribute(types.AttributeProjectID, strconv.FormatUint(msg.ProjectId, 10)),
		),
		// an event to signify the event comes from which module and which signer
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	// return result to gRPC server
	return &types.MsgEnableIDOResponse{}, nil
}

func (server msgServer) CommitParticipation(goCtx context.Context, msg *types.MsgCommitParticipationRequest) (*types.MsgCommitParticipationResponse, error) {
	// get ctx SDK context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get participant
	participant, err := sdk.AccAddressFromBech32(msg.Participant)
	if err != nil {
		return nil, err
	}

	// invoke logic EnableIDO
	err = server.Keeper.CommitParticipation(ctx, participant, msg)
	if err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		// an event to signify a project created
		sdk.NewEvent(
			types.TypeDistributionTokenAdded,
			sdk.NewAttribute(types.AttributeParticipantAddress, msg.Participant),
			sdk.NewAttribute(types.AttributeProjectID, strconv.FormatUint(msg.ProjectId, 10)),
		),
		// an event to signify the event comes from which module and which signer
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Participant),
		),
	})

	// return result to gRPC server
	return &types.MsgCommitParticipationResponse{}, nil
}
