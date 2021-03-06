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
		case MsgPrivateEventParticipate:
			return handelMsgPrivateEventParticipate(ctx, k, msg)
		case MsgPrivateEventValidate:
			return handelMsgPrivateEventValidate(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handelMsgPrivateCreateEvent(ctx sdk.Context, k Keeper, msg MsgPrivateCreateEvent) (*sdk.Result, error) {
	var event = types.CreateEvent{
		EventId:  msg.EventId,
		EndTime:  msg.EndTime,
		Question: msg.Question,
		Answers:  msg.Answers,
		Winner:   msg.Winner,
		Loser:    msg.Loser,
		Owner:    msg.Owner,
	}

	k.SetPrivateEvent(ctx, event)

	return &sdk.Result{}, nil
}

func handelMsgPrivateEventParticipate(ctx sdk.Context, k Keeper, msg MsgPrivateEventParticipate) (*sdk.Result, error) {
	var part = types.Participate{
		Participant:  msg.Participant,
		Answer:       msg.Answer,
		AnswerNumber: msg.AnswerNumber,
		Date:         msg.Date,
		EventId:      msg.EventId,
	}

	k.Participate(ctx, part)

	return &sdk.Result{}, nil
}

func handelMsgPrivateEventValidate(ctx sdk.Context, k Keeper, msg MsgPrivateEventValidate) (*sdk.Result, error) {

	var valid = types.Validate{
		Expert:       msg.Expert,
		Answer:       msg.Answer,
		AnswerNumber: msg.AnswerNumber,
		Date:         msg.Date,
		EventId:      msg.EventId,
	}

	k.Validate(ctx, valid)

	var finalAnswer = types.FinalAnswer{
		EventId:           msg.EventId,
		FinalAnswer:       msg.Answer,
		FinalAnswerNumber: msg.AnswerNumber,
	}
	k.SetFinalAnswer(ctx, finalAnswer)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.TypeEventFinish,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributePrivateEventId, strconv.Itoa(msg.EventId)),
			sdk.NewAttribute(types.AttributePrivateEventFinalAnswer, msg.Answer),
			sdk.NewAttribute(types.AttributePrivateEventFinalAnswerNumber, strconv.Itoa(msg.AnswerNumber)),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil

}
