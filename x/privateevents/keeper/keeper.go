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
func (k Keeper) GetPrivateEvent(ctx sdk.Context, eventId string) (types.CreateEvent, error) {
	store := ctx.KVStore(k.storeKey)
	var event types.CreateEvent
	byteKey := []byte(eventId)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &event)
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

// func (k Keeper) delete(ctx sdk.Context, key string) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Delete([]byte(key))
// }
