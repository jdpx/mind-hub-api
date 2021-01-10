package store_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/jdpx/mind-hub-api/pkg/store/builder"
	storemocks "github.com/jdpx/mind-hub-api/pkg/store/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	courseProgressTableName = "course_progress"
)

func TestCourseProgresshandlerGet(t *testing.T) {
	cID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	courseProgress := builder.NewCourseProgressBuilder().WithCourseID(cID).WithUserID(uID).Build()

	testCases := []struct {
		desc               string
		userID             string
		courseID           string
		clientExpectations func(client *storemocks.MockStorer)

		expectedCourseProgress *store.CourseProgress
		expectedErr            error
	}{
		{
			desc:     "given a courseProgress is returned from store, progress is returned",
			userID:   uID,
			courseID: cID,
			clientExpectations: func(client *storemocks.MockStorer) {
				params := map[string]string{
					"courseID": cID,
					"userID":   uID,
				}

				client.EXPECT().Get(gomock.Any(), courseProgressTableName, params, gomock.Any()).SetArg(3, courseProgress)
			},

			expectedCourseProgress: &courseProgress,
		},
		{
			desc:     "given a NotFound error is returned from the store, nil map is returned",
			userID:   uID,
			courseID: cID,
			clientExpectations: func(client *storemocks.MockStorer) {
				params := map[string]string{
					"courseID": cID,
					"userID":   uID,
				}

				client.EXPECT().Get(gomock.Any(), courseProgressTableName, params, gomock.Any()).Return(store.ErrNotFound)
			},

			expectedCourseProgress: nil,
		},
		{
			desc:     "given a generic error is returned from the store, error returned",
			userID:   uID,
			courseID: cID,
			clientExpectations: func(client *storemocks.MockStorer) {
				params := map[string]string{
					"courseID": cID,
					"userID":   uID,
				}

				client.EXPECT().Get(gomock.Any(), courseProgressTableName, params, gomock.Any()).Return(fmt.Errorf("error occurred"))
			},

			expectedErr: fmt.Errorf("error occurred"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockStorer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := store.NewCourseProgressHandler(clientMock)
			ctx := context.Background()

			n, err := resolver.Get(ctx, tt.courseID, tt.userID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedCourseProgress, n)
			}
		})
	}
}

func TestCourseProgresshandlerStart(t *testing.T) {
	cID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	courseProgress := builder.NewCourseProgressBuilder().WithCourseID(cID).WithUserID(uID).Build()

	testCases := []struct {
		desc               string
		userID             string
		courseID           string
		clientExpectations func(client *storemocks.MockStorer)

		expectedCourseProgress *store.CourseProgress
		expectedErr            error
	}{
		{
			desc:     "given a progress is returned from store, started progress is returned",
			userID:   uID,
			courseID: cID,
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Put(gomock.Any(), courseProgressTableName, gomock.Any()).Return(nil)
			},

			expectedCourseProgress: &courseProgress,
		},
		{
			desc:     "given a generic error is returned from the store, error returned",
			userID:   uID,
			courseID: cID,
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Put(gomock.Any(), courseProgressTableName, gomock.Any()).Return(fmt.Errorf("error occurred"))
			},

			expectedErr: fmt.Errorf("error occurred"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockStorer(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := store.NewCourseProgressHandler(clientMock)
			ctx := context.Background()

			n, err := resolver.Start(ctx, tt.courseID, tt.userID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.NotEqual(t, tt.expectedCourseProgress.ID, n.ID)
				assert.Equal(t, tt.expectedCourseProgress.CourseID, n.CourseID)
				assert.Equal(t, tt.expectedCourseProgress.UserID, n.UserID)
				assert.Equal(t, store.STATUS_STARTED, n.State)
			}
		})
	}
}
