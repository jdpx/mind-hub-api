package store_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
	storemocks "github.com/jdpx/mind-hub-api/pkg/store/mocks"
	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Foo string `json:"foo"`
}

func TestStoreGet(t *testing.T) {
	tableName := fake.CharactersN(5)
	pk := fake.CharactersN(10)
	sk := fake.CharactersN(10)
	value := fake.CharactersN(10)

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockDynamoDBer)

		expectedStruct testStruct
		expectedErr    error
	}{
		{
			desc: "given record is found in store, output struct set",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
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

				result := dynamodb.GetItemOutput{
					Item: map[string]types.AttributeValue{
						"foo": &types.AttributeValueMemberS{
							Value: value,
						},
					},
				}

				client.EXPECT().GetItem(gomock.Any(), &input).Return(&result, nil)
			},

			expectedStruct: testStruct{
				Foo: value,
			},
		},
		{
			desc: "given nil result is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
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

				client.EXPECT().GetItem(gomock.Any(), &input).Return(nil, nil)
			},

			expectedErr: fmt.Errorf("nil results returned"),
		},
		{
			desc: "given nil result item is returned, Not Found error returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
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

				result := dynamodb.GetItemOutput{}

				client.EXPECT().GetItem(gomock.Any(), &input).Return(&result, nil)
			},

			expectedErr: store.ErrNotFound,
		},
		{
			desc: "given a ThroughputExeeded error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
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

				client.EXPECT().GetItem(gomock.Any(), &input).Return(nil, &types.ProvisionedThroughputExceededException{
					Message: aws.String("something went wrong"),
				})
			},

			expectedErr: fmt.Errorf("error returned from dynamodb ProvisionedThroughputExceededException: something went wrong"),
		},
		{
			desc: "given a InternalServerError error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
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

				client.EXPECT().GetItem(gomock.Any(), &input).Return(nil, &types.InternalServerError{
					Message: aws.String("something went wrong"),
				})
			},

			expectedErr: fmt.Errorf("error returned from dynamodb InternalServerError: something went wrong"),
		},
		{
			desc: "given an error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
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

				client.EXPECT().GetItem(gomock.Any(), &input).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error returned from dynamodb something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeMock := storemocks.NewMockDynamoDBer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(storeMock)
			}

			store := store.NewStore(store.WithDynamoDB(storeMock))
			ctx := context.Background()
			m := testStruct{}

			err := store.Get(ctx, tableName, pk, sk, &m)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedStruct, m)
			}
		})
	}
}

func TestStoreBatchGet(t *testing.T) {
	tableName := fake.CharactersN(5)
	pk := fake.CharactersN(10)
	skOne := fake.CharactersN(10)
	skTwo := fake.CharactersN(10)

	valueOne := fake.CharactersN(10)
	valueTwo := fake.CharactersN(10)

	dynInput := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			tableName: {
				Keys: []map[string]types.AttributeValue{
					{
						"PK": &types.AttributeValueMemberS{
							Value: pk,
						},
						"SK": &types.AttributeValueMemberS{
							Value: skOne,
						},
					},
					{
						"PK": &types.AttributeValueMemberS{
							Value: pk,
						},
						"SK": &types.AttributeValueMemberS{
							Value: skTwo,
						},
					},
				},
			},
		},
	}

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockDynamoDBer)

		expectedStruct []testStruct
		expectedErr    error
	}{
		{
			desc: "given record is found in store, output struct set",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {

				result := dynamodb.BatchGetItemOutput{
					Responses: map[string][]map[string]types.AttributeValue{
						tableName: {
							map[string]types.AttributeValue{
								"foo": &types.AttributeValueMemberS{
									Value: valueOne,
								},
							},
							map[string]types.AttributeValue{
								"foo": &types.AttributeValueMemberS{
									Value: valueTwo,
								},
							},
						},
					},
				}

				client.EXPECT().BatchGetItem(gomock.Any(), &dynInput).Return(&result, nil)
			},

			expectedStruct: []testStruct{
				{
					Foo: valueOne,
				},
				{
					Foo: valueTwo,
				},
			},
		},
		{
			desc: "given nil result is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				client.EXPECT().BatchGetItem(gomock.Any(), &dynInput).Return(nil, nil)
			},

			expectedErr: fmt.Errorf("nil results returned"),
		},
		{
			desc: "given a ThroughputExeeded error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				client.EXPECT().BatchGetItem(gomock.Any(), &dynInput).Return(nil, &types.ProvisionedThroughputExceededException{
					Message: aws.String("something went wrong"),
				})
			},

			expectedErr: fmt.Errorf("error returned from dynamodb ProvisionedThroughputExceededException: something went wrong"),
		},
		{
			desc: "given a InternalServerError error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				client.EXPECT().BatchGetItem(gomock.Any(), &dynInput).Return(nil, &types.InternalServerError{
					Message: aws.String("something went wrong"),
				})
			},

			expectedErr: fmt.Errorf("error returned from dynamodb InternalServerError: something went wrong"),
		},
		{
			desc: "given an error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				client.EXPECT().BatchGetItem(gomock.Any(), &dynInput).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error returned from dynamodb something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeMock := storemocks.NewMockDynamoDBer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(storeMock)
			}

			store := store.NewStore(store.WithDynamoDB(storeMock))
			ctx := context.Background()
			m := []testStruct{}

			err := store.BatchGet(ctx, tableName, pk, []string{skOne, skTwo}, &m)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedStruct, m)
			}
		})
	}
}

