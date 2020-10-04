package graphcms

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

// Client ...
type Client struct {
	client *graphql.Client
}

// Request ...
type Request = graphql.Request

// NewRequest ...
func NewRequest(query string) *Request {
	return graphql.NewRequest(query)
}

// NewClient ...
func NewClient(url string) *Client {
	return &Client{
		client: graphql.NewClient(url),
	}
}

// Run ...
func (c Client) Run(ctx context.Context, req *Request, resp interface{}) error {
	err := c.client.Run(ctx, req, resp)
	if err != nil {
		return fmt.Errorf("error occurred making request to GraphCMS %v", err)
	}

	return nil
}
