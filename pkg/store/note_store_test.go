package store_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/jdpx/mind-hub-api/pkg/store/builder"
	storemocks "github.com/jdpx/mind-hub-api/pkg/store/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	userTableName = "user"
)

func TestNoteStoreGet(t *testing.T) {
	cID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	note := builder.NewNoteBuilder().
		WithEntityID(cID).
		WithUserID(uID).
		Build()

	testCases := []struct {
		desc               string
		userID             string
		courseID           string
		clientExpectations func(client *storemocks.MockStorer)

		expectedNote *store.Note
		expectedErr  error
	}{
		{
			desc:     "given a note is returned from store, note is returned",
			userID:   uID,
			courseID: cID,
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Get(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("NOTE#%s", cID),
					gomock.Any(),
				).SetArg(4, note)
			},

			expectedNote: &note,
		},
		{
			desc:     "given a NotFound error is returned from the store, nil map is returned",
			userID:   uID,
			courseID: cID,
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Get(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("NOTE#%s", cID),
					gomock.Any(),
				).Return(store.ErrNotFound)
			},

			expectedNote: nil,
		},
		{
			desc:     "given a generic error is returned from the store, error returned",
			userID:   uID,
			courseID: cID,
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Get(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("NOTE#%s", cID),
					gomock.Any(),
				).Return(fmt.Errorf("error occurred"))
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

			resolver := store.NewNoteStore(clientMock)
			ctx := context.Background()

			n, err := resolver.Get(ctx, tt.courseID, tt.userID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedNote, n)
			}
		})
	}
}

func TestNoteStoreCreate(t *testing.T) {
	now := time.Now()
	nID := fake.CharactersN(10)
	note := builder.NewNoteBuilder().WithID("").Build()

	createNote := builder.NewNoteBuilder().
		WithPK(fmt.Sprintf("USER#%s", note.UserID)).
		WithSK(fmt.Sprintf("NOTE#%s", note.EntityID)).
		WithEntityID(note.EntityID).
		WithUserID(note.UserID).
		WithValue(note.Value).
		WithID(nID).
		WithDateCreated(now).
		WithDateUpdated(now).
		Build()

	expectedNote := builder.NewNoteBuilder().
		WithEntityID(note.EntityID).
		WithUserID(note.UserID).
		WithValue(note.Value).
		WithID(nID).
		WithDateCreated(now).
		WithDateUpdated(now).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockStorer)

		expectedNote store.Note
		expectedErr  error
	}{
		{
			desc: "given create is successful, note is returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Put(gomock.Any(), userTableName, createNote).Return(nil)
			},

			expectedNote: expectedNote,
		},
		{
			desc: "given a generic error is returned from the store, error returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				client.EXPECT().Put(gomock.Any(), userTableName, gomock.Any()).Return(fmt.Errorf("error occurred"))
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

			gen := func() string {
				return nID
			}

			timer := func() time.Time {
				return now
			}

			resolver := store.NewNoteStore(clientMock, store.WithNoteIDGenerator(gen), store.WithNoteTimer(timer))
			ctx := context.Background()

			n, err := resolver.Create(ctx, note)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, &tt.expectedNote, n)
			}
		})
	}
}

func TestNoteStoreUpdate(t *testing.T) {
	now := time.Now()
	id := fake.CharactersN(10)
	cID := fake.CharactersN(10)
	uID := fake.CharactersN(10)
	note := builder.NewNoteBuilder().
		WithEntityID(cID).
		WithUserID(uID).
		Build()

	expectedNote := builder.NewNoteBuilder().
		WithEntityID(note.EntityID).
		WithUserID(note.UserID).
		WithValue(note.Value).
		WithID(note.ID).
		WithDateCreated(now).
		WithDateUpdated(now).
		Build()

	testCases := []struct {
		desc               string
		clientExpectations func(client *storemocks.MockStorer)

		expectedNote *store.Note
		expectedErr  error
	}{
		{
			desc: "given a note is returned from store, note is returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				upBuilder := expression.
					Set(expression.Name("id"), expression.Name("id").IfNotExists(expression.Value(id))).
					Set(expression.Name("entityID"), expression.Name("entityID").IfNotExists(expression.Value(note.EntityID))).
					Set(expression.Name("value"), expression.Value(note.Value)).
					Set(expression.Name("dateCreated"), expression.Name("dateCreated").IfNotExists(expression.Value(now))).
					Set(expression.Name("dateUpdated"), expression.Value(now))

				expr, _ := expression.NewBuilder().WithUpdate(upBuilder).Build()

				client.EXPECT().Update(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("NOTE#%s", cID),
					expr,
					gomock.Any(),
				).SetArg(5, expectedNote)
			},

			expectedNote: &expectedNote,
		},
		{
			desc: "given a generic error is returned from the store, error returned",
			clientExpectations: func(client *storemocks.MockStorer) {
				upBuilder := expression.
					Set(expression.Name("id"), expression.Name("id").IfNotExists(expression.Value(id))).
					Set(expression.Name("entityID"), expression.Name("entityID").IfNotExists(expression.Value(note.EntityID))).
					Set(expression.Name("value"), expression.Value(note.Value)).
					Set(expression.Name("dateCreated"), expression.Name("dateCreated").IfNotExists(expression.Value(now))).
					Set(expression.Name("dateUpdated"), expression.Value(now))

				expr, _ := expression.NewBuilder().WithUpdate(upBuilder).Build()

				client.EXPECT().Update(
					gomock.Any(),
					userTableName,
					fmt.Sprintf("USER#%s", uID),
					fmt.Sprintf("NOTE#%s", cID),
					expr,
					gomock.Any(),
				).Return(fmt.Errorf("error occurred"))
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

			gen := func() string {
				return id
			}

			timer := func() time.Time {
				return now
			}

			resolver := store.NewNoteStore(clientMock, store.WithNoteIDGenerator(gen), store.WithNoteTimer(timer))
			ctx := context.Background()

			n, err := resolver.Update(ctx, note)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedNote, n)
			}
		})
	}
}
