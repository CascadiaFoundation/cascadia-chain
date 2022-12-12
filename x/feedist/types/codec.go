package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&RegisterFeedistProposal{},
		&CancelFeedistProposal{},
	)
}
