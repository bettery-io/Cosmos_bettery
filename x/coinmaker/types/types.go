package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BankKeeper interface {
	SubtractCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, error)
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
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
