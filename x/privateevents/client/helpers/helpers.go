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

func ValidateExpert(EventId int, w http.ResponseWriter, r *http.Request, cliCtx context.CLIContext) bool {
	res, _, err := getEvent(EventId, w, r, cliCtx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusNotFound, "could not resolve event by id: "+strconv.FormatUint(uint64(EventId), 10)+" ,"+err.Error())
	}
	var event types.EventInfo
	cliCtx.Codec.MustUnmarshalJSON(res, &event)
	if event.FinalAnswer == "undefined" {
		// check if time is valid
		time := time.Now().Unix()
		if uint(time) > event.EndTime {
			return false
		} else {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Time for validator is not started")
			return true
		}
	} else {
		rest.WriteErrorResponse(w, http.StatusBadRequest, "Event is finish. Answer is alredy provided")
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

func ValidateParticipant(partWallet string, EventId int, w http.ResponseWriter, r *http.Request, cliCtx context.CLIContext) bool {
	res, _, err := getEvent(EventId, w, r, cliCtx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusNotFound, "could not resolve event by id: "+strconv.FormatUint(uint64(EventId), 10)+" ,"+err.Error())
	}
	var event types.EventInfo
	cliCtx.Codec.MustUnmarshalJSON(res, &event)

	// check if participant already participate
	for _, v := range event.Participants {
		if v.Participant.String() == partWallet {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Participant already participate")
			return true
		}
	}

	// check if time is valid
	time := time.Now().Unix()
	if uint(time) > event.EndTime {
		rest.WriteErrorResponse(w, http.StatusBadRequest, "Time for paticipate is finish")
		return true
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
