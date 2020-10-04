package privateevents

import (
	"github.com/VoroshilovMax/Bettery/x/privateevents/keeper"
	"github.com/VoroshilovMax/Bettery/x/privateevents/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	//QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper                     = keeper.NewKeeper
	NewQuerier                    = keeper.NewQuerier
	RegisterCodec                 = types.RegisterCodec
	NewGenesisState               = types.NewGenesisState
	DefaultGenesisState           = types.DefaultGenesisState
	ValidateGenesis               = types.ValidateGenesis
	NewMsgPrivateCreateEvent      = types.NewMsgPrivateCreateEvent
	NewMsgPrivateEventParticipate = types.NewMsgPrivateEventParticipate
	NewMsgPrivateEventValidate    = types.NewMsgPrivateEventValidate

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	MsgPrivateCreateEvent      = types.MsgPrivateCreateEvent
	MsgPrivateEventParticipate = types.MsgPrivateEventParticipate
	MsgPrivateEventValidate    = types.MsgPrivateEventValidate
)
