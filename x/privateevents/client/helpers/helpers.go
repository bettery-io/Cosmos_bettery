package helpers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/VoroshilovMax/Bettery/x/privateevents/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func GetAnswerNumber(EventId int, Answer string, w http.ResponseWriter, r *http.Request, cliCtx context.CLIContext) int {
	eventId := strconv.FormatUint(uint64(EventId), 10)

	cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
	if !ok {
		rest.WriteErrorResponse(w, http.StatusNotFound, "ParseQueryHeightOrReturnBadRequest")
	}

	res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryGetSinglePrivateEvent, eventId), nil)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusNotFound, "could not resolve quiz by id: "+eventId+" ,"+err.Error())
	}

	var event types.CreateEvent
	cliCtx.Codec.MustUnmarshalJSON(res, &event)
	return IndexOf(event.Answers, Answer)
}

func CheckIfEventExist(EventId int, w http.ResponseWriter, r *http.Request, cliCtx context.CLIContext) bool {
	eventId := strconv.FormatUint(uint64(EventId), 10)

	cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
	if !ok {
		rest.WriteErrorResponse(w, http.StatusNotFound, "ParseQueryHeightOrReturnBadRequest")
	}

	_, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryGetSinglePrivateEvent, eventId), nil)
	if err != nil {
		return false
	} else {
		return true
	}
}

func IndexOf(haystack []string, needle string) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func CurrentEpochTime() int64 {
	return time.Now().Unix()
}
