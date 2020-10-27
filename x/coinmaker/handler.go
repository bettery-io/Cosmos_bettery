package coinmaker

import (
	"fmt"

	"github.com/VoroshilovMax/Bettery/x/coinmaker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//	"github.com/VoroshilovMax/Bettery/x/coinmaker/types"
)

// NewHandler creates an sdk.Handler for all the coinmaker type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateCoin:
			return handelMsgCreateCoin(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handelMsgCreateCoin(ctx sdk.Context, k Keeper, msg MsgCreateCoin) (*sdk.Result, error) {
	var event = types.CreateCoin{
		Symbol:        msg.Symbol,
		Amount:        msg.Amount,
		Owner:         msg.Owner,
		SenderAddress: msg.SenderAddress,
	}

	fmt.Println(event)

	return &sdk.Result{}, nil
}
