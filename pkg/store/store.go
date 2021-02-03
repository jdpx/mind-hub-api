package store

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBer interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
}

type Storer interface {
	Get(ctx context.Context, tableName string, pk, sk string, i interface{}) error
	Query(ctx context.Context, tableName string, ex expression.Expression, i interface{}) error
	Put(ctx context.Context, tableName string, body interface{}) error
	Update(ctx context.Context, tableName string, pk, sk string, ex expression.Expression, i interface{}) error
}

type StoreOption func(*Store)

// Store ...
type Store struct {
	db DynamoDBer
}

func NewStore(opts ...StoreOption) *Store {
	s := &Store{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// WithDynamoDB ...
func WithDynamoDB(c DynamoDBer) func(*Store) {
	return func(r *Store) {
		r.db = c
	}
}
