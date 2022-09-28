package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCommitParticipationRequest{}

// msg types
const (
	TypeMsgCommitParticipationRequest = "commit_participation"
)

func NewMsgCommitParticipationRequest(participant string, projectId uint64, tokenCommit sdk.Coins) *MsgCommitParticipationRequest {
	return &MsgCommitParticipationRequest{
		Participant: participant,
		ProjectId:   projectId,
		TokenCommit: tokenCommit,
	}
}

func (msg *MsgCommitParticipationRequest) Route() string {
	return RouterKey
}

func (msg *MsgCommitParticipationRequest) Type() string {
	return TypeMsgEnableIDORequest
}

func (msg *MsgCommitParticipationRequest) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Participant)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgCommitParticipationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCommitParticipationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Participant)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
