package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"
)

type BankKeeper interface {
	SubtractCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, error)
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

// SupplyKeeper defines the expected supply keeper
type SupplyKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	SetModuleAccount(sdk.Context, supplyexported.ModuleAccountI)
}

type CreateCoin struct {
	Symbol        string         `json:"symbol"`
	Owner         sdk.AccAddress `json:"owner"`
	Amount        uint           `json:"amount"`
	SenderAddress string         `json:"sender_address"`
}

func (s CreateCoin) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`Symbol: %v
		Owner: %v
		Amount: %v
		SenderAddress: %v
	`,
		s.Symbol,
		s.Owner,
		s.Amount,
		s.SenderAddress,
	))
}
