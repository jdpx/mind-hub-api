package graphcms

import "github.com/machinebox/graphql"

// NewRequest creates a request containing the query to be sent to GraphCMS
func NewRequest(query string) *Request {
	return graphql.NewRequest(query)
}
