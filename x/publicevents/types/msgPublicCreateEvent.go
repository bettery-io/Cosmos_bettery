package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPublicCreateEvent{}

type MsgPublicCreateEvent struct {
	EventId         int            `json:"eventId"`
	EndTime         uint           `json:"endTime"`
	Owner           sdk.AccAddress `json:"owner"`
	CurrencyType    string         `json:"currencyType"`
	ValidatorAmount uint64         `json:"validatorAmount"`
	Question        string         `json:"question"`
	Answers         []string       `json:"answers"`
}

// NewMsgCreateEvent creates a new MsgPublicCreateEvent instance
func NewMsgPublicCreateEvent(
	_eventId int,
	_endTime uint,
	_owner sdk.AccAddress,
	_currencyType string,
	_validatorAmount uint64,
	_question string,
	_answers []string,
) MsgPublicCreateEvent {
	return MsgPublicCreateEvent{
		EventId:         _eventId,
		EndTime:         _endTime,
		Owner:           _owner,
		CurrencyType:    _currencyType,
		ValidatorAmount: _validatorAmount,
		Question:        _question,
		Answers:         _answers,
	}
}

const CreatePublicEventConst = "CreatePublicEvent"

// nolint
func (msg MsgPublicCreateEvent) Route() string { return RouterKey }
func (msg MsgPublicCreateEvent) Type() string  { return CreatePublicEventConst }
func (msg MsgPublicCreateEvent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgPublicCreateEvent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgPublicCreateEvent) ValidateBasic() error {
	if msg.EventId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "event id can't be empty")
	}
	if msg.EndTime == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "end time can't be empty")
	}
	if len(msg.Question) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "question can't be empty")
	}
	if len(msg.Answers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "answer can't be empty")
	}
	if len(msg.CurrencyType) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "currency type can't be empty")
	}
	if msg.ValidatorAmount == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validator amount can't be empty")
	}
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address owner can't be empty")
	}
	return nil
}
