package rest

import (
	// "bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	// "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	// "github.com/VoroshilovMax/Bettery/x/privateevents/types"
)

type createEvent struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	EventId   int          `json:"event_id"`
	StartTime int          `json:"start_time"`
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

		fmt.Println(req.Owner)
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
