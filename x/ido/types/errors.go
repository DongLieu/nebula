package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ido module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")

	ErrNotImplemented = sdkerrors.Register(ModuleName, 60, "function not implemented")

	ErrNotEnoughFunds                 = sdkerrors.Register(ModuleName, 1, "address doesn't have enough funds")
	ErrNoSupportedDenoms              = sdkerrors.Register(ModuleName, 3, "denom is not supported as payment denom")
	ErrNotEnoughIDOTokens             = sdkerrors.Register(ModuleName, 4, "there is not enough token in ido for purchase")
	ErrOutOfBoundPurchase             = sdkerrors.Register(ModuleName, 5, "you are purchasing amount not within limit")
	ErrStartTimeSmallerThanNow        = sdkerrors.Register(ModuleName, 6, "start time must greater than now")
	ErrProjectOwnerNotAllowedToCommit = sdkerrors.Register(ModuleName, 7, "project owner is not allowed to participate in their own project")
	ErrIDOIsNotActive                 = sdkerrors.Register(ModuleName, 8, "cannot commit because IDO is not active")
	ErrIDOAlreadyExist                = sdkerrors.Register(ModuleName, 9, "cannot enable ido of the same ID")
)
