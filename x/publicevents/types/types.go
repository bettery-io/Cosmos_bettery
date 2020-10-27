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
	EventId         int            `json:"eventId"`
	EndTime         uint           `json:"endTime"`
	Owner           sdk.AccAddress `json:"owner"`
	CurrencyType    string         `json:"currencyType"`
	ValidatorAmount uint64         `json:"validatorAmount"`
	Question        string         `json:"question"`
	Answers         []string       `json:"answers"`
}

func (s CreateEvent) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`EventId: %v
		EndTime: %v
		Owner: %v
		CurrencyType: %v
		ValidatorAmount: %v
		Question: %v
		Answers: %v
	`,
		s.EventId,
		s.EndTime,
		s.Owner,
		s.CurrencyType,
		s.ValidatorAmount,
		s.Question,
		s.Answers,
	))
}

type Participante struct {
	Participant  sdk.AccAddress `json:"participant"`
	Answer       string         `json:"answer"`
	AnswerNumber int            `json:"answer_number"`
	EventId      int            `json:"event_id"`
	Amount       sdk.Coins      `json:"amount"`
	Date         uint           `json:"date"`
}

func (s Participante) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`Participant: %v
		Answer: %v
		AnswerNumber: %v
		EventId: %v
		Amount: %v
		Date: %v
	`,
		s.Participant,
		s.Answer,
		s.AnswerNumber,
		s.EventId,
		s.Amount,
		s.Date,
	))
}

type Validate struct {
	Expert       sdk.AccAddress `json:"expert"`
	Answer       string         `json:"answer"`
	AnswerNumber int            `json:"answer_number"`
	Date         uint           `json:"date"`
	EventId      int            `json:"event_id"`
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
