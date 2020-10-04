package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPrivateEventValidate{}

type MsgPrivateEventValidate struct {
	Answer       string         `json:"answer"`
	Date         uint           `json:"date"`
	AnswerNumber int            `json:"answerNumber"`
	Expert       sdk.AccAddress `json:"expert"`
	Validator    sdk.AccAddress `json:"validator"`
	EventId      int            `json:"event_id"`
}

// NewMsgCreateEvent creates a new MsgPrivateCreateEvent instance
func NewMsgPrivateEventValidate(
	_answer string,
	_date uint,
	_answerNumber int,
	_expert sdk.AccAddress,
	_validator sdk.AccAddress,
	_eventId int,
) MsgPrivateEventValidate {
	return MsgPrivateEventValidate{
		Answer:       _answer,
		Date:         _date,
		AnswerNumber: _answerNumber,
		Expert:       _expert,
		Validator:    _validator,
		EventId:      _eventId,
	}
}

const PrivateEventValidateConst = "PrivateEventValidate"

// nolint
func (msg MsgPrivateEventValidate) Route() string { return RouterKey }
func (msg MsgPrivateEventValidate) Type() string  { return PrivateEventValidateConst }
func (msg MsgPrivateEventValidate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Validator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgPrivateEventValidate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgPrivateEventValidate) ValidateBasic() error {
	if len(msg.Answer) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "answer can't be empty")
	}
	if msg.Date == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "date can't be empty")
	}
	if msg.AnswerNumber == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "answer number can't be empty")
	}
	if msg.Expert.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address expert can't be empty")
	}
	if msg.Validator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address validator can't be empty")
	}
	if msg.EventId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "event id can't be empty")
	}
	return nil
}
