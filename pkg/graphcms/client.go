//go:generate mockgen -source=client.go -destination=./mocks/client.go -package=graphcmsmocks

package graphcms

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/machinebox/graphql"
)

// Requester defines the functionality to make requests to GraphCMS
type Requester interface {
	Run(ctx context.Context, req *Request, resp interface{}) error
}

// Client makes request to GraphCMS
type Client struct {
	client Requester
}

// Request ...
type Request = graphql.Request

// DefaultHTTPClient ...
func DefaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

// NewClient initalises a new Client
func NewClient(client Requester) *Client {
	return &Client{
		client: client,
	}
}

// Run makes a request to GraphCMS
func (c Client) Run(ctx context.Context, req *Request, resp interface{}) error {
	log := logging.NewFromResolver(ctx)

	log.Info("Making request to GraphCMS")

	err := c.client.Run(ctx, req, resp)
	if err != nil {
		log.WithError(err).
			Error("Error occurred making request to GraphCMS")

		return fmt.Errorf("error occurred making request to GraphCMS: %v", err)
	}

	log.Info("Completed request to GraphCMS")

	return nil
}