func TestStoreQuery(t *testing.T) {
	tableName := fake.CharactersN(5)
	value := fake.CharactersN(10)

	ex, _ := expression.NewBuilder().
		WithCondition(expression.Name("foo").Equal(expression.Value(5))).
		WithFilter(expression.Name("bar").LessThan(expression.Value(6))).
		WithProjection(expression.NamesList(expression.Name("foo"), expression.Name("bar"), expression.Name("baz"))).
		WithKeyCondition(expression.Key("foo").Equal(expression.Value(5))).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockDynamoDBer)

		expectedStruct []testStruct
		expectedErr    error
	}{
		{
			desc: "given record is found in store, output struct set",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				input := dynamodb.QueryInput{
					KeyConditionExpression:    ex.KeyCondition(),
					ProjectionExpression:      ex.Projection(),
					ExpressionAttributeNames:  ex.Names(),
					ExpressionAttributeValues: ex.Values(),
					TableName:                 aws.String(tableName),
				}

				result := dynamodb.QueryOutput{
					Items: []map[string]types.AttributeValue{
						{
							"foo": &types.AttributeValueMemberS{
								Value: value,
							},
						},
					},
				}

				client.EXPECT().Query(gomock.Any(), &input).Return(&result, nil)
			},

			expectedStruct: []testStruct{
				{
					Foo: value,
				},
			},
		},
		{
			desc: "given nil result is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				input := dynamodb.QueryInput{
					KeyConditionExpression:    ex.KeyCondition(),
					ProjectionExpression:      ex.Projection(),
					ExpressionAttributeNames:  ex.Names(),
					ExpressionAttributeValues: ex.Values(),
					TableName:                 aws.String(tableName),
				}

				client.EXPECT().Query(gomock.Any(), &input).Return(nil, nil)
			},

			expectedErr: fmt.Errorf("nil results returned"),
		},
		{
			desc: "given nil result item is returned, Not Found error returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				input := dynamodb.QueryInput{
					KeyConditionExpression:    ex.KeyCondition(),
					ProjectionExpression:      ex.Projection(),
					ExpressionAttributeNames:  ex.Names(),
					ExpressionAttributeValues: ex.Values(),
					TableName:                 aws.String(tableName),
				}

				result := dynamodb.QueryOutput{}

				client.EXPECT().Query(gomock.Any(), &input).Return(&result, nil)
			},
			expectedStruct: []testStruct{},
		},
		{
			desc: "given a ThroughputExeeded error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				input := dynamodb.QueryInput{
					KeyConditionExpression:    ex.KeyCondition(),
					ProjectionExpression:      ex.Projection(),
					ExpressionAttributeNames:  ex.Names(),
					ExpressionAttributeValues: ex.Values(),
					TableName:                 aws.String(tableName),
				}

				client.EXPECT().Query(gomock.Any(), &input).Return(nil, &types.ProvisionedThroughputExceededException{
					Message: aws.String("something went wrong"),
				})
			},

			expectedErr: fmt.Errorf("error returned from dynamodb ProvisionedThroughputExceededException: something went wrong"),
		},
		{
			desc: "given a InternalServerError error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				input := dynamodb.QueryInput{
					KeyConditionExpression:    ex.KeyCondition(),
					ProjectionExpression:      ex.Projection(),
					ExpressionAttributeNames:  ex.Names(),
					ExpressionAttributeValues: ex.Values(),
					TableName:                 aws.String(tableName),
				}

				client.EXPECT().Query(gomock.Any(), &input).Return(nil, &types.InternalServerError{
					Message: aws.String("something went wrong"),
				})
			},

			expectedErr: fmt.Errorf("error returned from dynamodb InternalServerError: something went wrong"),
		},
		{
			desc: "given an error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				input := dynamodb.QueryInput{
					KeyConditionExpression:    ex.KeyCondition(),
					ProjectionExpression:      ex.Projection(),
					ExpressionAttributeNames:  ex.Names(),
					ExpressionAttributeValues: ex.Values(),
					TableName:                 aws.String(tableName),
				}

				client.EXPECT().Query(gomock.Any(), &input).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error returned from dynamodb something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeMock := storemocks.NewMockDynamoDBer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(storeMock)
			}

			store := store.NewStore(store.WithDynamoDB(storeMock))
			ctx := context.Background()
			m := []testStruct{}

			err := store.Query(ctx, tableName, ex, &m)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedStruct, m)
			}
		})
	}
}

