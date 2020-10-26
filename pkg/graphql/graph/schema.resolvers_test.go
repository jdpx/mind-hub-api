package graph_test

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	builder "github.com/jdpx/mind-hub-api/pkg/graphcms/builders"
	graphcmsmocks "github.com/jdpx/mind-hub-api/pkg/graphcms/mocks"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/generated"
	"github.com/stretchr/testify/assert"
)

func TestCoursesResolver(t *testing.T) {
	coursesQuery := "{ courses { id title description sessions { id title description } } }"
	cmsCourse := builder.NewCourseBuilder().
		Build()

	session := builder.NewSessionBuilder().
		Build()

	testCases := []struct {
		desc               string
		query              string
		clientExpectations func(client *graphcmsmocks.MockResolverer)

		expectedErr error
	}{
		{
			desc:  "given a valid courses query with session body, correct courses returned",
			query: coursesQuery,
			clientExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourses(gomock.Any()).Return([]*graphcms.Course{
					&cmsCourse,
				}, nil)
				client.EXPECT().ResolveCourseSessions(gomock.Any(), cmsCourse.ID).Return([]*graphcms.Session{
					&session,
				}, nil)
			},
		},
		{
			desc:  "given the course resolver returns an error, error returned",
			query: coursesQuery,
			clientExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourses(gomock.Any()).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("[{\"message\":\"something went wrong\",\"path\":[\"courses\"]}]"),
		},
		{
			desc:  "given the session resolver returns an error, error returned",
			query: coursesQuery,
			clientExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourses(gomock.Any()).Return([]*graphcms.Course{
					&cmsCourse,
				}, nil)
				client.EXPECT().ResolveCourseSessions(gomock.Any(), cmsCourse.ID).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("[{\"message\":\"something went wrong\",\"path\":[\"courses\",0,\"sessions\"]}]"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockResolverer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := graph.NewResolver(
				graph.WithCMSClient(clientMock),
			)

			c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

			var resp graph.CoursesResponse
			err := c.Post(tt.query, &resp)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cmsCourse.ID, resp.Courses[0].ID)
				assert.Equal(t, cmsCourse.Title, resp.Courses[0].Title)
				assert.Equal(t, cmsCourse.Description, resp.Courses[0].Description)
				assert.Equal(t, session.ID, resp.Courses[0].Sessions[0].ID)
				assert.Equal(t, session.Title, resp.Courses[0].Sessions[0].Title)
				assert.Equal(t, session.Description, resp.Courses[0].Sessions[0].Description)
			}
		})
	}
}

func TestCourseResolver(t *testing.T) {
	courseID := fake.CharactersN(10)

	coursesQuery := fmt.Sprintf(`{ course(where: {id: "%s"}) { id title description }}`, courseID)
	cmsCourse := builder.NewCourseBuilder().
		WithID(courseID).
		Build()

	testCases := []struct {
		desc               string
		query              string
		clientExpectations func(client *graphcmsmocks.MockResolverer)

		expectedErr error
	}{
		{
			desc:  "given a valid course query, correct course returned",
			query: coursesQuery,
			clientExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourse(gomock.Any(), courseID).Return(&cmsCourse, nil)
			},
		},
		{
			desc:  "given the query returns an error, error returned",
			query: coursesQuery,
			clientExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveCourse(gomock.Any(), courseID).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("[{\"message\":\"something went wrong\",\"path\":[\"course\"]}]"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockResolverer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := graph.NewResolver(
				graph.WithCMSClient(clientMock),
			)

			c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

			var resp graph.CourseResponse
			err := c.Post(tt.query, &resp)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cmsCourse.ID, resp.Course.ID)
				assert.Equal(t, cmsCourse.Title, resp.Course.Title)
				assert.Equal(t, cmsCourse.Description, resp.Course.Description)
			}
		})
	}
}

func TestSessionResolver(t *testing.T) {
	sessionID := fake.CharactersN(10)

	sessionsQuery := fmt.Sprintf(`{ session(where: {id: "%s"}) { id title description }}`, sessionID)

	session := builder.NewSessionBuilder().
		WithID(sessionID).
		Build()

	testCases := []struct {
		desc               string
		query              string
		clientExpectations func(client *graphcmsmocks.MockResolverer)

		expectedErr error
	}{
		{
			desc:  "given a valid session query, correct session returned",
			query: sessionsQuery,
			clientExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveSession(gomock.Any(), sessionID).Return(&session, nil)
			},
		},
		{
			desc:  "given the query returns an error, error returned",
			query: sessionsQuery,
			clientExpectations: func(client *graphcmsmocks.MockResolverer) {
				client.EXPECT().ResolveSession(gomock.Any(), sessionID).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("[{\"message\":\"something went wrong\",\"path\":[\"session\"]}]"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockResolverer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := graph.NewResolver(
				graph.WithCMSClient(clientMock),
			)

			c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

			var resp graph.SessionResponse
			err := c.Post(tt.query, &resp)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, session.ID, resp.Session.ID)
				assert.Equal(t, session.Title, resp.Session.Title)
				assert.Equal(t, session.Description, resp.Session.Description)
			}
		})
	}
}
