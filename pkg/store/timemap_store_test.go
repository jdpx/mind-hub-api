package store_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/jdpx/mind-hub-api/pkg/store/builder"
	storemocks "github.com/jdpx/mind-hub-api/pkg/store/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTimemapStoreGet(t *testing.T) {
	tID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	cID := fake.CharactersN(10)
	timemap := builder.NewTimemapBuilder().
		WithID(tID).
		WithCourseID(cID).
		WithUserID(uID).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockStorer)

		expectedTimemap *store.Timemap
		expectedErr     error
	}{
		{
			desc: "given a timemap is returned from store, timemap is returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Get(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("COURSE#%s#TIMEMAP#%s", cID, tID),
					gomock.Any(),
				).SetArg(4, timemap)
			},

			expectedTimemap: &timemap,
		},
		{
			desc: "given a NotFound error is returned from the store, nil map is returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Get(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("COURSE#%s#TIMEMAP#%s", cID, tID),
					gomock.Any(),
				).Return(store.ErrNotFound)
			},

			expectedTimemap: nil,
		},
		{
			desc: "given a generic error is returned from the store, error returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Get(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("COURSE#%s#TIMEMAP#%s", cID, tID),
					gomock.Any(),
				).Return(fmt.Errorf("error occurred"))
			},

			expectedErr: fmt.Errorf("error occurred"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockStorer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := store.NewTimemapStore(clientMock)
			ctx := context.Background()

			n, err := resolver.Get(ctx, uID, cID, tID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedTimemap, n)
			}
		})
	}
}

func TestTimemapStoreGetByCourseID(t *testing.T) {
	uID := fake.CharactersN(10)
	cID := fake.CharactersN(10)

	timemapOne := builder.NewTimemapBuilder().
		WithCourseID(cID).
		WithUserID(uID).
		Build()

	timemapTwo := builder.NewTimemapBuilder().
		WithCourseID(cID).
		WithUserID(uID).
		Build()

	timemaps := []store.Timemap{
		timemapOne,
		timemapTwo,
	}

	expectedExpression, _ := expression.NewBuilder().
		WithKeyCondition(expression.Key("PK").Equal(expression.Value(store.UserPK(uID)))).
		WithFilter(expression.Name("SK").BeginsWith(store.CourseTimemapsSK(cID))).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockStorer)

		expectedTimemap []store.Timemap
		expectedErr     error
	}{
		{
			desc: "given a timemap is returned from store, timemap is returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Query(
					gomock.Any(),
					userTableName,
					expectedExpression,
					gomock.Any(),
				).SetArg(3, timemaps)
			},

			expectedTimemap: timemaps,
		},
		{
			desc: "given a generic error is returned from the store, error returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Query(
					gomock.Any(),
					userTableName,
					expectedExpression,
					gomock.Any(),
				).Return(fmt.Errorf("error occurred"))
			},

			expectedErr: fmt.Errorf("error occurred"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockStorer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := store.NewTimemapStore(clientMock)
			ctx := context.Background()

			n, err := resolver.GetByCourseID(ctx, uID, cID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedTimemap, n)
			}
		})
	}
}

func TestTimemapStoreCreate(t *testing.T) {
	now := time.Now()
	uID := fake.CharactersN(10)
	cID := fake.CharactersN(10)
	tID := fake.CharactersN(10)
	timemap := builder.NewTimemapBuilder().
		WithCourseID(cID).
		WithUserID(uID).
		WithID("").
		Build()

	createTimemap := builder.NewTimemapBuilder().
		WithPK(fmt.Sprintf("USER#%s", uID)).
		WithSK(fmt.Sprintf("COURSE#%s#TIMEMAP#%s", cID, tID)).
		WithID(tID).
		WithCourseID(cID).
		WithUserID(uID).
		WithMap(timemap.Map).
		WithDateUpdated(now).
		WithDateCreated(now).
		Build()

	expectedTimemap := builder.NewTimemapBuilder().
		WithCourseID(timemap.CourseID).
		WithUserID(timemap.UserID).
		WithMap(timemap.Map).
		WithID(tID).
		WithDateUpdated(now).
		WithDateCreated(now).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockStorer)

		expectedTimemap store.Timemap
		expectedErr     error
	}{
		{
			desc: "given create is successful, timemap is returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Put(gomock.Any(), userTableName, createTimemap).Return(nil)
			},

			expectedTimemap: expectedTimemap,
		},
		{
			desc: "given a generic error is returned from the store, error returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Put(gomock.Any(), userTableName, gomock.Any()).Return(fmt.Errorf("error occurred"))
			},

			expectedErr: fmt.Errorf("error occurred"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockStorer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			gen := func() string {
				return tID
			}

			timer := func() time.Time {
				return now
			}

			resolver := store.NewTimemapStore(clientMock, store.WithTimemapIDGenerator(gen), store.WithTimemapTimer(timer))
			ctx := context.Background()

			n, err := resolver.Create(ctx, timemap)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, &tt.expectedTimemap, n)
			}
		})
	}
}

