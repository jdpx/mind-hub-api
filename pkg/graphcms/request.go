package graphcms

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/machinebox/graphql"
)

// NewRequest creates a request containing the query to be sent to GraphCMS
func NewRequest(ctx context.Context, query string) *Request {
	r := graphql.NewRequest(query)

	cID, err := request.ContextCorrelationID(ctx)
	if err != nil {
		return r
	}

	r.Header.Add(request.CorrelationIDHeader, cID)

	return r
}
