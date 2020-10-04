package keeper

import (
	"fmt"
	"strconv"

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
	var event types.EventInfo
	eventId := path[0]
	id, err := strconv.Atoi(eventId)
	if err != nil {
		return nil, err
	}
	eventItself, err := k.GetPrivateEventById(ctx, id)
	if err != nil {
		return nil, err
	}

	event.Question = eventItself.Question
	event.Answers = eventItself.Answers
	event.EndTime = eventItself.EndTime
	event.EventId = eventItself.EventId
	event.Loser = eventItself.Loser
	event.Winner = eventItself.Question
	event.Owner = eventItself.Owner

	event.FinalAnswer = "TO DO" // TO DO

	participants := k.GetPartIteratorByEventId(ctx, id)

	for ; participants.Valid(); participants.Next() {
		partWallet := removePrefixFromHash(participants.Key(), []byte(types.ParticipantPrefix+strconv.Itoa(event.EventId)))
		fmt.Println(string(partWallet))
		part, err := k.getParticipantById(ctx, event.EventId, string(partWallet))
		if err != nil {
			return
		}

		event.Participants = append(event.Participants, part)
	}

	res, err = codec.MarshalJSONIndent(k.cdc, event)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func removePrefixFromHash(key []byte, prefix []byte) (hash []byte) {
	hash = key[len(prefix):]
	return hash
}
