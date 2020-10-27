package types

import "strings"

// Query endpoints supported by the coinmaker querier
const (
// TODO: Describe query parameters, update <action> with your query
// Query<Action>    = "<action>"
)

// QueryResList Queries Result Payload for a query
type QueryResList []string

// implement fmt.Stringer
func (n QueryResList) String() string {
	return strings.Join(n[:], "\n")
}
