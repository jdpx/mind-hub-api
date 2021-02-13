package client

import "encoding/json"

// RawJSONError is a json formatted error from a GraphQL server.
type RawJSONError struct {
	json.RawMessage
}

func (r RawJSONError) Error() string {
	return string(r.RawMessage)
}
