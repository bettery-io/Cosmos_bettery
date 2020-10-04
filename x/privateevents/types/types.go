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

type CreateEvent struct {
	EventId  int            `json:"event_id"`
	EndTime  uint           `json:"end_time"`
	Question string         `json:"question"`
	Answers  []string       `json:"answers"`
	Winner   string         `json:"winner"`
	Loser    string         `json:"loser"`
	Owner    sdk.AccAddress `json:"owner"`
}

func (s CreateEvent) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`EventId: %v
	EndTime: %v
	Question: %v
	Answers: %v
	Winner: %v
	Loser: %v
	Owner: %v
	`,
		s.EventId,
		s.EndTime,
		s.Question,
		s.Answers,
		s.Winner,
		s.Loser,
		s.Owner,
	))
}
