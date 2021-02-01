//go:generate mockgen -source=client.go -destination=./mocks/client.go -package=storemocks

package store

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/sirupsen/logrus"
)

var (
	// ErrNotFound ...
	ErrNotFound = fmt.Errorf("record not found")
)

const (
	dbRegion         = "eu-west-1"
	localDynamoDBURL = "http://localhost:8000"
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

// NewClient ...
func NewClient() *dynamodb.Client {
	config, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(dbRegion))
	if err != nil {
		log.Fatal(err)
	}

	return dynamodb.NewFromConfig(config)
}

func NewLocalClient() *dynamodb.Client {
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID && region == dbRegion {
			return aws.Endpoint{
				PartitionID:       "aws",
				URL:               localDynamoDBURL,
				SigningRegion:     dbRegion,
				HostnameImmutable: true,
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	config, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(dbRegion),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		log.Fatal(err)
	}

	dbSvc := dynamodb.NewFromConfig(config)

	setupDbV2Tables(dbSvc)

	return dbSvc
}

// Get ...
func (c Store) Get(ctx context.Context, tableName string, pk, sk string, i interface{}) error {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.PKKey: pk,
		logging.SKKey: sk,
	})

	input := dynamodb.GetItemInput{
		TableName: aws.String(tableName),
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

// Query ...
func (c Store) Query(ctx context.Context, tableName string, ex expression.Expression, i interface{}) error {
	log := logging.NewFromResolver(ctx)

	queryInput := dynamodb.QueryInput{
		KeyConditionExpression:    ex.KeyCondition(),
		ProjectionExpression:      ex.Projection(),
		ExpressionAttributeNames:  ex.Names(),
		ExpressionAttributeValues: ex.Values(),
		TableName:                 aws.String(tableName),
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
		TableName: aws.String(tableName),
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
		TableName:                 aws.String(tableName),
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

func setupDbV2Tables(dbSvc *dynamodb.Client) {
	_, err := dbSvc.CreateTable(context.Background(), &dynamodb.CreateTableInput{
		TableName: aws.String("user"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("PK"),
				AttributeType: "S",
			},
			{
				AttributeName: aws.String("SK"),
				AttributeType: "S",
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("PK"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("SK"),
				KeyType:       types.KeyTypeRange,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	})
	if err != nil {
		log.Println(err)
	}
}
