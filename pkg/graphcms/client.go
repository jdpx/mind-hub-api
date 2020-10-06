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

type QueryRequest struct {
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
	Query         string                 `json:"query"`
}

// NewRequest ...
func NewRequest(query string) *Request {
	return graphql.NewRequest(query)
}

// NewQueryRequest ...
func NewQueryRequest(query string, name string, variables map[string]interface{}) (*Request, error) {
	fmt.Println("1111", variables)

	// req := QueryRequest{
	// 	OperationName: name,
	// 	Query:         query,
	// 	Variables:     variables,
	// }

	// b, err := json.Marshal(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, fmt.Errorf("error converting struct to json %v", err)
	// }

	req := graphql.NewRequest(query)

	req.Var("id", variables["id"])
	return req, nil
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
