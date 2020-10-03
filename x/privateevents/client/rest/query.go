package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/VoroshilovMax/Bettery/x/privateevents/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/privateevents/{id}", getEventById(cliCtx)).Methods("GET")
}

// TODO: Make better structure for single event with all participant and validators
func getEventById(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		eventId := mux.Vars(r)["id"]
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryGetSinglePrivateEvent, eventId), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, "could not resolve quiz by id: "+eventId+" ,"+err.Error())
			fmt.Printf("could not resolve quiz %s \n%s\n", eventId, err.Error())
			return
		}

		var out types.CreateEvent
		cliCtx.Codec.MustUnmarshalJSON(res, &out)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
