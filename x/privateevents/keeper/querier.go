package keeper

import (
	//	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	//	"github.com/cosmos/cosmos-sdk/client"
	"github.com/VoroshilovMax/Bettery/x/privateevents/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for privateevents clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryGetSinglePrivateEvent:
			return queryGetSinglePrivateEvent(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown privateevents query endpoint")
		}
	}
}

func queryGetSinglePrivateEvent(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	eventId := path[0]
	event, err := k.GetPrivateEvent(ctx, eventId)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, event)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
