package rest

import (
	"fmt"
	"net/http"

	helpers "github.com/VoroshilovMax/Bettery/x/privateevents/client/helpers"
	restTypes "github.com/VoroshilovMax/Bettery/x/privateevents/client/types"
	"github.com/VoroshilovMax/Bettery/x/privateevents/types"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/privateevent/create", createPrivateEvent(cliCtx)).Methods("POST")
	r.HandleFunc("/privateevent/participate", participatePrivateEvent(cliCtx)).Methods("POST")
	r.HandleFunc("/privateevent/validate", validatePrivateEvent(cliCtx)).Methods("POST")
}

// TO DO
// 1 check if exist event when create new event
// 2 check time to participate
// 3 check if participate whant to participate send time
// 4 check time for validator
// 5 build better struct for get event
// 6 finish event when validator did his job

func createPrivateEvent(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req restTypes.CreateEvent

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		onwer, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		validator, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// _, err := k.GetPrivateEventById(ctx, event.EventId)
		// if err == nil {
		// 	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Quiz already exists")
		// }

		msg := types.NewMsgPrivateCreateEvent(
			req.EventId,
			req.StartTime,
			req.Question,
			req.Answers,
			req.Winner,
			req.Loser,
			onwer,
			validator,
		)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

func participatePrivateEvent(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req restTypes.Participate

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		participant, err := sdk.AccAddressFromBech32(req.Participant)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		validator, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		answerNumber := helpers.GetAnswerNumber(req.EventId, req.Answer, w, r, cliCtx)
		if answerNumber == -1 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Can't find answer in event answers")
			return
		}

		fmt.Println(answerNumber)

		msg := types.NewMsgPrivateEventParticipate(
			req.Answer,
			req.Date,
			answerNumber,
			participant,
			validator,
			req.EventId,
		)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

func validatePrivateEvent(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req restTypes.Validate

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		expert, err := sdk.AccAddressFromBech32(req.Expert)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		validator, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		answerNumber := helpers.GetAnswerNumber(req.EventId, req.Answer, w, r, cliCtx)
		if answerNumber == -1 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Can't find answer in event answers")
			return
		}

		fmt.Println(answerNumber)

		msg := types.NewMsgPrivateEventValidate(
			req.Answer,
			req.Date,
			answerNumber,
			expert,
			validator,
			req.EventId,
		)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
