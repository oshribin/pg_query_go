// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node A_Star) MarshalJSON() ([]byte, error) {
	type A_StarMarshalAlias A_Star
	return json.Marshal(map[string]interface{}{
		"A_STAR": (*A_StarMarshalAlias)(&node),
	})
}

func (node *A_Star) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node A_Star) Deparse() string {
	panic("Not Implemented")
}