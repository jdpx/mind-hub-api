package graphql_test

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	builder "github.com/jdpx/mind-hub-api/pkg/graphcms/builders"
	graphcmsmocks "github.com/jdpx/mind-hub-api/pkg/graphcms/mocks"
	"github.com/jdpx/mind-hub-api/pkg/graphql"
	"github.com/jdpx/mind-hub-api/pkg/graphql/generated"
	"github.com/jdpx/mind-hub-api/tools/client"
	"github.com/stretchr/testify/assert"
)

func TestCoursesResolver(t *testing.T) {
	coursesQuery := "{ courses { id title description sessionCount sessions { id title description } } }"
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
				}, nil).AnyTimes()
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
				client.EXPECT().ResolveCourseSessions(gomock.Any(), cmsCourse.ID).Return(nil, fmt.Errorf("something went wrong")).AnyTimes()
			},

			expectedErr: fmt.Errorf("[{\"message\":\"something went wrong\",\"path\":[\"courses\",0,\"sessionCount\"]},{\"message\":\"something went wrong\",\"path\":[\"courses\",0,\"sessions\"]}]"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := graphcmsmocks.NewMockResolverer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := graphql.NewResolver(
				graphql.WithCMSClient(clientMock),
			)

			c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

			var resp graphql.CoursesResponse
			err := c.Post(tt.query, &resp)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cmsCourse.ID, resp.Courses[0].ID)
				assert.Equal(t, cmsCourse.Title, resp.Courses[0].Title)
				assert.Equal(t, cmsCourse.Description, resp.Courses[0].Description)
				assert.Equal(t, 1, resp.Courses[0].SessionCount)
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

			resolver := graphql.NewResolver(
				graphql.WithCMSClient(clientMock),
			)

			c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

			var resp graphql.CourseResponse
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

			resolver := graphql.NewResolver(
				graphql.WithCMSClient(clientMock),
			)

			c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

			var resp graphql.SessionResponse
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

// func TestCourseStartedResolver(t *testing.T) {
// 	courseID := fake.CharactersN(10)
// 	testUserID := tools.GenerateTestUserID()
// 	uS := strings.Split(testUserID, "|")
// 	userID := uS[1]

// 	mutation := `mutation courseStarted ($courseID: ID!) { courseStarted(input: {courseID: $courseID}) { id progress { started } } }`

// 	event := store.Progress{
// 		CourseID: courseID,
// 		UserID:   userID,
// 	}

// 	dbProgress := store.Progress{
// 		CourseID: courseID,
// 		UserID:   userID,
// 	}

// 	testCases := []struct {
// 		desc               string
// 		query              string
// 		tokenClaims        jwt.MapClaims
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedErr error
// 	}{
// 		{
// 			desc:  "given a valid Course Started mutation, event stored in database",
// 			query: mutation,
// 			tokenClaims: jwt.MapClaims{
// 				"sub": testUserID,
// 			},
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				client.EXPECT().Put(gomock.Any(), "course_progress", event, gomock.Any()).SetArg(3, dbProgress)
// 				client.EXPECT().Get(gomock.Any(), "course_progress", gomock.Any(), gomock.Any()).SetArg(3, dbProgress)
// 			},
// 		},
// 		{
// 			desc:        "given there is no user ID in the request token, error returned",
// 			query:       mutation,
// 			tokenClaims: jwt.MapClaims{},

// 			expectedErr: fmt.Errorf("[{\"message\":\"error occurred getting request user ID token user ID is an invalid Auth0 user ID\",\"path\":[\"courseStarted\"]}]"),
// 		},
// 		{
// 			desc:  "given an error occurs when saving the event to the store, error returned",
// 			query: mutation,
// 			tokenClaims: jwt.MapClaims{
// 				"sub": testUserID,
// 			},
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				client.EXPECT().Put(gomock.Any(), "course_progress", event, gomock.Any()).Return(fmt.Errorf("something went wrong"))
// 			},

// 			expectedErr: fmt.Errorf("[{\"message\":\"something went wrong\",\"path\":[\"courseStarted\"]}]"),
// 		},
// 	}
// 	for _, tt := range testCases {
// 		t.Run(tt.desc, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			clientMock := storemocks.NewMockStorer(ctrl)

// 			if tt.clientExpectations != nil {
// 				tt.clientExpectations(clientMock)
// 			}

// 			courseRepo := store.NewCourseProgressHandler(clientMock)

// 			resolver := graphql.NewResolver(
// 				graphql.WithStore(clientMock),
// 				graphql.WithCourseProgressHandler(courseRepo),
// 			)

// 			h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

// 			h.AroundOperations(func(ctx context.Context, next gqlgen.OperationHandler) gqlgen.ResponseHandler {
// 				return next(tools.GenerateTestGinContext(ctx, tt.tokenClaims))
// 			})

// 			c := client.New(h)

// 			var resp graphql.CourseStartedResponse

// 			err := c.Post(tt.query, &resp, client.Var("courseID", courseID))

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.Equal(t, courseID, resp.CourseStarted.ID)
// 				assert.True(t, resp.CourseStarted.Progress.Started)
// 			}
// 		})
// 	}
// }
