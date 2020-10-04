package resTypes

import "github.com/cosmos/cosmos-sdk/types/rest"

type CreateEvent struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	EventId  int          `json:"event_id"`
	EndTime  uint         `json:"end_time"`
	Question string       `json:"question"`
	Answers  []string     `json:"answers"`
	Winner   string       `json:"winner"`
	Loser    string       `json:"loser"`
	Owner    string       `json:"owner"`
}

type Participate struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Participant string       `json:"participant"`
	Answer      string       `json:"answer"`
	EventId     int          `json:"event_id"`
}

type Validate struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Expert  string       `json:"expert"`
	Answer  string       `json:"answer"`
	EventId int          `json:"event_id"`
}
