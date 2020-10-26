//go:generate mockgen -source=client.go -destination=./mocks/client.go -package=graphcmsmocks

package graphcms

import (
	"context"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/logging"
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

// Run ...
func (c Client) Run(ctx context.Context, req *Request, resp interface{}) error {
	log := logging.NewFromResolver(ctx)

	log.Info("Making request to GraphCMS")

	err := c.client.Run(ctx, req, resp)
	if err != nil {
		log.WithError(err).
			Error("Error occurred making request to GraphCMS")

		return fmt.Errorf("error occurred making request to GraphCMS: %v", err)
	}

	return nil
}
