package rest

import (
	// "bytes"
	"fmt"
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
	r.HandleFunc(
		"/private_event/create", createPrivateEvent(cliCtx),
	).Methods("POST")
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

		fmt.Println(req.Owner)

		onwer, err := sdk.AccAddressFromBech32(req.Owner)
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
		)
		err = msg.ValidateBasic()
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		fmt.Println("WORK")

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})

	}
}

/*
// Action TX body
type <Action>Req struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	// TODO: Define more types if needed
}

func <Action>RequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req <Action>Req
		vars := mux.Vars(r)

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// TODO: Define the module tx logic for this action

		utils.WriteGenerateStdTxResponse(w, cliCtx, BaseReq, []sdk.Msg{msg})
	}
}
*/
