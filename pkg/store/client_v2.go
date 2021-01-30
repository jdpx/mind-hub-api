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
)

type StorerV2 interface {
	Get(ctx context.Context, tableName string, pk, sk string, i interface{}) error
	Query(ctx context.Context, tableName string, ex expression.Expression, i interface{}) error
	Put(ctx context.Context, tableName string, body interface{}) error
	Update(ctx context.Context, tableName string, pk, sk string, ex expression.Expression, i interface{}) error
}

// ClientV2 ...
type ClientV2 struct {
	db *dynamodb.Client
}

type Key struct {
	PK string
	SK string
}

// NewClientV2 ...
func NewClientV2(c Config) (*ClientV2, error) {
	config, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(dbRegion))
	if err != nil {
		log.Fatal(err)
	}

	if c.Env == "local" {
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

		config.EndpointResolver = customResolver
	}

	dbSvc := dynamodb.NewFromConfig(config)

	if c.Env == "local" {
		setupDbV2Tables(dbSvc)
	}

	return &ClientV2{
		db: dbSvc,
	}, nil
}

// Get ...
func (c ClientV2) Get(ctx context.Context, tableName string, pk, sk string, i interface{}) error {
	log := logging.NewFromResolver(ctx)

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
		log.Error("error getting item from store", err)

		return err
	}

	if results.Item == nil {
		log.Info(fmt.Sprintf("No %s records found", tableName))

		return ErrNotFound
	}

	err = attributevalue.UnmarshalMap(results.Item, &i)
	if err != nil {
		return fmt.Errorf("failed to unmarshal Items, %w", err)
	}

	return nil
}

// Query ...
func (c ClientV2) Query(ctx context.Context, tableName string, ex expression.Expression, i interface{}) error {
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
		log.Error("error querying store", err)

		return err
	}

	err = attributevalue.UnmarshalListOfMaps(result.Items, &i)
	if err != nil {
		return fmt.Errorf("failed to unmarshal Items, %w", err)
	}

	return nil
}

// Put ...
func (c ClientV2) Put(ctx context.Context, tableName string, body interface{}) error {
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
		log.Error(fmt.Sprintf("error putting item to %s store", tableName), err)

		return fmt.Errorf("error putting item to %s store %w", tableName, err)
	}

	return nil
}

func (c ClientV2) Update(ctx context.Context, tableName string, pk, sk string, ex expression.Expression, i interface{}) error {
	input := &dynamodb.UpdateItemInput{
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

	result, err := c.db.UpdateItem(ctx, input)
	if err != nil {
		if apiErr := new(types.ProvisionedThroughputExceededException); errors.As(err, &apiErr) {
			fmt.Println("throughput exceeded")
		} else if apiErr := new(types.ResourceNotFoundException); errors.As(err, &apiErr) {
			fmt.Println("resource not found")
		} else if apiErr := new(types.InternalServerError); errors.As(err, &apiErr) {
			fmt.Println("internal server error")
		} else {
			fmt.Println(err)
		}
		return fmt.Errorf("error updating item to %s store %w", tableName, err)
	}

	err = attributevalue.UnmarshalMap(result.Attributes, &i)
	if err != nil {
		return fmt.Errorf("failed to unmarshal Items, %w", err)
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
