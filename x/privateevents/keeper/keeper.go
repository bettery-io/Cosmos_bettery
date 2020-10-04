package keeper

import (
	"fmt"
	"strconv"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/VoroshilovMax/Bettery/x/privateevents/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// Keeper of the privateevents store
type Keeper struct {
	CoinKeeper bank.Keeper
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

// NewKeeper creates a privateevents keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, coinKeeper bank.Keeper) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		CoinKeeper: coinKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Get returns the pubkey from the adddress-pubkey relation
func (k Keeper) GetPrivateEventById(ctx sdk.Context, eventId int) (types.CreateEvent, error) {
	store := ctx.KVStore(k.storeKey)
	var event types.CreateEvent
	key := []byte(types.EventPrefix + strconv.Itoa(int(eventId)))
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(key), &event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (k Keeper) SetPrivateEvent(ctx sdk.Context, event types.CreateEvent) {
	eventId := event.EventId
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(event)
	key := []byte(types.EventPrefix + strconv.Itoa(int(eventId)))
	store.Set([]byte(key), bz)
}

func (k Keeper) Participate(ctx sdk.Context, event types.Participate) {
	keys := event.Participant.String()
	eventId := strconv.Itoa(int(event.EventId))
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(event)
	key := []byte(types.ParticipantPrefix + eventId + keys)
	store.Set(key, bz)
}

func (k Keeper) getParticipantById(ctx sdk.Context, eventId string, partWallet string) (types.Participate, error) {
	store := ctx.KVStore(k.storeKey)
	var patr types.Participate
	byteKey := []byte(types.ParticipantPrefix + eventId + partWallet)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &patr)
	if err != nil {
		return patr, err
	}
	return patr, nil
}

func (k Keeper) GetPartIteratorByEventId(ctx sdk.Context, eventId string) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.ParticipantPrefix+eventId))
}

func (k Keeper) Validate(ctx sdk.Context, event types.Validate) {
	keys := event.Expert.String()
	eventId := strconv.Itoa(int(event.EventId))
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(event)
	key := []byte(types.ValidatorPrefix + eventId + keys)
	store.Set(key, bz)
}

func (k Keeper) getValidatorById(ctx sdk.Context, eventId string, validWallet string) (types.Validate, error) {
	store := ctx.KVStore(k.storeKey)
	var valid types.Validate
	byteKey := []byte(types.ValidatorPrefix + eventId + validWallet)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &valid)
	if err != nil {
		return valid, err
	}
	return valid, nil
}

func (k Keeper) GetValidIteratorByEventId(ctx sdk.Context, eventId string) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.ValidatorPrefix+eventId))
}
