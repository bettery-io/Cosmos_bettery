package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/VoroshilovMax/Bettery/x/coinmaker/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// Keeper of the publicevents store
type Keeper struct {
	CoinKeeper   bank.Keeper
	storeKey     sdk.StoreKey
	supplyKeeper types.SupplyKeeper
	cdc          *codec.Codec
}

// NewKeeper creates a privateevents keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, supplyKeeper types.SupplyKeeper, coinKeeper bank.Keeper) Keeper {
	return Keeper{
		cdc:          cdc,
		storeKey:     storeKey,
		supplyKeeper: supplyKeeper,
		CoinKeeper:   coinKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreateCoin(ctx sdk.Context, event types.CreateCoin) error {
	var coins sdk.Coins
	symbol := event.Symbol
	coins = sdk.Coins{sdk.NewInt64Coin(symbol, int64(event.Amount))}
	err := k.supplyKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return err
	}
	if err := k.supplyKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, event.Owner, coins,
	); err != nil {
		panic(err)
	}

	return nil
}
