package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/service"
	"github.com/jdpx/mind-hub-api/pkg/store/builder"
	storemocks "github.com/jdpx/mind-hub-api/pkg/store/mocks"
	"github.com/stretchr/testify/assert"
)

func TestStepProgressServiceGet(t *testing.T) {
	uID := fake.CharactersN(10)
	sID := fake.CharactersN(10)
	progress := builder.NewProgressBuilder().
		WithEntityID(sID).
		WithUserID(uID).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockProgressRepositor)

		expectedStepProgress *service.StepProgress
		expectedErr          error
	}{
		{
			desc: "given progress returned from store, progress returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					sID,
					uID,
				).Return(&progress, nil)
			},

			expectedStepProgress: &service.StepProgress{
				ID:            progress.ID,
				UserID:        progress.UserID,
				StepID:        progress.EntityID,
				State:         progress.State,
				DateStarted:   progress.DateStarted,
				DateCompleted: progress.DateCompleted,
			},
		},
		{
			desc: "given an error is returned from the store, err returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					sID,
					uID,
				).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
		{
			desc: "given the returned entity is nil, ErrNotFound returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					sID,
					uID,
				).Return(nil, nil)
			},

			expectedErr: service.ErrNotFound,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockProgressRepositor(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := service.NewStepProgressService(clientMock)
			ctx := context.Background()

			n, err := resolver.Get(ctx, sID, uID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedStepProgress, n)
			}
		})
	}
}

func TestStepProgressServiceStart(t *testing.T) {
	sID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	progress := builder.NewProgressBuilder().
		WithUserID(uID).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockProgressRepositor)

		expectedStepProgress *service.StepProgress
		expectedErr          error
	}{
		{
			desc: "given progress returned from store, progress returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Start(
					gomock.Any(),
					sID,
					uID,
				).Return(&progress, nil)
			},

			expectedStepProgress: &service.StepProgress{
				ID:          progress.ID,
				StepID:      progress.EntityID,
				UserID:      progress.UserID,
				State:       progress.State,
				DateStarted: progress.DateStarted,
			},
		},
		{
			desc: "given an error is returned from the store, err returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Start(
					gomock.Any(),
					sID,
					uID,
				).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockProgressRepositor(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := service.NewStepProgressService(clientMock)
			ctx := context.Background()

			n, err := resolver.Start(ctx, sID, uID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedStepProgress, n)
			}
		})
	}
}

func TestStepProgressServiceComplete(t *testing.T) {
	sID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	progress := builder.NewProgressBuilder().
		WithUserID(uID).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockProgressRepositor)

		expectedStepProgress *service.StepProgress
		expectedErr          error
	}{
		{
			desc: "given progress returned from store, progress returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Complete(
					gomock.Any(),
					sID,
					uID,
				).Return(&progress, nil)
			},

			expectedStepProgress: &service.StepProgress{
				ID:            progress.ID,
				StepID:        progress.EntityID,
				UserID:        progress.UserID,
				State:         progress.State,
				DateStarted:   progress.DateStarted,
				DateCompleted: progress.DateCompleted,
			},
		},
		{
			desc: "given an error is returned from the store, err returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Complete(
					gomock.Any(),
					sID,
					uID,
				).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockProgressRepositor(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := service.NewStepProgressService(clientMock)
			ctx := context.Background()

			n, err := resolver.Complete(ctx, sID, uID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedStepProgress, n)
			}
		})
	}
}
