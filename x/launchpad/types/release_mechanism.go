package types

import (
	proto "github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	RM_INIT   = uint64(0)
	RM_ACTIVE = uint64(1)
	RM_ENDED  = uint64(2)
)

type ReleaseMechanismI interface {
	proto.Message

	GetId() uint64

	// get active status of release mechanism
	// 0: has not started
	// 1: active
	// 2: ended
	// 3: request delete
	GetReleaseMechanismStatus() uint64

	// get information on tokens for each release mechanism
	GetTokens() sdk.Coins

	// get information on which type this is
	GetType() string
}
