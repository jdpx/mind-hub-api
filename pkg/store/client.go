//go:generate mockgen -source=client.go -destination=./mocks/client.go -package=storemocks

package store

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
	BatchGetItem(ctx context.Context, params *dynamodb.BatchGetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error)
	Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
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
