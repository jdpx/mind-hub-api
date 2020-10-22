package dynamo

import "context"

var (
	db map[string]interface{}
)

// Client ...
type Client struct {
}

// Config ...
type Config struct {
}

// NewClient ...
func NewClient(c Config) *Client {
	return &Client{}
}

// Put ...
func (c Client) Put(ctx context.Context, i interface{}) error {
	return nil
}
