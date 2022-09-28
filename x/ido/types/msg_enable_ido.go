package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEnableIDORequest{}

// msg types
const (
	TypeMsgEnableIDORequest = "enable_ido"
)

func NewMsgEnableIDORequest(owner string, projectId uint64, tokenForDistribution sdk.Coins, tokenListingPrice sdk.Coins, allocationLimit []AllocationLimit, startTime time.Time) *MsgEnableIDORequest {
	return &MsgEnableIDORequest{
		Owner:                owner,
		ProjectId:            projectId,
		TokenForDistribution: tokenForDistribution,
		TokenListingPrice:    tokenListingPrice,
		AllocationLimit:      allocationLimit,
		StartTime:            startTime,
	}
}

func (msg *MsgEnableIDORequest) Route() string {
	return RouterKey
}

func (msg *MsgEnableIDORequest) Type() string {
	return TypeMsgEnableIDORequest
}

func (msg *MsgEnableIDORequest) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgEnableIDORequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableIDORequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
