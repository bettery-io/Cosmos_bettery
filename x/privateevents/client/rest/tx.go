package rest

import (
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
// 2 check time to participate
// 3 check if participant already participate
// 4 check time for validator
// 6 finish event when validator did his job
// 7 BUG when I sen first answer like paricipant
// 10 build registration with DB

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

		owner, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		eventAlreadyExist := helpers.CheckIfEventExist(req.EventId, w, r, cliCtx)
		if eventAlreadyExist {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Event with this id already exist")
			return
		}

		msg := types.NewMsgPrivateCreateEvent(
			req.EventId,
			req.EndTime,
			req.Question,
			req.Answers,
			req.Winner,
			req.Loser,
			owner,
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

		participant, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		answerNumber := helpers.GetAnswerNumber(req.EventId, req.Answer, w, r, cliCtx)
		if answerNumber == -1 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Can't find answer in event answers")
			return
		}

		date := helpers.CurrentEpochTime()

		msg := types.NewMsgPrivateEventParticipate(
			req.Answer,
			uint(date),
			answerNumber,
			participant,
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

		expert, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		answerIsKnown := helpers.AnswerIsKnown(req.EventId, w, r, cliCtx)
		if answerIsKnown {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Event is finish. Answer is alredy provided")
			return
		}

		answerNumber := helpers.GetAnswerNumber(req.EventId, req.Answer, w, r, cliCtx)
		if answerNumber == -1 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Can't find answer in event answers")
			return
		}

		date := helpers.CurrentEpochTime()

		msg := types.NewMsgPrivateEventValidate(
			req.Answer,
			uint(date),
			answerNumber,
			expert,
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
