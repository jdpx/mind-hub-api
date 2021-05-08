//go:generate mockgen -source=store.go -destination=./mocks/store.go -package=storemocks
package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/sirupsen/logrus"
)

type Storer interface {
	Get(ctx context.Context, tableName string, pk, sk string, i interface{}) error
	BatchGet(ctx context.Context, tableName string, pk string, sk []string, i interface{}) error
	Query(ctx context.Context, tableName string, ex expression.Expression, i interface{}) error
	Put(ctx context.Context, tableName string, body interface{}) error
	Update(ctx context.Context, tableName string, pk, sk string, ex expression.Expression, i interface{}) error
}

type Option func(*Store)

// Store ...
type Store struct {
	db DynamoDBer
}

func NewStore(opts ...Option) *Store {
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

// Get ...
func (c Store) Get(ctx context.Context, tableName string, pk, sk string, i interface{}) error {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.PKKey: pk,
		logging.SKKey: sk,
	})

	input := dynamodb.GetItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: pk,
			},
			"SK": &types.AttributeValueMemberS{
				Value: sk,
			},
		},
	}

	results, err := c.db.GetItem(ctx, &input)
	if err != nil {
		if apiErr := new(types.ProvisionedThroughputExceededException); errors.As(err, &apiErr) {
			log.Error("error dynamodb throughput exceeded", err)
		} else if apiErr := new(types.InternalServerError); errors.As(err, &apiErr) {
			log.Error("internal server error from dynamodb", err)
		} else {
			log.Error("error getting item from dynamodb", err)
		}

		return fmt.Errorf("error returned from dynamodb %w", err)
	}

	if results == nil {
		log.Error("nil results returned")

		return fmt.Errorf("nil results returned")
	}

	if results.Item == nil {
		log.Info(fmt.Sprintf("No %s records found", tableName))

		return ErrNotFound
	}

	err = attributevalue.UnmarshalMap(results.Item, &i)
	if err != nil {
		log.Error("error unmarshalling dynamodb data", err)

		return fmt.Errorf("error unmarshalling dynamodb data, %w", err)
	}

	return nil
}

// Get ...
func (c Store) BatchGet(ctx context.Context, tableName string, pk string, sks []string, i interface{}) error {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.PKKey: pk,
	})

	var keys []map[string]types.AttributeValue

	for _, sk := range sks {
		keys = append(keys, map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: pk,
			},
			"SK": &types.AttributeValueMemberS{
				Value: sk,
			},
		})
	}

	input := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			tableName: {
				Keys: keys,
			},
		},
	}

	results, err := c.db.BatchGetItem(ctx, &input)
	if err != nil {
		if apiErr := new(types.ProvisionedThroughputExceededException); errors.As(err, &apiErr) {
			log.Error("error dynamodb throughput exceeded", err)
		} else if apiErr := new(types.InternalServerError); errors.As(err, &apiErr) {
			log.Error("internal server error from dynamodb", err)
		} else {
			log.Error("error getting item from dynamodb", err)
		}

		return fmt.Errorf("error returned from dynamodb %w", err)
	}

	if results == nil {
		log.Error("nil results returned")

		return fmt.Errorf("nil results returned")
	}

	items := results.Responses[tableName]

	err = attributevalue.UnmarshalListOfMaps(items, &i)
	if err != nil {
		log.Error("error unmarshalling dynamodb data", err)

		return fmt.Errorf("error unmarshalling dynamodb data, %w", err)
	}

	return nil
}

// Query ...
func (c Store) Query(ctx context.Context, tableName string, ex expression.Expression, i interface{}) error {
	log := logging.NewFromResolver(ctx)

	queryInput := dynamodb.QueryInput{
		KeyConditionExpression:    ex.KeyCondition(),
		ProjectionExpression:      ex.Projection(),
		ExpressionAttributeNames:  ex.Names(),
		ExpressionAttributeValues: ex.Values(),
		TableName:                 &tableName,
	}

	result, err := c.db.Query(ctx, &queryInput)
	if err != nil {
		if apiErr := new(types.ProvisionedThroughputExceededException); errors.As(err, &apiErr) {
			log.Error("error dynamodb throughput exceeded", err)
		} else if apiErr := new(types.InternalServerError); errors.As(err, &apiErr) {
			log.Error("internal server error from dynamodb", err)
		} else {
			log.Error("error getting item from dynamodb", err)
		}

		return fmt.Errorf("error returned from dynamodb %w", err)
	}

	if result == nil {
		log.Error("nil results returned")

		return fmt.Errorf("nil results returned")
	}

	err = attributevalue.UnmarshalListOfMaps(result.Items, &i)
	if err != nil {
		log.Error("error unmarshalling dynamodb data", err)

		return fmt.Errorf("error unmarshalling dynamodb data, %w", err)
	}

	return nil
}

// Put ...
func (c Store) Put(ctx context.Context, tableName string, body interface{}) error {
	log := logging.NewFromResolver(ctx)

	av, err := attributevalue.MarshalMap(body)
	if err != nil {
		return fmt.Errorf("failed to marshal Record, %w", err)
	}

	_, err = c.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      av,
	})

	if err != nil {
		if apiErr := new(types.ProvisionedThroughputExceededException); errors.As(err, &apiErr) {
			log.Error("error dynamodb throughput exceeded", err)
		} else if apiErr := new(types.InternalServerError); errors.As(err, &apiErr) {
			log.Error("internal server error from dynamodb", err)
		} else {
			log.Error("error getting item from dynamodb", err)
		}

		return fmt.Errorf("error returned from dynamodb %w", err)
	}

	return nil
}

func (c Store) Update(ctx context.Context, tableName string, pk, sk string, ex expression.Expression, i interface{}) error {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.PKKey: pk,
		logging.SKKey: sk,
	})

	input := dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: pk,
			},
			"SK": &types.AttributeValueMemberS{
				Value: sk,
			},
		},
		TableName:                 &tableName,
		ExpressionAttributeNames:  ex.Names(),
		ExpressionAttributeValues: ex.Values(),
		ReturnValues:              types.ReturnValueUpdatedNew,
		UpdateExpression:          ex.Update(),
	}

	result, err := c.db.UpdateItem(ctx, &input)
	if err != nil {
		if apiErr := new(types.ProvisionedThroughputExceededException); errors.As(err, &apiErr) {
			log.Error("error dynamodb throughput exceeded", err)
		} else if apiErr := new(types.InternalServerError); errors.As(err, &apiErr) {
			log.Error("internal server error from dynamodb", err)
		} else {
			log.Error("error getting item from dynamodb", err)
		}

		return fmt.Errorf("error returned from dynamodb %w", err)
	}

	if result == nil {
		log.Error("nil results returned")

		return fmt.Errorf("nil results returned")
	}

	err = attributevalue.UnmarshalMap(result.Attributes, &i)
	if err != nil {
		log.Error("error unmarshalling dynamodb data", err)

		return fmt.Errorf("error unmarshalling dynamodb data, %w", err)
	}

	return nil
}
