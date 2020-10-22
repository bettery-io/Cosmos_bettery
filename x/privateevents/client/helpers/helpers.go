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
	res, _, err := getEvent(EventId, w, r, cliCtx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusNotFound, "could not resolve event by id: "+strconv.FormatUint(uint64(EventId), 10)+" ,"+err.Error())
	}

	var event types.CreateEvent
	cliCtx.Codec.MustUnmarshalJSON(res, &event)
	return IndexOf(event.Answers, Answer)
}

func CheckIfEventExist(EventId int, w http.ResponseWriter, r *http.Request, cliCtx context.CLIContext) bool {
	_, _, err := getEvent(EventId, w, r, cliCtx)
	if err != nil {
		return false
	} else {
		return true
	}
}

func AnswerIsKnown(EventId int, w http.ResponseWriter, r *http.Request, cliCtx context.CLIContext) bool {
	res, _, err := getEvent(EventId, w, r, cliCtx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusNotFound, "could not resolve event by id: "+strconv.FormatUint(uint64(EventId), 10)+" ,"+err.Error())
	}
	var event types.EventInfo
	cliCtx.Codec.MustUnmarshalJSON(res, &event)
	if event.FinalAnswer == "undefined" {
		return false
	} else {
		return true
	}
}

func getEvent(EventId int, w http.ResponseWriter, r *http.Request, cliCtx context.CLIContext) ([]byte, int64, error) {
	eventId := strconv.FormatUint(uint64(EventId), 10)

	cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
	if !ok {
		rest.WriteErrorResponse(w, http.StatusNotFound, "ParseQueryHeightOrReturnBadRequest")
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryGetSinglePrivateEvent, eventId), nil)
}

func CheckIfParticipantParticipate(partWallet string, EventId int, w http.ResponseWriter, r *http.Request, cliCtx context.CLIContext) bool {
	res, _, err := getEvent(EventId, w, r, cliCtx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusNotFound, "could not resolve event by id: "+strconv.FormatUint(uint64(EventId), 10)+" ,"+err.Error())
	}
	var event types.EventInfo
	cliCtx.Codec.MustUnmarshalJSON(res, &event)
	for _, v := range event.Participants {
		if v.Participant.String() == partWallet {
			return true
		}
	}

	return false
}

func IndexOf(answers []string, answer string) int {
	for i, v := range answers {
		if v == answer {
			return i
		}
	}
	return -1
}

func CurrentEpochTime() int64 {
	return time.Now().Unix()
}
