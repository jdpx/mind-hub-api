package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/service"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/jdpx/mind-hub-api/pkg/store/builder"
	storemocks "github.com/jdpx/mind-hub-api/pkg/store/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTimemapServiceGet(t *testing.T) {
	uID := fake.CharactersN(10)
	timemap := builder.NewTimemapBuilder().
		WithUserID(uID).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockTimemapRepositor)

		expectedTimemap *service.Timemap
		expectedErr     error
	}{
		{
			desc: "given map returned from store, map returned",
			clientExpectations: func(client *storemocks.MockTimemapRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					uID,
				).Return(&timemap, nil)
			},

			expectedTimemap: &service.Timemap{
				ID:          timemap.ID,
				UserID:      timemap.UserID,
				Map:         timemap.Map,
				DateUpdated: timemap.DateUpdated,
			},
		},
		{
			desc: "given an error is returned from the store, err returned",
			clientExpectations: func(client *storemocks.MockTimemapRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					uID,
				).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
		{
			desc: "given the returned entity is nil, ErrNotFound returned",
			clientExpectations: func(client *storemocks.MockTimemapRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					uID,
				).Return(nil, nil)
			},

			expectedErr: service.ErrNotFound,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockTimemapRepositor(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := service.NewTimemapService(clientMock)
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

func TestTimemapServiceUpdate(t *testing.T) {
	uID := fake.CharactersN(10)
	value := fake.CharactersN(10)
	timemap := builder.NewTimemapBuilder().
		WithUserID(uID).
		WithMap(value).
		Build()

	uTimemap := store.Timemap{
		UserID: uID,
		Map:    value,
	}

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockTimemapRepositor)

		expectedTimemap *service.Timemap
		expectedErr     error
	}{
		{
			desc: "given map returned from store, map returned",
			clientExpectations: func(client *storemocks.MockTimemapRepositor) {
				client.EXPECT().Update(
					gomock.Any(),
					&uTimemap,
				).Return(&timemap, nil)
			},

			expectedTimemap: &service.Timemap{
				ID:          timemap.ID,
				UserID:      timemap.UserID,
				Map:         timemap.Map,
				DateUpdated: timemap.DateUpdated,
			},
		},
		{
			desc: "given an error is returned from the store, err returned",
			clientExpectations: func(client *storemocks.MockTimemapRepositor) {
				client.EXPECT().Update(
					gomock.Any(),
					&uTimemap,
				).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockTimemapRepositor(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := service.NewTimemapService(clientMock)
			ctx := context.Background()

			n, err := resolver.Update(ctx, uID, value)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedTimemap, n)
			}
		})
	}
}
