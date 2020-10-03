package rest

import (
	// "bytes"

	"net/http"

	"github.com/VoroshilovMax/Bettery/x/privateevents/types"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/cosmos/cosmos-sdk/types/rest"
)

type createEvent struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	EventId   uint         `json:"event_id"`
	StartTime uint         `json:"start_time"`
	Question  string       `json:"question"`
	Answers   []string     `json:"answers"`
	Winner    string       `json:"winner"`
	Loser     string       `json:"loser"`
	Owner     string       `json:"owner"`
}

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/privateevent/create", createPrivateEvent(cliCtx)).Methods("POST")
}

func createPrivateEvent(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createEvent

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
