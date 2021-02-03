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
	uID := fake.CharactersN(10)
	timemap := builder.NewTimemapBuilder().
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
					"TIMEMAP",
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
					"TIMEMAP",
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
					"TIMEMAP",
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

			n, err := resolver.Get(ctx, uID)

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
	nID := fake.CharactersN(10)
	timemap := builder.NewTimemapBuilder().WithID("").Build()

	createTimemap := builder.NewTimemapBuilder().
		WithPK(fmt.Sprintf("USER#%s", timemap.UserID)).
		WithSK("TIMEMAP").
		WithUserID(timemap.UserID).
		WithMap(timemap.Map).
		WithID(nID).
		WithDateUpdated(now).
		Build()

	expectedTimemap := builder.NewTimemapBuilder().
		WithUserID(timemap.UserID).
		WithMap(timemap.Map).
		WithID(nID).
		WithDateUpdated(now).
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
				return nID
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
	uID := fake.CharactersN(10)
	timemap := builder.NewTimemapBuilder().
		WithUserID(uID).
		Build()

	expectedTimemap := builder.NewTimemapBuilder().
		WithUserID(timemap.UserID).
		WithMap(timemap.Map).
		WithID(timemap.ID).
		WithDateUpdated(now).
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
				upBuilder := expression.
					Set(expression.Name("map"), expression.Value(timemap.Map)).
					Set(expression.Name("dateUpdated"), expression.Value(now))

				expr, _ := expression.NewBuilder().WithUpdate(upBuilder).Build()

				client.EXPECT().Update(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					"TIMEMAP",
					expr,
					gomock.Any(),
				).SetArg(5, expectedTimemap)
			},

			expectedTimemap: &expectedTimemap,
		},
		{
			desc: "given a generic error is returned from the store, error returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				upBuilder := expression.
					Set(expression.Name("map"), expression.Value(timemap.Map)).
					Set(expression.Name("dateUpdated"), expression.Value(now))

				expr, _ := expression.NewBuilder().WithUpdate(upBuilder).Build()

				client.EXPECT().Update(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					"TIMEMAP",
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

			timer := func() time.Time {
				return now
			}

			resolver := store.NewTimemapStore(clientMock, store.WithTimemapTimer(timer))
			ctx := context.Background()

			n, err := resolver.Update(ctx, &timemap)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedTimemap, n)
			}
		})
	}
}

// const (
// 	timemapTableName = "timemap"
// )

// func TestTimemaphandlerGet(t *testing.T) {
// 	uID := fake.CharactersN(10)
// 	timemap := builder.NewTimemapBuilder().WithUserID(uID).Build()

// 	testCases := []struct {
// 		desc               string
// 		userID             string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedTimemap *store.Timemap
// 		expectedErr     error
// 	}{
// 		{
// 			desc:   "given a timemap is returned from store, map is returned",
// 			userID: uID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"userID": uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), timemapTableName, params, gomock.Any()).SetArg(3, timemap)
// 			},

// 			expectedTimemap: &timemap,
// 		},
// 		{
// 			desc:   "given a NotFound error is returned from the store, nil map is returned",
// 			userID: uID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"userID": uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), timemapTableName, params, gomock.Any()).Return(store.ErrNotFound)
// 			},

// 			expectedTimemap: nil,
// 		},
// 		{
// 			desc:   "given a generic error is returned from the store, error returned",
// 			userID: uID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"userID": uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), timemapTableName, params, gomock.Any()).Return(fmt.Errorf("error occurred"))
// 			},

// 			expectedErr: fmt.Errorf("error occurred"),
// 		},
// 	}
// 	for _, tt := range testCases {
// 		t.Run(tt.desc, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			clientMock := storemocks.NewMockStorer(ctrl)

// 			if tt.clientExpectations != nil {
// 				tt.clientExpectations(clientMock)
// 			}

// 			resolver := store.NewTimemapHandler(clientMock)
// 			ctx := context.Background()

// 			tm, err := resolver.Get(ctx, tt.userID)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.Equal(t, tt.expectedTimemap, tm)
// 			}
// 		})
// 	}
// }

// func TestTimemaphandlerCreate(t *testing.T) {
// 	timemap := builder.NewTimemapBuilder().WithID("").Build()

// 	testCases := []struct {
// 		desc               string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedTimemap *store.Timemap
// 		expectedErr     error
// 	}{
// 		{
// 			desc: "given a timemap is returned from store, map is returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				client.EXPECT().Put(gomock.Any(), timemapTableName, gomock.Any()).Return(nil)
// 			},

// 			expectedTimemap: &timemap,
// 		},
// 		{
// 			desc: "given a generic error is returned from the store, error returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				client.EXPECT().Put(gomock.Any(), timemapTableName, gomock.Any()).Return(fmt.Errorf("error occurred"))
// 			},

// 			expectedErr: fmt.Errorf("error occurred"),
// 		},
// 	}
// 	for _, tt := range testCases {
// 		t.Run(tt.desc, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			clientMock := storemocks.NewMockStorer(ctrl)

// 			if tt.clientExpectations != nil {
// 				tt.clientExpectations(clientMock)
// 			}

// 			resolver := store.NewTimemapHandler(clientMock)
// 			ctx := context.Background()

// 			tm, err := resolver.Create(ctx, timemap)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.NotEqual(t, tt.expectedTimemap.ID, tm.ID)
// 				assert.Equal(t, tt.expectedTimemap.UserID, tm.UserID)
// 				assert.Equal(t, tt.expectedTimemap.Map, tm.Map)
// 				assert.Equal(t, tt.expectedTimemap.UpdatedAt, tm.UpdatedAt)
// 			}
// 		})
// 	}
// }

// func TestTimemaphandlerUpdate(t *testing.T) {
// 	timemap := builder.NewTimemapBuilder().Build()

// 	testCases := []struct {
// 		desc               string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedTimemap *store.Timemap
// 		expectedErr     error
// 	}{
// 		{
// 			desc: "given a timemap is returned from store, map is returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				keys := map[string]string{
// 					"userID": timemap.UserID,
// 				}

// 				exp := "set info.map = :map, updatedAt = :updatedAt"

// 				client.EXPECT().Update(gomock.Any(), timemapTableName, keys, exp, gomock.Any(), gomock.Any()).SetArg(5, timemap)
// 			},

// 			expectedTimemap: &timemap,
// 		},
// 		{
// 			desc: "given a generic error is returned from the store, error returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				keys := map[string]string{
// 					"userID": timemap.UserID,
// 				}

// 				exp := "set info.map = :map, updatedAt = :updatedAt"

// 				client.EXPECT().Update(gomock.Any(), timemapTableName, keys, exp, gomock.Any(), gomock.Any()).Return(fmt.Errorf("error occurred"))
// 			},

// 			expectedErr: fmt.Errorf("error occurred"),
// 		},
// 	}
// 	for _, tt := range testCases {
// 		t.Run(tt.desc, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			clientMock := storemocks.NewMockStorer(ctrl)

// 			if tt.clientExpectations != nil {
// 				tt.clientExpectations(clientMock)
// 			}

// 			resolver := store.NewTimemapHandler(clientMock)
// 			ctx := context.Background()

// 			tm, err := resolver.Update(ctx, &timemap)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.Equal(t, tt.expectedTimemap.ID, tm.ID)
// 				assert.Equal(t, tt.expectedTimemap.Map, tm.Map)
// 				assert.NotEqual(t, tt.expectedTimemap.UpdatedAt, tm.UpdatedAt)
// 			}
// 		})
// 	}
// }
