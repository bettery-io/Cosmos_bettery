package types

import "strings"

const (
	QueryListPrivateEvent  = "list"
	QueryGetSinglePrivateEvent   = "get"
	QueryGetPrivateEventStatus = "status"
)

// QueryResQuiz Queries Result Payload for a names query
type QueryResPrivateEvent []string

// implement fmt.Stringer
func (n QueryResPrivateEvent) String() string {
	return strings.Join(n[:], "\n")
}

