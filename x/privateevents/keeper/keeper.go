package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/VoroshilovMax/Bettery/x/privateevents/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper of the privateevents store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	paramspace types.ParamSubspace
}

// NewKeeper creates a privateevents keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramspace types.ParamSubspace) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Get returns the pubkey from the adddress-pubkey relation
func (k Keeper) Get(ctx sdk.Context, key string) (types.CreateEvent, error) {
	store := ctx.KVStore(k.storeKey)
	var item types.CreateEvent
	byteKey := []byte(key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (k Keeper) set(ctx sdk.Context, key string, value types.CreateEvent) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(value)
	store.Set([]byte(key), bz)
}

func (k Keeper) delete(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}
