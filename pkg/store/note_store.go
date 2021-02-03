//go:generate mockgen -source=course_note_store.go -destination=./mocks/course_note_store.go -package=storemocks

package store

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/gofrs/uuid"
)

const (
	userTableName = "user"
)

// NoteRepositor ...
type NoteRepositor interface {
	Get(ctx context.Context, id, uID string) (*Note, error)
	Create(ctx context.Context, note Note) (*Note, error)
	Update(ctx context.Context, note Note) (*Note, error)
}

// NoteStore ...
type NoteStore struct {
	db Storer
}

// NewNoteStore ...
func NewNoteStore(client Storer) NoteStore {
	return NoteStore{
		db: client,
	}
}

// Get ...
func (c NoteStore) Get(ctx context.Context, id, uID string) (*Note, error) {
	res := Note{}
	err := c.db.Get(ctx, userTableName, UserPK(uID), NoteSK(id), &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

// Create ...
func (c NoteStore) Create(ctx context.Context, note Note) (*Note, error) {
	id, _ := uuid.NewV4()
	note.ID = id.String()

	note.BaseEntity = BaseEntity{
		PK: UserPK(note.UserID),
		SK: NoteSK(note.EntityID),
	}

	err := c.db.Put(ctx, userTableName, note)
	if err != nil {
		return nil, err
	}

	return &Note{
		ID:       note.ID,
		EntityID: note.EntityID,
		UserID:   note.UserID,
		Value:    note.Value,
	}, nil
}

// Update ...
func (c NoteStore) Update(ctx context.Context, note Note) (*Note, error) {
	p := Note{
		ID:       note.ID,
		EntityID: note.EntityID,
		UserID:   note.UserID,
		Value:    note.Value,
	}

	upBuilder := expression.Set(expression.Name("Value"), expression.Value(note.Value))

	expr, err := expression.NewBuilder().WithUpdate(upBuilder).Build()
	if err != nil {
		return nil, err
	}

	res := Note{}
	err = c.db.Update(ctx, userTableName, UserPK(note.UserID), NoteSK(note.EntityID), expr, &res)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
