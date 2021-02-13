package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	gBuilder "github.com/jdpx/mind-hub-api/pkg/graphcms/builder"
	graphcmsmocks "github.com/jdpx/mind-hub-api/pkg/graphcms/mocks"
	"github.com/jdpx/mind-hub-api/pkg/service"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/jdpx/mind-hub-api/pkg/store/builder"
	storemocks "github.com/jdpx/mind-hub-api/pkg/store/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCourseProgressServiceGet(t *testing.T) {
	uID := fake.CharactersN(10)
	cID := fake.CharactersN(10)
	progress := builder.NewProgressBuilder().
		WithEntityID(cID).
		WithUserID(uID).
		Build()

	cmsStepOne := gBuilder.NewStepBuilder().
		Build()
	cmsStepTwo := gBuilder.NewStepBuilder().
		Build()

	sProgress := builder.NewProgressBuilder().
		WithEntityID(cmsStepOne.ID).
		WithUserID(uID).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockProgressRepositor)
		cmsExpectations    func(client *graphcmsmocks.MockCMSResolver)

		expectedCourseProgress *service.CourseProgress
		expectedErr            error
	}{
		{
			desc: "given progress returned from store, progress returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					cID,
					uID,
				).Return(&progress, nil)
				client.EXPECT().GetCompletedByIDs(gomock.Any(), uID, cmsStepOne.ID, cmsStepTwo.ID).Return([]*store.Progress{
					&sProgress,
				}, nil)
			},

			cmsExpectations: func(client *graphcmsmocks.MockCMSResolver) {
				client.EXPECT().GetStepIDsByCourseID(gomock.Any(), cID).Return([]string{
					cmsStepOne.ID,
					cmsStepTwo.ID,
				}, nil)
			},

			expectedCourseProgress: &service.CourseProgress{
				ID:             progress.ID,
				UserID:         progress.UserID,
				CourseID:       progress.EntityID,
				State:          progress.State,
				CompletedSteps: 1,
				DateStarted:    progress.DateStarted,
				DateCompleted:  progress.DateCompleted,
			},
		},
		{
			desc: "given an error is returned from the when getting course progress, err returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					cID,
					uID,
				).Return(&progress, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
		{
			desc: "given the returned course progress is nil, ErrNotFound returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					cID,
					uID,
				).Return(nil, nil)
			},

			expectedErr: service.ErrNotFound,
		},
		{
			desc: "given an error is returned from the when getting course step ids, err returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					cID,
					uID,
				).Return(&progress, nil)
			},
			cmsExpectations: func(client *graphcmsmocks.MockCMSResolver) {
				client.EXPECT().GetStepIDsByCourseID(gomock.Any(), cID).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error occurred getting course step ids something went wrong"),
		},
		{
			desc: "given no steps are returned for a course, progress returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					cID,
					uID,
				).Return(&progress, nil)
			},

			cmsExpectations: func(client *graphcmsmocks.MockCMSResolver) {
				client.EXPECT().GetStepIDsByCourseID(gomock.Any(), cID).Return([]string{}, nil)
			},

			expectedCourseProgress: &service.CourseProgress{
				ID:             progress.ID,
				UserID:         progress.UserID,
				CourseID:       progress.EntityID,
				State:          progress.State,
				CompletedSteps: 0,
				DateStarted:    progress.DateStarted,
				DateCompleted:  progress.DateCompleted,
			},
		},
		{
			desc: "given an error occurs getting completed step ids, error returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					cID,
					uID,
				).Return(&progress, nil)
				client.EXPECT().GetCompletedByIDs(gomock.Any(), uID, cmsStepOne.ID, cmsStepTwo.ID).Return(nil, fmt.Errorf("something went wrong"))
			},

			cmsExpectations: func(client *graphcmsmocks.MockCMSResolver) {
				client.EXPECT().GetStepIDsByCourseID(gomock.Any(), cID).Return([]string{
					cmsStepOne.ID,
					cmsStepTwo.ID,
				}, nil)
			},

			expectedErr: fmt.Errorf("error occurred getting course progress something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			clientMock := storemocks.NewMockProgressRepositor(ctrl)
			cmsMock := graphcmsmocks.NewMockCMSResolver(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}
			if tt.cmsExpectations != nil {
				tt.cmsExpectations(cmsMock)
			}

			resolver := service.NewCourseProgressService(cmsMock, clientMock)
			ctx := context.Background()

			n, err := resolver.Get(ctx, cID, uID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedCourseProgress, n)
			}
		})
	}
}

func TestCourseProgressServiceStart(t *testing.T) {
	cID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	progress := builder.NewProgressBuilder().
		WithUserID(uID).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockProgressRepositor)
		cmsExpectations    func(client *graphcmsmocks.MockCMSResolver)

		expectedCourseProgress *service.CourseProgress
		expectedErr            error
	}{
		{
			desc: "given progress returned from store, progress returned",
			clientExpectations: func(client *storemocks.MockProgressRepositor) {
				client.EXPECT().Start(
					gomock.Any(),
					cID,
					uID,
				).Return(&progress, nil)
			},

			expectedCourseProgress: &service.CourseProgress{
				ID:          progress.ID,
				CourseID:    progress.EntityID,
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
					cID,
					uID,
				).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			clientMock := storemocks.NewMockProgressRepositor(ctrl)
			cmsMock := graphcmsmocks.NewMockCMSResolver(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			if tt.cmsExpectations != nil {
				tt.cmsExpectations(cmsMock)
			}

			resolver := service.NewCourseProgressService(cmsMock, clientMock)
			ctx := context.Background()

			n, err := resolver.Start(ctx, cID, uID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedCourseProgress, n)
			}
		})
	}
}
