package privateevents

import (
	"fmt"
	"strconv"

	"github.com/VoroshilovMax/Bettery/x/privateevents/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the privateevents type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgPrivateCreateEvent:
			return handelMsgPrivateCreateEvent(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handelMsgPrivateCreateEvent(ctx sdk.Context, k Keeper, msg MsgPrivateCreateEvent) (*sdk.Result, error) {
	var event = types.CreateEvent{
		EventId:   msg.EventId,
		StartTime: msg.StartTime,
		Question:  msg.Question,
		Answers:   msg.Answers,
		Winner:    msg.Winner,
		Loser:     msg.Loser,
		Owner:     msg.Owner,
	}
	// error handel don't work

	_, err := k.GetPrivateEvent(ctx, strconv.Itoa(int(event.EventId)))
	if err == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Quiz already exists")
	}
	k.SetPrivateEvent(ctx, event)

	// ctx.EventManager().EmitEvents(sdk.Events{
	// 	sdk.NewEvent(
	// 		sdk.EventTypeMessage,
	// 		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
	// 	sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddress.String()),
	// ),
	// sdk.NewEvent(
	// 	types.EventTypeCreateClaim,
	// 	sdk.NewAttribute(types.AttributeKeyEthereumSender, msg.EthereumSender.String()),
	// 	sdk.NewAttribute(types.AttributeKeyCosmosReceiver, msg.CosmosReceiver.String()),
	// 	sdk.NewAttribute(types.AttributeKeyAmount, strconv.FormatInt(msg.Amount, 10)),
	// 	sdk.NewAttribute(types.AttributeKeySymbol, msg.Symbol),
	// 	sdk.NewAttribute(types.AttributeKeyTokenContract, msg.TokenContractAddress.String()),
	// 	sdk.NewAttribute(types.AttributeKeyClaimType, msg.ClaimType.String()),
	// ),
	// sdk.NewEvent(
	// 	types.EventTypeProphecyStatus,
	// 	sdk.NewAttribute(types.AttributeKeyStatus, status.Text.String()),
	//		),
	//	})

	//return &sdk.Result{Events: ctx.EventManager().Events()}, nil
	return &sdk.Result{}, nil
}
