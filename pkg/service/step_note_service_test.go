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

func TestStepNoteServiceGet(t *testing.T) {
	uID := fake.CharactersN(10)
	sID := fake.CharactersN(10)
	note := builder.NewNoteBuilder().
		WithEntityID(sID).
		WithUserID(uID).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockNoteRepositor)

		expectedStepNote *service.StepNote
		expectedErr      error
	}{
		{
			desc: "given note returned from store, note returned",
			clientExpectations: func(client *storemocks.MockNoteRepositor) {
				client.EXPECT().Get(
					gomock.Any(),
					sID,
					uID,
				).Return(&note, nil)
			},

			expectedStepNote: &service.StepNote{
				ID:          note.ID,
				UserID:      note.UserID,
				StepID:      note.EntityID,
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
					sID,
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
					sID,
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

			resolver := service.NewStepNoteService(clientMock)
			ctx := context.Background()

			n, err := resolver.Get(ctx, sID, uID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedStepNote, n)
			}
		})
	}
}

func TestStepNoteServiceUpdate(t *testing.T) {
	sID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	value := fake.CharactersN(10)
	note := builder.NewNoteBuilder().
		WithUserID(uID).
		Build()

	uStepNote := store.Note{
		EntityID: sID,
		UserID:   uID,
		Value:    value,
	}

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockNoteRepositor)

		expectedStepNote *service.StepNote
		expectedErr      error
	}{
		{
			desc: "given note returned from store, map returned",
			clientExpectations: func(client *storemocks.MockNoteRepositor) {
				client.EXPECT().Update(
					gomock.Any(),
					uStepNote,
				).Return(&note, nil)
			},

			expectedStepNote: &service.StepNote{
				ID:          note.ID,
				StepID:      note.EntityID,
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
					uStepNote,
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

			resolver := service.NewStepNoteService(clientMock)
			ctx := context.Background()

			n, err := resolver.Update(ctx, sID, uID, value)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedStepNote, n)
			}
		})
	}
}
