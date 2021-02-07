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
	"github.com/stretchr/testify/assert"
)

func TestStepServiceGetByID(t *testing.T) {
	sID := fake.CharactersN(10)

	cmsStep := gBuilder.NewStepBuilder().
		WithID(sID).
		Build()

	testCases := []struct {
		desc            string
		cmsExpectations func(client *graphcmsmocks.MockResolverer)

		expectedStep *service.Step
		expectedErr  error
	}{
		{
			desc: "given step returned from store, step returned",
			cmsExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveStep(gomock.Any(), sID).Return(&cmsStep, nil)
			},

			expectedStep: &service.Step{
				ID:          cmsStep.ID,
				Title:       cmsStep.Title,
				Description: cmsStep.Description,
				Type:        cmsStep.Type,
				VideoURL:    cmsStep.VideoURL,
				Question:    cmsStep.Question,
			},
		},
		{
			desc: "given a nil step is returned from graphcms, return ErroNotFound",
			cmsExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveStep(gomock.Any(), sID).Return(nil, nil)
			},

			expectedErr: service.ErrNotFound,
		},
		{
			desc: "given an error is returned from the when getting step, err returned",
			cmsExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveStep(gomock.Any(), sID).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cmsMock := graphcmsmocks.NewMockResolverer(ctrl)

			if tt.cmsExpectations != nil {
				tt.cmsExpectations(cmsMock)
			}

			resolver := service.NewStepService(cmsMock)
			ctx := context.Background()

			n, err := resolver.GetByID(ctx, sID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedStep, n)
			}
		})
	}
}

func TestStepServiceCountByCourseID(t *testing.T) {
	cID := fake.CharactersN(10)
	cmsStepOne := gBuilder.NewStepBuilder().
		Build()
	cmsStepTwo := gBuilder.NewStepBuilder().
		Build()

	testCases := []struct {
		desc            string
		cmsExpectations func(client *graphcmsmocks.MockResolverer)

		expectedStepCount int
		expectedErr       error
	}{
		{
			desc: "given steps returned from graphcms, number of steps returned",
			cmsExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourseStepIDs(gomock.Any(), cID).Return([]string{
					cmsStepOne.ID,
					cmsStepTwo.ID,
				}, nil)
			},

			expectedStepCount: 2,
		},
		{
			desc: "given an error is returned from the when getting step step ids, err returned",
			cmsExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourseStepIDs(gomock.Any(), cID).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cmsMock := graphcmsmocks.NewMockResolverer(ctrl)

			if tt.cmsExpectations != nil {
				tt.cmsExpectations(cmsMock)
			}

			resolver := service.NewStepService(cmsMock)
			ctx := context.Background()

			n, err := resolver.CountByCourseID(ctx, cID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedStepCount, n)
			}
		})
	}
}
