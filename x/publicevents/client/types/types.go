package resTypes

import "github.com/cosmos/cosmos-sdk/types/rest"

type CreateEvent struct {
	BaseReq         rest.BaseReq `json:"base_req"`
	EventId         int          `json:"event_id"`
	EndTime         uint         `json:"end_time"`
	Question        string       `json:"question"`
	Answers         []string     `json:"answers"`
	CurrencyType    string       `json:"currencyType"`
	ValidatorAmount uint64       `json:"validatorAmount"`
}
