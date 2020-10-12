package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPrivateCreateEvent{}

type MsgPrivateCreateEvent struct {
	EventId  int            `json:"event_id"`
	EndTime  uint           `json:"end_time"`
	Question string         `json:"question"`
	Answers  []string       `json:"answers"`
	Winner   string         `json:"winner"`
	Loser    string         `json:"loser"`
	Owner    sdk.AccAddress `json:"owner"`
}

// NewMsgCreateEvent creates a new MsgPrivateCreateEvent instance
func NewMsgPrivateCreateEvent(
	_eventId int,
	_endTime uint,
	_question string,
	_answers []string,
	_winner string,
	_loser string,
	_owner sdk.AccAddress,
) MsgPrivateCreateEvent {
	return MsgPrivateCreateEvent{
		EventId:  _eventId,
		EndTime:  _endTime,
		Question: _question,
		Answers:  _answers,
		Winner:   _winner,
		Loser:    _loser,
		Owner:    _owner,
	}
}

const CreatePrivateEventConst = "CreatePrivateEvent"

// nolint
func (msg MsgPrivateCreateEvent) Route() string { return RouterKey }
func (msg MsgPrivateCreateEvent) Type() string  { return CreatePrivateEventConst }
func (msg MsgPrivateCreateEvent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgPrivateCreateEvent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgPrivateCreateEvent) ValidateBasic() error {
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
	if len(msg.Winner) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "winner can't be empty")
	}
	if len(msg.Loser) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "loser can't be empty")
	}
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address owner can't be empty")
	}
	return nil
}
