package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*ReleaseMechanismI)(nil), nil)
	cdc.RegisterConcrete(&MsgCreateProjectRequest{}, "nebula/launchpad/create-project", nil)
	cdc.RegisterConcrete(&MsgDeleteProjectRequest{}, "nebula/launchpad/delete-project", nil)
	cdc.RegisterConcrete(&MsgWithdrawAllTokensRequest{}, "nebula/launchpad/withdraw-all-tokens", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"nebula.launchpad.ReleaseMechanismI",
		(*ReleaseMechanismI)(nil),
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateProjectRequest{},
		&MsgDeleteProjectRequest{},
		&MsgWithdrawAllTokensRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(Amino)
	Amino.Seal()
}
