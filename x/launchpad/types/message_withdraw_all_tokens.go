package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgWithdrawAllTokensRequest{}

// msg types
const (
	TypeMsgWithdrawAllTokensRequest = "withdraw_all_tokens"
)

func NewMsgWithdrawAllTokensRequest(owner string, id uint64) *MsgWithdrawAllTokensRequest {
	return &MsgWithdrawAllTokensRequest{
		Owner:     owner,
		ProjectId: id,
	}
}

func (msg *MsgWithdrawAllTokensRequest) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawAllTokensRequest) Type() string {
	return TypeMsgDeleteProjectRequest
}

func (msg *MsgWithdrawAllTokensRequest) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgWithdrawAllTokensRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawAllTokensRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
