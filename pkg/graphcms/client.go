//go:generate mockgen -source=client.go -destination=./mocks/client.go -package=graphcmsmocks

package graphcms

import (
	"context"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/machinebox/graphql"
)

// Requester defines the functionality to make requests to GraphCMS
type Requester interface {
	Run(ctx context.Context, req *Request, res interface{}) error
}

// Client makes request to GraphCMS
type Client struct {
	client map[string]Requester
}

// Request ...
type Request = graphql.Request

// Option ...
type Option func(*Client)

// NewClient initialises a new Client
func NewClient(opts ...Option) *Client {
	c := &Client{
		client: map[string]Requester{},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithOrganisationClient sets a client for a specific Organisation
func WithOrganisationClient(orgID string, client Requester) func(r *Client) {
	return func(r *Client) {
		r.client[orgID] = client
	}
}

// Run makes a request to GraphCMS
func (c Client) Run(ctx context.Context, req *Request, res interface{}) error {
	log := logging.NewFromResolver(ctx)

	oID, err := request.GetOrganisationID(ctx)
	if err != nil {
		log.WithError(err).
			Error("no organisation ID in context")

		return fmt.Errorf("no organisation ID in context %w", err)
	}

	client, ok := c.client[oID]
	if !ok {
		log.Error(fmt.Errorf("no client registered for organisation %s", oID))

		return fmt.Errorf("no client registered for organisation %s", oID)
	}

	err = client.Run(ctx, req, res)
	if err != nil {
		log.WithError(err).
			Error("Error occurred making request to GraphCMS")

		return fmt.Errorf("error occurred making request to GraphCMS: %v", err)
	}

	return nil
}
