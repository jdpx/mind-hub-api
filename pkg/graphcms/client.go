//go:generate mockgen -source=client.go -destination=./mocks/client.go -package=graphcmsmocks

package graphcms

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

// CMSRequster ...
type CMSRequster interface {
	Run(ctx context.Context, req *graphql.Request, resp interface{}) error
}

// Client ...
type Client struct {
	client CMSRequster
}

// Request ...
type Request = graphql.Request

// NewClient ...
func NewClient(client CMSRequster) *Client {
	return &Client{
		client: client,
	}
}

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

// Run ...
func (c Client) Run(ctx context.Context, req *Request, resp interface{}) error {
	err := c.client.Run(ctx, req, resp)
	if err != nil {
		return fmt.Errorf("error occurred making request to GraphCMS: %v", err)
	}

	return nil
}
