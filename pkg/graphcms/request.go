package graphcms

import "github.com/machinebox/graphql"

// NewRequest ...
func NewRequest(query string) *Request {
	return graphql.NewRequest(query)
}

// NewQueryRequest ...
func NewQueryRequest(query string, variables map[string]interface{}) *Request {
	req := graphql.NewRequest(query)

	for k, v := range variables {
		req.Var(k, v)
	}

	return req
}
