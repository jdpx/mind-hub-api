//go:generate mockgen -source=client.go -destination=./mocks/client.go -package=storemocks

package store

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jdpx/mind-hub-api/pkg/logging"
)

var (
	// ErrNotFound ...
	ErrNotFound = fmt.Errorf("record not found")
)

// Storer ...
type Storer interface {
	Get(ctx context.Context, tableName string, searchBody interface{}, i interface{}) error
	Put(ctx context.Context, tableName string, body interface{}) error
	Update(ctx context.Context, tableName string, keys map[string]string, expression string, body interface{}, i interface{}) error
}

// Client ...
type Client struct {
	db *dynamodb.DynamoDB
}

// Config ...
type Config struct {
}

// NewClient ...
func NewClient(config Config) (*Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("eu-west-1"),
		Endpoint: aws.String("http://localhost:8000"),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dbSvc := dynamodb.New(sess)
	_, err = dbSvc.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String("course_progress"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("courseID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("userID"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("courseID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("userID"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	})
	if err != nil {
		log.Println(err)
	}

	_, err = dbSvc.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String("course_note"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("courseID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("userID"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("courseID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("userID"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	})
	if err != nil {
		log.Println(err)
	}

	result, err := dbSvc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Tables:")
	for _, table := range result.TableNames {
		log.Println(*table)
	}

	return &Client{
		db: dbSvc,
	}, nil
}

// Get ...
func (c Client) Get(ctx context.Context, tableName string, searchBody interface{}, i interface{}) error {
	log := logging.NewFromResolver(ctx)

	k, err := dynamodbattribute.MarshalMap(searchBody)
	if err != nil {
		return err
	}

	input := dynamodb.GetItemInput{
		Key:       k,
		TableName: aws.String(tableName),
	}

	result, err := c.db.GetItem(&input)
	if err != nil {
		log.Error("error getting item from store", err)

		return err
	}

	if len(result.Item) == 0 {
		log.Info(fmt.Sprintf("No %s records found", tableName))

		return ErrNotFound
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &i)
	if err != nil {
		log.Error("error unmarshalling from store", err)

		return err
	}

	return nil
}

// Put ...
func (c Client) Put(ctx context.Context, tableName string, body interface{}) error {
	log := logging.NewFromResolver(ctx)

	info, err := dynamodbattribute.MarshalMap(body)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal the movie, %v", err))
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String(tableName),
	}

	_, err = c.db.PutItemWithContext(ctx, input)
	if err != nil {
		log.Error("error putting item from store", err)
		return err
	}

	return nil
}

// Update ...
func (c Client) Update(ctx context.Context, tableName string, keys map[string]string, expression string, body interface{}, i interface{}) error {
	log := logging.NewFromResolver(ctx)

	k, err := dynamodbattribute.MarshalMap(keys)
	if err != nil {
		return err
	}

	b, err := dynamodbattribute.MarshalMap(body)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal the movie, %v", err))
	}

	input := &dynamodb.UpdateItemInput{
		Key:                       k,
		ExpressionAttributeValues: b,
		TableName:                 aws.String(tableName),
		ReturnValues:              aws.String("UPDATED_NEW"),
		UpdateExpression:          aws.String(expression),
	}

	result, err := c.db.UpdateItemWithContext(ctx, input)
	if err != nil {
		log.Error("error updating item from store", err)
		return err
	}

	err = dynamodbattribute.UnmarshalMap(result.Attributes, &i)
	if err != nil {
		log.Error("error unmarshalling from store", err)

		return err
	}

	return nil
}
