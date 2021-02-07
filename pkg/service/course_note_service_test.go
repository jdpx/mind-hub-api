package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/service"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/jdpx/mind-hub-api/pkg/store/builder"
	storemocks "github.com/jdpx/mind-hub-api/pkg/store/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCourseNoteServiceGet(t *testing.T) {
	uID := fake.CharactersN(10)
	cID := fake.CharactersN(10)
	note := builder.NewNoteBuilder().
		WithEntityID(cID).
		WithUserID(uID).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockNoteRepositor)

		expectedCourseNote *service.CourseNote
		expectedErr        error
	}{
		{
			desc: "given note returned from store, note returned",
			clientExpectations: func(client *storemocks.MockNoteRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					cID,
					uID,
				).Return(&note, nil)
			},

			expectedCourseNote: &service.CourseNote{
				ID:          note.ID,
				UserID:      note.UserID,
				CourseID:    note.EntityID,
				Value:       note.Value,
				DateCreated: note.DateCreated,
				DateUpdated: note.DateUpdated,
			},
		},
		{
			desc: "given an error is returned from the store, err returned",
			clientExpectations: func(client *storemocks.MockNoteRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					cID,
					uID,
				).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
		{
			desc: "given the returned entity is nil, ErrNotFound returned",
			clientExpectations: func(client *storemocks.MockNoteRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					cID,
					uID,
				).Return(nil, nil)
			},

			expectedErr: service.ErrNotFound,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockNoteRepositor(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := service.NewCourseNoteService(clientMock)
			ctx := context.Background()

			n, err := resolver.Get(ctx, cID, uID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedCourseNote, n)
			}
		})
	}
}

func TestCourseNoteServiceUpdate(t *testing.T) {
	cID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	value := fake.CharactersN(10)
	note := builder.NewNoteBuilder().
		WithUserID(uID).
		Build()

	uCourseNote := store.Note{
		EntityID: cID,
		UserID:   uID,
		Value:    value,
	}

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockNoteRepositor)

		expectedCourseNote *service.CourseNote
		expectedErr        error
	}{
		{
			desc: "given note returned from store, map returned",
			clientExpectations: func(client *storemocks.MockNoteRepositor) {
				client.EXPECT().Update(
					gomock.Any(),
					uCourseNote,
				).Return(&note, nil)
			},

			expectedCourseNote: &service.CourseNote{
				ID:          note.ID,
				CourseID:    note.EntityID,
				UserID:      note.UserID,
				Value:       note.Value,
				DateCreated: note.DateCreated,
				DateUpdated: note.DateUpdated,
			},
		},
		{
			desc: "given an error is returned from the store, err returned",
			clientExpectations: func(client *storemocks.MockNoteRepositor) {
				client.EXPECT().Update(
					gomock.Any(),
					uCourseNote,
				).Return(nil, fmt.Errorf("something went wrong"))
			},

			expectedErr: fmt.Errorf("something went wrong"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			clientMock := storemocks.NewMockNoteRepositor(ctrl)

			if tt.clientExpectations != nil {
				tt.clientExpectations(clientMock)
			}

			resolver := service.NewCourseNoteService(clientMock)
			ctx := context.Background()

			n, err := resolver.Update(ctx, cID, uID, value)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedCourseNote, n)
			}
		})
	}
}
