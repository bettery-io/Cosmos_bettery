package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CreateEvent struct {
	EventId   uint           `json:"event_id"`
	StartTime uint           `json:"start_time"`
	Question  string         `json:"question"`
	Answers   []string       `json:"answers"`
	Winner    string         `json:"winner"`
	Loser     string         `json:"loser"`
	Owner     sdk.AccAddress `json:"owner"`
}

func (s CreateEvent) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`EventId: %v
	StartTime: %v
	Question: %v
	Answers: %v
	Winner: %v
	Loser: %v
	Owner: %v
	`,
		s.EventId,
		s.StartTime,
		s.Question,
		s.Answers,
		s.Winner,
		s.Loser,
		s.Owner,
	))
}
