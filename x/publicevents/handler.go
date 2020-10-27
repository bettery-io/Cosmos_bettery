package publicevents

import (
	"fmt"

	"github.com/VoroshilovMax/Bettery/x/publicevents/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the publicevents type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgPublicCreateEvent:
			return handelMsgPublicCreateEvent(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handelMsgPublicCreateEvent(ctx sdk.Context, k Keeper, msg MsgPublicCreateEvent) (*sdk.Result, error) {
	var event = types.CreateEvent{
		EventId:         msg.EventId,
		EndTime:         msg.EndTime,
		Owner:           msg.Owner,
		CurrencyType:    msg.CurrencyType,
		ValidatorAmount: msg.ValidatorAmount,
		Question:        msg.Question,
		Answers:         msg.Answers,
	}
	k.SetPublicEvent(ctx, event)
	return &sdk.Result{}, nil
}
