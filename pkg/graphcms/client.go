//go:generate mockgen -source=client.go -destination=./mocks/client.go -package=graphcmsmocks

package graphcms

import (
	"context"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/machinebox/graphql"
)

// Requester defines the functionality to make requests to GraphCMS
type Requester interface {
	Run(ctx context.Context, req *Request, res interface{}) error
}

// Client makes request to GraphCMS
type Client struct {
	client Requester
}

// Request ...
type Request = graphql.Request

// NewClient initialises a new Client
func NewClient(client Requester) *Client {
	return &Client{
		client: client,
	}
}

// Run makes a request to GraphCMS
func (c Client) Run(ctx context.Context, req *Request, res interface{}) error {
	log := logging.NewFromResolver(ctx)

	err := c.client.Run(ctx, req, res)
	if err != nil {
		log.WithError(err).
			Error("Error occurred making request to GraphCMS")

		return fmt.Errorf("error occurred making request to GraphCMS: %v", err)
	}

	return nil
}
