package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	gBuilder "github.com/jdpx/mind-hub-api/pkg/graphcms/builder"
	graphcmsmocks "github.com/jdpx/mind-hub-api/pkg/graphcms/mocks"
	"github.com/jdpx/mind-hub-api/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestSessionServiceGetByID(t *testing.T) {
	sID := fake.CharactersN(10)

	cmsSession := gBuilder.NewSessionBuilder().
		WithID(sID).
		Build()

	testCases := []struct {
		desc            string
		cmsExpectations func(client *graphcmsmocks.MockCMSResolver)

		expectedSession *service.Session
		expectedErr     error
	}{
		{
			desc: "given step returned from store, step returned",
			cmsExpectations: func(client *graphcmsmocks.MockCMSResolver) {
				client.EXPECT().GetSessionByID(gomock.Any(), sID).Return(&cmsSession, nil)
			},

			expectedSession: &service.Session{
				ID:          cmsSession.ID,
				Title:       cmsSession.Title,
				Description: cmsSession.Description,
				Steps:       []*service.Step{},
			},
		},
		{
			desc: "given a nil step is returned from graphcms, return ErroNotFound",
			cmsExpectations: func(client *graphcmsmocks.MockCMSResolver) {
				client.EXPECT().GetSessionByID(gomock.Any(), sID).Return(nil, nil)
			},

			expectedErr: service.ErrNotFound,
		},
		{
			desc: "given an error is returned from the when getting step, err returned",
			cmsExpectations: func(client *graphcmsmocks.MockCMSResolver) {
				client.EXPECT().GetSessionByID(gomock.Any(), sID).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cmsMock := graphcmsmocks.NewMockCMSResolver(ctrl)

			if tt.cmsExpectations != nil {
				tt.cmsExpectations(cmsMock)
			}

			resolver := service.NewSessionService(cmsMock)
			ctx := context.Background()

			n, err := resolver.GetByID(ctx, sID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedSession, n)
			}
		})
	}
}

func TestSessionServiceGetByCourseID(t *testing.T) {
	cID := fake.CharactersN(10)
	cmsSessionOne := gBuilder.NewSessionBuilder().
		Build()
	cmsSessionTwo := gBuilder.NewSessionBuilder().
		Build()

	testCases := []struct {
		desc            string
		cmsExpectations func(client *graphcmsmocks.MockCMSResolver)

		expectedSession []*service.Session
		expectedErr     error
	}{
		{
			desc: "given steps returned from graphcms, number of steps returned",
			cmsExpectations: func(client *graphcmsmocks.MockCMSResolver) {
				client.EXPECT().GetSessionsByCourseID(gomock.Any(), cID).Return([]*graphcms.Session{
					&cmsSessionOne,
					&cmsSessionTwo,
				}, nil)
			},

			expectedSession: []*service.Session{
				&service.Session{
					ID:          cmsSessionOne.ID,
					Title:       cmsSessionOne.Title,
					Description: cmsSessionOne.Description,
					Steps:       []*service.Step{},
				},
				&service.Session{
					ID:          cmsSessionTwo.ID,
					Title:       cmsSessionTwo.Title,
					Description: cmsSessionTwo.Description,
					Steps:       []*service.Step{},
				},
			},
		},
		{
			desc: "given an error is returned from the when getting step step ids, err returned",
			cmsExpectations: func(client *graphcmsmocks.MockCMSResolver) {
				client.EXPECT().GetSessionsByCourseID(gomock.Any(), cID).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cmsMock := graphcmsmocks.NewMockCMSResolver(ctrl)

			if tt.cmsExpectations != nil {
				tt.cmsExpectations(cmsMock)
			}

			resolver := service.NewSessionService(cmsMock)
			ctx := context.Background()

			n, err := resolver.GetByCourseID(ctx, cID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedSession, n)
			}
		})
	}
}

func TestSessionServiceCountByCourseID(t *testing.T) {
	cID := fake.CharactersN(10)
	cmsSessionOne := gBuilder.NewSessionBuilder().
		Build()
	cmsSessionTwo := gBuilder.NewSessionBuilder().
		Build()

	testCases := []struct {
		desc            string
		cmsExpectations func(client *graphcmsmocks.MockCMSResolver)

		expectedSessionCount int
		expectedErr          error
	}{
		{
			desc: "given steps returned from graphcms, number of steps returned",
			cmsExpectations: func(client *graphcmsmocks.MockCMSResolver) {
				client.EXPECT().GetSessionsByCourseID(gomock.Any(), cID).Return([]*graphcms.Session{
					&cmsSessionOne,
					&cmsSessionTwo,
				}, nil)
			},

			expectedSessionCount: 2,
		},
		{
			desc: "given an error is returned from the when getting step step ids, err returned",
			cmsExpectations: func(client *graphcmsmocks.MockCMSResolver) {
				client.EXPECT().GetSessionsByCourseID(gomock.Any(), cID).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cmsMock := graphcmsmocks.NewMockCMSResolver(ctrl)

			if tt.cmsExpectations != nil {
				tt.cmsExpectations(cmsMock)
			}

			resolver := service.NewSessionService(cmsMock)
			ctx := context.Background()

			n, err := resolver.CountByCourseID(ctx, cID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedSessionCount, n)
			}
		})
	}
}
