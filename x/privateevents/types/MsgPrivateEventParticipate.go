package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPrivateEventParticipate{}

type MsgPrivateEventParticipate struct {
	Answer       string         `json:"answer"`
	Date         uint           `json:"date"`
	AnswerNumber int            `json:"answerNumber"`
	Participant  sdk.AccAddress `json:"owner"`
	EventId      int            `json:"event_id"`
}

// NewMsgCreateEvent creates a new MsgPrivateCreateEvent instance
func NewMsgPrivateEventParticipate(
	_answer string,
	_date uint,
	_answerNumber int,
	_participant sdk.AccAddress,
	_eventId int,
) MsgPrivateEventParticipate {
	return MsgPrivateEventParticipate{
		Answer:       _answer,
		Date:         _date,
		AnswerNumber: _answerNumber,
		Participant:  _participant,
		EventId:      _eventId,
	}
}

const PrivateEventParticipateConst = "PrivateEventParticipate"

// nolint
func (msg MsgPrivateEventParticipate) Route() string { return RouterKey }
func (msg MsgPrivateEventParticipate) Type() string  { return PrivateEventParticipateConst }
func (msg MsgPrivateEventParticipate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Participant)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgPrivateEventParticipate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgPrivateEventParticipate) ValidateBasic() error {
	if len(msg.Answer) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "answer can't be empty")
	}
	if msg.Date == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "date can't be empty")
	}
	if msg.Participant.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address participant can't be empty")
	}
	if msg.EventId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "event id can't be empty")
	}
	return nil
}
