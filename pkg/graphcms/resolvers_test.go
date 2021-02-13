package graphcms_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/graphcms/builder"
	graphcmsmocks "github.com/jdpx/mind-hub-api/pkg/graphcms/mocks"
	"github.com/stretchr/testify/assert"
)

func TestResolverResolveCourses(t *testing.T) {
	ctx := context.Background()
	cmsCourse := builder.NewCourseBuilder().
		Build()

	req := graphcms.NewRequest(ctx, `{
  courses {
    id
    title
    description
  }
}`)

	resp := struct {
		Courses []*graphcms.Course `json:"courses"`
	}{
		Courses: []*graphcms.Course{
			&cmsCourse,
		},
	}

	testCases := []struct {
		desc               string
		clientExpectations func(client *graphcmsmocks.MockRequester)

		expectedCourse *graphcms.Course
		expectedErr    error
	}{
		{
			desc: "given a successful request to GraphCMS, Courses model returned",
			clientExpectations: func(client *graphcmsmocks.MockRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).SetArg(2, resp)
			},
		},
		{
			desc: "given an error is returned from the client",
			clientExpectations: func(client *graphcmsmocks.MockRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).Return(fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error occurred getting GraphCMS Courses"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockRequester(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := graphcms.NewResolver(clientMock)
			ctx := context.Background()

			courses, err := resolver.ResolveCourses(ctx)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, &cmsCourse, courses[0])
			}
		})
	}
}

func TestResolverResolveCourse(t *testing.T) {
	ctx := context.Background()
	cmsCourse := builder.NewCourseBuilder().
		Build()

	req := graphcms.NewRequest(ctx, `
  query Course($id: ID) {
      course(where: { id: $id }) {
          id
          title
          description
      }
  }`)
	req.Var("id", cmsCourse.ID)

	resp := struct {
		Course *graphcms.Course `json:"course"`
	}{
		Course: &cmsCourse,
	}

	testCases := []struct {
		desc               string
		clientExpectations func(client *graphcmsmocks.MockRequester)

		expectedCourse *graphcms.Course
		expectedErr    error
	}{
		{
			desc: "given a successful request to GraphCMS, Courses model returned",
			clientExpectations: func(client *graphcmsmocks.MockRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).SetArg(2, resp)
			},
		},
		{
			desc: "given an error is returned from the client",
			clientExpectations: func(client *graphcmsmocks.MockRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).Return(fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error occurred getting GraphCMS Course"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockRequester(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := graphcms.NewResolver(clientMock)
			ctx := context.Background()

			course, err := resolver.ResolveCourse(ctx, cmsCourse.ID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, &cmsCourse, course)
			}
		})
	}
}

func TestResolverResolveCourseSessions(t *testing.T) {
	ctx := context.Background()
	courseID := fake.CharactersN(10)
	cmsSession := builder.NewSessionBuilder().
		Build()

	req := graphcms.NewRequest(ctx, `query sessions($id: ID){
  sessions(where: { course: { id: $id } }) {
    id
    title
    description

    steps {
      id
      title
    }
  }
}`)
	req.Var("id", courseID)

	resp := struct {
		Sessions []*graphcms.Session `json:"sessions"`
	}{
		Sessions: []*graphcms.Session{
			&cmsSession,
		},
	}

	testCases := []struct {
		desc               string
		clientExpectations func(client *graphcmsmocks.MockRequester)

		expectedCourse *graphcms.Course
		expectedErr    error
	}{
		{
			desc: "given a successful request to GraphCMS, Sessions model returned",
			clientExpectations: func(client *graphcmsmocks.MockRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).SetArg(2, resp)
			},
		},
		{
			desc: "given an error is returned from the client",
			clientExpectations: func(client *graphcmsmocks.MockRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).Return(fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error occurred getting GraphCMS Course Sessions"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockRequester(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := graphcms.NewResolver(clientMock)
			ctx := context.Background()

			sessions, err := resolver.ResolveCourseSessions(ctx, courseID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, &cmsSession, sessions[0])
			}
		})
	}
}

func TestResolverResolveSession(t *testing.T) {
	ctx := context.Background()
	cmsSession := builder.NewSessionBuilder().
		Build()

	req := graphcms.NewRequest(ctx, `query Session($id: ID) {
    session(where: { id: $id }) {
        id
        title
        description

        steps {
            id
            title
            description
            type
            videoUrl
            audio {
                id
                url
            }
            question
        }

        course {
            id
            title
            description
        }
    }
}`)
	req.Var("id", cmsSession.ID)

	resp := struct {
		Session *graphcms.Session `json:"session"`
	}{
		Session: &cmsSession,
	}

	testCases := []struct {
		desc               string
		clientExpectations func(client *graphcmsmocks.MockRequester)

		expectedSession *graphcms.Session
		expectedErr     error
	}{
		{
			desc: "given a successful request to GraphCMS, Sessions model returned",
			clientExpectations: func(client *graphcmsmocks.MockRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).SetArg(2, resp)
			},
		},
		{
			desc: "given an error is returned from the client",
			clientExpectations: func(client *graphcmsmocks.MockRequester) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).Return(fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("error occurred getting GraphCMS Session"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockRequester(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := graphcms.NewResolver(clientMock)
			ctx := context.Background()

			session, err := resolver.ResolveSession(ctx, cmsSession.ID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, &cmsSession, session)
			}
		})
	}
}
