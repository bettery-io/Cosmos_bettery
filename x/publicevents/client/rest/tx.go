package rest

// The packages below are commented out at first to prevent an error if this file isn't initially saved.
import (
	// "bytes"
	"net/http"

	"github.com/gorilla/mux"

	restTypes "github.com/VoroshilovMax/Bettery/x/publicevents/client/types"
	"github.com/VoroshilovMax/Bettery/x/publicevents/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/publicevent/create", createPublicEvent(cliCtx)).Methods("POST")
}

func createPublicEvent(cliCtx context.CLIContext) http.HandlerFunc {
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

		// eventAlreadyExist := helpers.CheckIfEventExist(req.EventId, w, r, cliCtx)
		// if eventAlreadyExist {
		// 	rest.WriteErrorResponse(w, http.StatusBadRequest, "Event with this id already exist")
		// 	return
		// }

		msg := types.NewMsgPublicCreateEvent(
			req.EventId,
			req.EndTime,
			owner,
			req.CurrencyType,
			req.ValidatorAmount,
			req.Question,
			req.Answers,
		)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
