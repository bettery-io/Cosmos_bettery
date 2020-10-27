package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateCoin{}

type MsgCreateCoin struct {
	Symbol        string         `json:"symbol"`
	Owner         sdk.AccAddress `json:"owner"`
	Amount        uint           `json:"amount"`
	SenderAddress string         `json:"sender_address"`
}

// NewMsgCreateEvent creates a new MsgPrivateCreateEvent instance
func NewMsgCreateCoin(
	_symbol string,
	_owner sdk.AccAddress,
	_amount uint,
	_senderAddress string,
) MsgCreateCoin {
	return MsgCreateCoin{
		Symbol:        _symbol,
		Owner:         _owner,
		Amount:        _amount,
		SenderAddress: _senderAddress,
	}
}

const CreateCoinConst = "CreateCoinConst"

// nolint
func (msg MsgCreateCoin) Route() string { return RouterKey }
func (msg MsgCreateCoin) Type() string  { return CreateCoinConst }
func (msg MsgCreateCoin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateCoin) ValidateBasic() error {
	if len(msg.Symbol) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "symbol can't be empty")
	}
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "owner can't be empty")
	}
	if msg.Amount == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount can't be less that 0")
	}
	if len(msg.SenderAddress) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Sender address can't be empty")
	}
	return nil
}
