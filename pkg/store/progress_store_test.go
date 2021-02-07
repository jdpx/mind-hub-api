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

func TestProgressStoreGet(t *testing.T) {
	cID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	progress := builder.NewProgressBuilder().
		WithEntityID(cID).
		WithUserID(uID).
		Build()

	testCases := []struct {
		desc               string
		userID             string
		courseID           string
		clientExpectations func(client *storemocks.MockStorer)

		expectedProgress *store.Progress
		expectedErr      error
	}{
		{
			desc:     "given a progress is returned from store, progress is returned",
			userID:   uID,
			courseID: cID,
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Get(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("PROGRESS#%s", cID),
					gomock.Any(),
				).SetArg(4, progress)
			},

			expectedProgress: &progress,
		},
		{
			desc:     "given a NotFound error is returned from the store, nil map is returned",
			userID:   uID,
			courseID: cID,
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Get(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("PROGRESS#%s", cID),
					gomock.Any(),
				).Return(store.ErrNotFound)
			},

			expectedProgress: nil,
		},
		{
			desc:     "given a generic error is returned from the store, error returned",
			userID:   uID,
			courseID: cID,
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Get(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("PROGRESS#%s", cID),
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

			resolver := store.NewProgressStore(clientMock)
			ctx := context.Background()

			n, err := resolver.Get(ctx, tt.courseID, tt.userID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedProgress, n)
			}
		})
	}
}

func TestProgressStoreCreate(t *testing.T) {
	now := time.Now()
	id := fake.CharactersN(10)
	eID := fake.CharactersN(10)
	uID := fake.CharactersN(10)

	progress := builder.NewProgressBuilder().
		WithPK(fmt.Sprintf("USER#%s", uID)).
		WithSK(fmt.Sprintf("PROGRESS#%s", eID)).
		WithID(id).
		WithEntityID(eID).
		WithUserID(uID).
		WithDateStarted(now).
		Build()

	expectedProgress := builder.NewProgressBuilder().
		WithID(id).
		WithEntityID(eID).
		WithUserID(uID).
		WithDateStarted(now).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockStorer)

		expectedProgress store.Progress
		expectedErr      error
	}{
		{
			desc: "given create is successful, progress is returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Put(gomock.Any(), userTableName, progress).Return(nil)
			},

			expectedProgress: expectedProgress,
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
				return id
			}

			timer := func() time.Time {
				return now
			}

			resolver := store.NewProgressStore(clientMock, store.WithProgressIDGenerator(gen), store.WithProgressTimer(timer))
			ctx := context.Background()

			n, err := resolver.Start(ctx, eID, uID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, &tt.expectedProgress, n)
			}
		})
	}
}

func TestProgressStoreUpdate(t *testing.T) {
	now := time.Now()
	id := fake.CharactersN(10)
	eID := fake.CharactersN(10)
	uID := fake.CharactersN(10)

	expectedProgress := builder.NewProgressBuilder().
		WithID(id).
		WithEntityID(eID).
		WithUserID(uID).
		WithDateStarted(now).
		Completed().
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockStorer)

		expectedProgress *store.Progress
		expectedErr      error
	}{
		{
			desc: "given a progress is returned from store, progress is returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				upBuilder := expression.
					Set(expression.Name("id"), expression.Name("id").IfNotExists(expression.Value(id))).
					Set(expression.Name("entityID"), expression.Name("entityID").IfNotExists(expression.Value(eID))).
					Set(expression.Name("userID"), expression.Name("userID").IfNotExists(expression.Value(uID))).
					Set(expression.Name("state"), expression.Value(store.StatusCompleted)).
					Set(expression.Name("dateStarted"), expression.Name("dateStarted").IfNotExists(expression.Value(now))).
					Set(expression.Name("dateCompleted"), expression.Value(now))

				expr, _ := expression.NewBuilder().WithUpdate(upBuilder).Build()

				client.EXPECT().Update(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("PROGRESS#%s", eID),
					expr,
					gomock.Any(),
				).SetArg(5, expectedProgress)
			},

			expectedProgress: &expectedProgress,
		},
		{
			desc: "given a generic error is returned from the store, error returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				upBuilder := expression.
					Set(expression.Name("id"), expression.Name("id").IfNotExists(expression.Value(id))).
					Set(expression.Name("entityID"), expression.Name("entityID").IfNotExists(expression.Value(eID))).
					Set(expression.Name("userID"), expression.Name("userID").IfNotExists(expression.Value(uID))).
					Set(expression.Name("state"), expression.Value(store.StatusCompleted)).
					Set(expression.Name("dateStarted"), expression.Name("dateStarted").IfNotExists(expression.Value(now))).
					Set(expression.Name("dateCompleted"), expression.Value(now))

				expr, _ := expression.NewBuilder().WithUpdate(upBuilder).Build()

				client.EXPECT().Update(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("PROGRESS#%s", eID),
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
				return id
			}

			timer := func() time.Time {
				return now
			}

			resolver := store.NewProgressStore(clientMock, store.WithProgressIDGenerator(gen), store.WithProgressTimer(timer))
			ctx := context.Background()

			n, err := resolver.Complete(ctx, eID, uID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedProgress, n)
			}
		})
	}
}
