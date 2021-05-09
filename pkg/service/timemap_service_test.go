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
	cID := fake.CharactersN(10)
	tID := fake.CharactersN(10)
	timemap := builder.NewTimemapBuilder().
		WithCourseID(cID).
		WithUserID(uID).
		WithID(tID).
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
					cID,
					tID,
				).Return(&timemap, nil)
			},

			expectedTimemap: &service.Timemap{
				ID:          timemap.ID,
				CourseID:    timemap.CourseID,
				UserID:      timemap.UserID,
				Map:         timemap.Map,
				DateCreated: timemap.DateCreated,
				DateUpdated: timemap.DateUpdated,
			},
		},
		{
			desc: "given an error is returned from the store, err returned",
			clientExpectations: func(client *storemocks.MockTimemapRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					uID,
					cID,
					tID,
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
					cID,
					tID,
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

func TestTimemapServiceGetByCourseID(t *testing.T) {
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

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockTimemapRepositor)

		expectedTimemap []service.Timemap
		expectedErr     error
	}{
		{
			desc: "given map returned from store, map returned",
			clientExpectations: func(client *storemocks.MockTimemapRepositor) {
				client.EXPECT().GetByCourseID(
					gomock.Any(),
					uID,
					cID,
				).Return(timemaps, nil)
			},

			expectedTimemap: []service.Timemap{
				{
					ID:          timemapOne.ID,
					CourseID:    timemapOne.CourseID,
					UserID:      timemapOne.UserID,
					Map:         timemapOne.Map,
					DateCreated: timemapOne.DateCreated,
					DateUpdated: timemapOne.DateUpdated,
				},
				{
					ID:          timemapTwo.ID,
					CourseID:    timemapTwo.CourseID,
					UserID:      timemapTwo.UserID,
					Map:         timemapTwo.Map,
					DateCreated: timemapTwo.DateCreated,
					DateUpdated: timemapTwo.DateUpdated,
				},
			},
		},
		{
			desc: "given an error is returned from the store, err returned",
			clientExpectations: func(client *storemocks.MockTimemapRepositor) {
				client.EXPECT().GetByCourseID(
					gomock.Any(),
					uID,
					cID,
				).Return([]store.Timemap{}, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
		{
			desc: "given the returned an empty slice of timemaps, empty slice returned",
			clientExpectations: func(client *storemocks.MockTimemapRepositor) {
				client.EXPECT().GetByCourseID(
					gomock.Any(),
					uID,
					cID,
				).Return([]store.Timemap{}, nil)
			},

			expectedTimemap: []service.Timemap{},
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

func TestTimemapServiceUpdate(t *testing.T) {
	uID := fake.CharactersN(10)
	cID := fake.CharactersN(10)
	tID := fake.CharactersN(10)
	value := fake.CharactersN(10)

	timemap := builder.NewTimemapBuilder().
		WithCourseID(cID).
		WithUserID(uID).
		WithID(tID).
		WithMap(value).
		Build()

	uTimemap := store.Timemap{
		ID:       tID,
		CourseID: cID,
		UserID:   uID,
		Map:      value,
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
					uTimemap,
				).Return(&timemap, nil)
			},

			expectedTimemap: &service.Timemap{
				ID:          timemap.ID,
				CourseID:    timemap.CourseID,
				UserID:      timemap.UserID,
				Map:         timemap.Map,
				DateCreated: timemap.DateCreated,
				DateUpdated: timemap.DateUpdated,
			},
		},
		{
			desc: "given an error is returned from the store, err returned",
			clientExpectations: func(client *storemocks.MockTimemapRepositor) {
				client.EXPECT().Update(
					gomock.Any(),
					uTimemap,
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

			n, err := resolver.Update(ctx, uID, cID, &tID, value)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedTimemap, n)
			}
		})
	}
}
