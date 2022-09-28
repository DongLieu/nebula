package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	launchpadtypes "github.com/nebula-labs/nebula/x/launchpad/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&IDO{}, "nebula/ido/ido", nil)
	cdc.RegisterConcrete(&MsgEnableIDORequest{}, "nebula/ido/enable-ido", nil)
	cdc.RegisterConcrete(&MsgCommitParticipationRequest{}, "nebula/ido/commit-participation", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"nebula.launchpad.ReleaseMechanismI",
		(*launchpadtypes.ReleaseMechanismI)(nil),
		&IDO{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgEnableIDORequest{},
		&MsgCommitParticipationRequest{},
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