func TestTimemapStoreUpdate(t *testing.T) {
	now := time.Now()
	tID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	cID := fake.CharactersN(10)

	existingTimemap := builder.NewTimemapBuilder().
		WithID(tID).
		WithCourseID(cID).
		WithUserID(uID).
		Build()

	timemapWithoutID := builder.NewTimemapBuilder().
		WithID("").
		WithCourseID(cID).
		WithUserID(uID).
		WithMap(existingTimemap.Map).
		Build()

	expectedTimemap := builder.NewTimemapBuilder().
		WithUserID(uID).
		WithCourseID(cID).
		WithID(tID).
		WithMap(existingTimemap.Map).
		WithDateCreated(now).
		WithDateUpdated(now).
		Build()

	expectedExpression := expression.
		Set(expression.Name("id"), expression.Name("id").IfNotExists(expression.Value(tID))).
		Set(expression.Name("courseID"), expression.Name("courseID").IfNotExists(expression.Value(cID))).
		Set(expression.Name("userID"), expression.Name("userID").IfNotExists(expression.Value(uID))).
		Set(expression.Name("map"), expression.Value(existingTimemap.Map)).
		Set(expression.Name("dateCreated"), expression.Name("dateCreated").IfNotExists(expression.Value(now))).
		Set(expression.Name("dateUpdated"), expression.Value(now))

	testCases := []struct {
		desc               string
		timemap            store.Timemap
		clientExpectations func(client *storemocks.MockStorer)

		expectedTimemap *store.Timemap
		expectedErr     error
	}{
		{
			desc:    "given a timemap that has an ID and updated timemap is returned from store, timemap is returned",
			timemap: existingTimemap,
			clientExpectations: func(client *storemocks.MockStorer) {
				expr, _ := expression.NewBuilder().WithUpdate(expectedExpression).Build()

				client.EXPECT().Update(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("COURSE#%s#TIMEMAP#%s", cID, tID),
					expr,
					gomock.Any(),
				).SetArg(5, expectedTimemap)
			},

			expectedTimemap: &expectedTimemap,
		},
		{
			desc:    "given a timemap without an ID and updated timemap is returned from store, timemap is returned",
			timemap: timemapWithoutID,
			clientExpectations: func(client *storemocks.MockStorer) {
				expr, _ := expression.NewBuilder().WithUpdate(expectedExpression).Build()

				client.EXPECT().Update(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("COURSE#%s#TIMEMAP#%s", cID, tID),
					expr,
					gomock.Any(),
				).SetArg(5, expectedTimemap)
			},

			expectedTimemap: &expectedTimemap,
		},
		{
			desc:    "given a generic error is returned from the store, error returned",
			timemap: existingTimemap,
			clientExpectations: func(client *storemocks.MockStorer) {
				expr, _ := expression.NewBuilder().WithUpdate(expectedExpression).Build()

				client.EXPECT().Update(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("COURSE#%s#TIMEMAP#%s", cID, tID),
					expr,
					gomock.Any(),
				).Return(fmt.Errorf("error occurred"))
			},

			expectedErr: fmt.Errorf("error occurred"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockStorer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			gen := func() string {
				return tID
			}

			timer := func() time.Time {
				return now
			}

			resolver := store.NewTimemapStore(
				clientMock,
				store.WithTimemapIDGenerator(gen),
				store.WithTimemapTimer(timer),
			)

			ctx := context.Background()

			n, err := resolver.Update(ctx, tt.timemap)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedTimemap, n)
			}
		})
	}
}
