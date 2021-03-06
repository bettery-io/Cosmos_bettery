package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgPrivateCreateEvent{}, "privateevent/create", nil)
	cdc.RegisterConcrete(MsgPrivateEventParticipate{}, "privateevent/participate", nil)
	cdc.RegisterConcrete(MsgPrivateEventValidate{}, "privateevent/validate", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
