package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	gBuilder "github.com/jdpx/mind-hub-api/pkg/graphcms/builders"
	graphcmsmocks "github.com/jdpx/mind-hub-api/pkg/graphcms/mocks"
	"github.com/jdpx/mind-hub-api/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestCourseServiceGetAll(t *testing.T) {
	cmsCourseOne := gBuilder.NewCourseBuilder().
		Build()
	cmsCourseTwo := gBuilder.NewCourseBuilder().
		Build()

	testCases := []struct {
		desc            string
		cmsExpectations func(client *graphcmsmocks.MockResolverer)

		expectedCourse []*service.Course
		expectedErr    error
	}{
		{
			desc: "given courses returned from graphcms, courses returned",
			cmsExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourses(gomock.Any()).Return([]*graphcms.Course{
					&cmsCourseOne,
					&cmsCourseTwo,
				}, nil)
			},

			expectedCourse: []*service.Course{
				{
					ID:          cmsCourseOne.ID,
					Title:       cmsCourseOne.Title,
					Description: cmsCourseOne.Description,
					Sessions:    []*service.Session{},
				},
				{
					ID:          cmsCourseTwo.ID,
					Title:       cmsCourseTwo.Title,
					Description: cmsCourseTwo.Description,
					Sessions:    []*service.Session{},
				},
			},
		},
		{
			desc: "given an error is returned from the when getting courses from graphcms, err returned",
			cmsExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourses(gomock.Any()).Return(nil, fmt.Errorf("something went wrong"))
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

			resolver := service.NewCourseService(cmsMock)
			ctx := context.Background()

			n, err := resolver.GetAll(ctx)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedCourse, n)
			}
		})
	}
}

func TestCourseServiceGetByID(t *testing.T) {
	cID := fake.CharactersN(10)

	cmsCourse := gBuilder.NewCourseBuilder().
		WithID(cID).
		Build()

	testCases := []struct {
		desc            string
		cmsExpectations func(client *graphcmsmocks.MockResolverer)

		expectedCourse *service.Course
		expectedErr    error
	}{
		{
			desc: "given course returned from store, course returned",

			cmsExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourse(gomock.Any(), cID).Return(&cmsCourse, nil)
			},

			expectedCourse: &service.Course{
				ID:          cmsCourse.ID,
				Title:       cmsCourse.Title,
				Description: cmsCourse.Description,
				Sessions:    []*service.Session{},
			},
		},
		{
			desc: "given a nil course is returned from graphcms, return ErroNotFound",
			cmsExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourse(gomock.Any(), cID).Return(nil, nil)
			},

			expectedErr: service.ErrNotFound,
		},
		{
			desc: "given an error is returned from the when getting course, err returned",
			cmsExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourse(gomock.Any(), cID).Return(nil, fmt.Errorf("something went wrong"))
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

			resolver := service.NewCourseService(cmsMock)
			ctx := context.Background()

			n, err := resolver.GetByID(ctx, cID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedCourse, n)
			}
		})
	}
}
