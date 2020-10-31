//go:generate mockgen -source=client.go -destination=./mocks/client.go -package=storemocks

package store

import "context"

var (
	db map[string]map[string]interface{}
)

// Storer ...
type Storer interface {
	Get(ctx context.Context, tableName string, key string) (interface{}, error)
	Put(ctx context.Context, tableName string, key string, i interface{}) error
}

// Client ...
type Client struct {
}

// Config ...
type Config struct {
}

// NewClient ...
func NewClient() *Client {
	db = map[string]map[string]interface{}{}

	return &Client{}
}

// Get ...
func (c Client) Get(ctx context.Context, tableName string, key string) (interface{}, error) {
	_, ok := db[tableName]
	if !ok {
		db[tableName] = map[string]interface{}{}
	}

	return db[tableName][key], nil
}

// Put ...
func (c Client) Put(ctx context.Context, tableName string, key string, i interface{}) error {
	_, ok := db[tableName]
	if !ok {
		db[tableName] = map[string]interface{}{}
	}

	db[tableName][key] = i

	return nil
}
