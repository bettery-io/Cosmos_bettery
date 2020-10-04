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

type Participate struct {
	Participant sdk.AccAddress `json:"participant"`
	Answer      string         `json:"answer"`
	Date        uint           `json:"date"`
	EventId     int            `json:"event_id"`
}

func (s Participate) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`Participant: %v
		Answer: %v
		Date: %v
		EventId: %v
	`,
		s.Participant,
		s.Answer,
		s.Date,
		s.EventId,
	))
}

type Validate struct {
	Expert  sdk.AccAddress `json:"expert"`
	Answer  string         `json:"answer"`
	Date    uint           `json:"date"`
	EventId int            `json:"event_id"`
}

func (s Validate) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`Expert: %v
		Answer: %v
		Date: %v
		EventId: %v
	`,
		s.Expert,
		s.Answer,
		s.Date,
		s.EventId,
	))
}

type EventInfo struct {
	EventId      int            `json:"event_id"`
	EndTime      uint           `json:"end_time"`
	Question     string         `json:"question"`
	Answers      []string       `json:"answers"`
	Winner       string         `json:"winner"`
	Loser        string         `json:"loser"`
	Owner        sdk.AccAddress `json:"owner"`
	FinalAnswer  string         `json:"final_answer"`
	Participants []Participate  `json:"participants"`
	Validator    Validate       `json:"validator"`
}
