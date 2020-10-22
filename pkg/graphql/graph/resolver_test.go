package graph_test

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang/mock/gomock"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	builder "github.com/jdpx/mind-hub-api/pkg/graphcms/builders"
	graphcmsmocks "github.com/jdpx/mind-hub-api/pkg/graphcms/mocks"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/generated"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCoursesResolver(t *testing.T) {
	coursesQuery := "{ courses { title } }"
	req := graphcms.NewRequest(coursesQuery)
	title := "some title"
	cmsCourse := builder.NewCourseBuilder().
		WithTitle(title).
		Build()
	resp := graph.CoursesResponse{
		Courses: []*model.Course{
			&cmsCourse,
		},
	}

	testCases := []struct {
		desc               string
		query              string
		clientExpectations func(client *graphcmsmocks.MockCMSRequster)

		expectedErr error
	}{
		{
			desc:  "given a valid courses query, correct courses returned",
			query: coursesQuery,
			clientExpectations: func(client *graphcmsmocks.MockCMSRequster) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).SetArg(2, resp)
			},
		},
		{
			desc:  "given the query returns an error, error returned",
			query: coursesQuery,
			clientExpectations: func(client *graphcmsmocks.MockCMSRequster) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).Return(fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("[{\"message\":\"something went wrong\",\"path\":[\"courses\"]}]"),
		},
		{
			desc:  "given an invalid query, error returned",
			query: "{{{",

			expectedErr: fmt.Errorf("http 422: {\"errors\":[{\"message\":\"Expected Name, found {\",\"locations\":[{\"line\":1,\"column\":2}],\"extensions\":{\"code\":\"GRAPHQL_PARSE_FAILED\"}}],\"data\":null}"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockCMSRequster(ctrl)

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
				require.Equal(t, title, resp.Courses[0].Title)
			}
		})
	}
}

func TestCourseResolver(t *testing.T) {
	coursesQuery := "{  course(where: {id: 1}) {    title  }}"
	req := graphcms.NewRequest(coursesQuery)
	title := "some title"
	cmsCourse := builder.NewCourseBuilder().
		WithTitle(title).
		Build()
	resp := graph.CourseResponse{
		Course: &cmsCourse,
	}

	testCases := []struct {
		desc               string
		query              string
		clientExpectations func(client *graphcmsmocks.MockCMSRequster)

		expectedErr error
	}{
		{
			desc:  "given a valid courses query, correct courses returned",
			query: coursesQuery,
			clientExpectations: func(client *graphcmsmocks.MockCMSRequster) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).SetArg(2, resp)
			},
		},
		{
			desc:  "given the query returns an error, error returned",
			query: coursesQuery,
			clientExpectations: func(client *graphcmsmocks.MockCMSRequster) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).Return(fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("[{\"message\":\"something went wrong\",\"path\":[\"course\"]}]"),
		},
		{
			desc:  "given an invalid query, error returned",
			query: "{{{",

			expectedErr: fmt.Errorf("http 422: {\"errors\":[{\"message\":\"Expected Name, found {\",\"locations\":[{\"line\":1,\"column\":2}],\"extensions\":{\"code\":\"GRAPHQL_PARSE_FAILED\"}}],\"data\":null}"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockCMSRequster(ctrl)

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
				assert.Equal(t, title, resp.Course.Title)
			}
		})
	}
}

