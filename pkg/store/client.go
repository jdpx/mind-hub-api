package store

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
func NewClient() *Client {
	db = map[string]interface{}{}

	return &Client{}
}

// Get ...
func (c Client) Get(ctx context.Context, key string) interface{} {
	return db[key]
}

// Put ...
func (c Client) Put(ctx context.Context, key string, i interface{}) error {
	db[key] = i

	return nil
}