func TestStorePut(t *testing.T) {
	tableName := fake.CharactersN(5)
	value := fake.CharactersN(10)

	entity := testStruct{
		Foo: value,
	}

	input := dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"Foo": &types.AttributeValueMemberS{
				Value: value,
			},
		},
	}

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockDynamoDBer)

		expectedStruct []testStruct
		expectedErr    error
	}{
		{
			desc: "given record is found in store, output struct set",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {

				client.EXPECT().PutItem(gomock.Any(), &input).Return(nil, nil)
			},
		},
		{
			desc: "given a ThroughputExeeded error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				client.EXPECT().PutItem(gomock.Any(), &input).Return(nil, &types.ProvisionedThroughputExceededException{
					Message: aws.String("something went wrong"),
				})
			},

			expectedErr: fmt.Errorf("error returned from dynamodb ProvisionedThroughputExceededException: something went wrong"),
		},
		{
			desc: "given a InternalServerError error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				client.EXPECT().PutItem(gomock.Any(), &input).Return(nil, &types.InternalServerError{
					Message: aws.String("something went wrong"),
				})
			},

			expectedErr: fmt.Errorf("error returned from dynamodb InternalServerError: something went wrong"),
		},
		{
			desc: "given an error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				client.EXPECT().PutItem(gomock.Any(), &input).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error returned from dynamodb something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeMock := storemocks.NewMockDynamoDBer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(storeMock)
			}

			store := store.NewStore(store.WithDynamoDB(storeMock))
			ctx := context.Background()

			err := store.Put(ctx, tableName, entity)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestStoreUpdate(t *testing.T) {
	tableName := fake.CharactersN(5)
	pk := fake.CharactersN(10)
	sk := fake.CharactersN(10)
	value := fake.CharactersN(10)

	ex, _ := expression.NewBuilder().
		WithCondition(expression.Name("foo").Equal(expression.Value(5))).
		WithFilter(expression.Name("bar").LessThan(expression.Value(6))).
		WithProjection(expression.NamesList(expression.Name("foo"), expression.Name("bar"), expression.Name("baz"))).
		WithKeyCondition(expression.Key("foo").Equal(expression.Value(5))).
		WithUpdate(expression.Set(expression.Name("foo"), expression.Value(value))).
		Build()

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

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockDynamoDBer)

		expectedStruct testStruct
		expectedErr    error
	}{
		{
			desc: "given record is found in store, output struct set",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				result := dynamodb.UpdateItemOutput{
					Attributes: map[string]types.AttributeValue{
						"foo": &types.AttributeValueMemberS{
							Value: value,
						},
					},
				}

				client.EXPECT().UpdateItem(gomock.Any(), &input).Return(&result, nil)
			},

			expectedStruct: testStruct{
				Foo: value,
			},
		},
		{
			desc: "given nil result is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				client.EXPECT().UpdateItem(gomock.Any(), &input).Return(nil, nil)
			},

			expectedErr: fmt.Errorf("nil results returned"),
		},
		{
			desc: "given a ThroughputExeeded error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				client.EXPECT().UpdateItem(gomock.Any(), &input).Return(nil, &types.ProvisionedThroughputExceededException{
					Message: aws.String("something went wrong"),
				})
			},

			expectedErr: fmt.Errorf("error returned from dynamodb ProvisionedThroughputExceededException: something went wrong"),
		},
		{
			desc: "given a InternalServerError error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				client.EXPECT().UpdateItem(gomock.Any(), &input).Return(nil, &types.InternalServerError{
					Message: aws.String("something went wrong"),
				})
			},

			expectedErr: fmt.Errorf("error returned from dynamodb InternalServerError: something went wrong"),
		},
		{
			desc: "given an error is returned, err returned",
			clientExpectations: func(client *storemocks.MockDynamoDBer) {
				client.EXPECT().UpdateItem(gomock.Any(), &input).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error returned from dynamodb something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeMock := storemocks.NewMockDynamoDBer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(storeMock)
			}

			store := store.NewStore(store.WithDynamoDB(storeMock))
			ctx := context.Background()
			m := testStruct{}

			err := store.Update(ctx, tableName, pk, sk, ex, &m)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedStruct, m)
			}
		})
	}
}