func TestSessionsResolver(t *testing.T) {
	sessionsQuery := "{ sessions { title } }"
	req := graphcms.NewRequest(sessionsQuery)
	title := "some title"
	session := builder.NewSessionBuilder().
		WithTitle(title).
		Build()
	resp := graph.SessionsResponse{
		Sessions: []*model.Session{
			&session,
		},
	}

	testCases := []struct {
		desc               string
		query              string
		clientExpectations func(client *graphcmsmocks.MockCMSRequster)

		expectedErr error
	}{
		{
			desc:  "given a valid sessions query, correct sessions returned",
			query: sessionsQuery,
			clientExpectations: func(client *graphcmsmocks.MockCMSRequster) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).SetArg(2, resp)
			},
		},
		{
			desc:  "given the query returns an error, error returned",
			query: sessionsQuery,
			clientExpectations: func(client *graphcmsmocks.MockCMSRequster) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).Return(fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("[{\"message\":\"something went wrong\",\"path\":[\"sessions\"]}]"),
		},
		{
			desc:  "given an invalid query, error returned",
			query: "{{{",

			expectedErr: fmt.Errorf("http 422: {\"errors\":[{\"message\":\"Expected Name, found {\",\"locations\":[{\"line\":1,\"column\":2}],\"extensions\":{\"code\":\"GRAPHQL_PARSE_FAILED\"}}],\"data\":null}"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockCMSRequster(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := graph.NewResolver(
				graph.WithCMSClient(clientMock),
			)

			c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

			var resp graph.SessionsResponse
			err := c.Post(tt.query, &resp)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				require.Equal(t, title, resp.Sessions[0].Title)
			}
		})
	}
}

func TestSessionResolver(t *testing.T) {
	sessionsQuery := "{  session(where: {id: 1}) {    title  }}"
	req := graphcms.NewRequest(sessionsQuery)
	title := "some title"
	session := builder.NewSessionBuilder().
		WithTitle(title).
		Build()
	resp := graph.SessionResponse{
		Session: &session,
	}

	testCases := []struct {
		desc               string
		query              string
		clientExpectations func(client *graphcmsmocks.MockCMSRequster)

		expectedErr error
	}{
		{
			desc:  "given a valid session query, correct session returned",
			query: sessionsQuery,
			clientExpectations: func(client *graphcmsmocks.MockCMSRequster) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).SetArg(2, resp)
			},
		},
		{
			desc:  "given the query returns an error, error returned",
			query: sessionsQuery,
			clientExpectations: func(client *graphcmsmocks.MockCMSRequster) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).Return(fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("[{\"message\":\"something went wrong\",\"path\":[\"session\"]}]"),
		},
		{
			desc:  "given an invalid query, error returned",
			query: "{{{",

			expectedErr: fmt.Errorf("http 422: {\"errors\":[{\"message\":\"Expected Name, found {\",\"locations\":[{\"line\":1,\"column\":2}],\"extensions\":{\"code\":\"GRAPHQL_PARSE_FAILED\"}}],\"data\":null}"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockCMSRequster(ctrl)

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
				require.Equal(t, title, resp.Session.Title)
			}
		})
	}
}

func TestStepResolver(t *testing.T) {
	stepsQuery := "{  step(where: {id: 1}) { title  }}"
	req := graphcms.NewRequest(stepsQuery)
	title := "some title"
	step := builder.NewStepBuilder().
		WithTitle(title).
		Build()
	resp := graph.StepResponse{
		Step: &step,
	}

	testCases := []struct {
		desc               string
		query              string
		clientExpectations func(client *graphcmsmocks.MockCMSRequster)

		expectedErr error
	}{
		{
			desc:  "given a valid step query, correct step returned",
			query: stepsQuery,
			clientExpectations: func(client *graphcmsmocks.MockCMSRequster) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).SetArg(2, resp)
			},
		},
		{
			desc:  "given the query returns an error, error returned",
			query: stepsQuery,
			clientExpectations: func(client *graphcmsmocks.MockCMSRequster) {
				client.EXPECT().Run(gomock.Any(), req, gomock.Any()).Return(fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("[{\"message\":\"something went wrong\",\"path\":[\"step\"]}]"),
		},
		{
			desc:  "given an invalid query, error returned",
			query: "{{{",

			expectedErr: fmt.Errorf("http 422: {\"errors\":[{\"message\":\"Expected Name, found {\",\"locations\":[{\"line\":1,\"column\":2}],\"extensions\":{\"code\":\"GRAPHQL_PARSE_FAILED\"}}],\"data\":null}"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockCMSRequster(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := graph.NewResolver(
				graph.WithCMSClient(clientMock),
			)

			c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

			var resp graph.StepResponse
			err := c.Post(tt.query, &resp)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				require.Equal(t, title, resp.Step.Title)
			}
		})
	}
}
