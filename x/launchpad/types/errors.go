package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/launchpad module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")

	ErrInvalidProject                           = sdkerrors.Register(ModuleName, 1, "attempting to create an invalid project")
	ErrGenesisProjectIDDoesNotMatchProjectArray = sdkerrors.Register(ModuleName, 2, "attempting to init genesis with unequal project array length and GlobalProjectIdStart")
	ErrCannotModifyProject                      = sdkerrors.Register(ModuleName, 3, "cannot modify an active project")
	ErrCannotWithdrawTokens                     = sdkerrors.Register(ModuleName, 4, "not allowed to withdraw tokens")
	ErrCannotRegisterRm                         = sdkerrors.Register(ModuleName, 5, "not allowed to register release mechanism")
	ErrNotProjecOwner                           = sdkerrors.Register(ModuleName, 6, "not project owner")
	ErrCannotDeleteProject                      = sdkerrors.Register(ModuleName, 7, "not allowed to delete project")

	ErrNotImplemented = sdkerrors.Register(ModuleName, 60, "function not implemented")

	ErrProjectNotFound = sdkerrors.Register(ModuleName, 404, "project not found")
)
