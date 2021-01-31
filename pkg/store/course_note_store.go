package store

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/gofrs/uuid"
)

const (
	courseNoteTableName = "user"
)

// CourseNoteRepositor ...
type CourseNoteRepositor interface {
	Get(ctx context.Context, cID, uID string) (*CourseNote, error)
	Create(ctx context.Context, note CourseNote) (*CourseNote, error)
	Update(ctx context.Context, note CourseNote) (*CourseNote, error)
}

// CourseNoteStore ...
type CourseNoteStore struct {
	db Storer
}

// NewCourseNoteStore ...
func NewCourseNoteStore(client Storer) CourseNoteStore {
	return CourseNoteStore{
		db: client,
	}
}

// Get ...
func (c CourseNoteStore) Get(ctx context.Context, cID, uID string) (*CourseNote, error) {
	res := CourseNote{}
	err := c.db.Get(ctx, userTableName, UserPK(uID), NoteSK(cID), &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

// Create ...
func (c CourseNoteStore) Create(ctx context.Context, note CourseNote) (*CourseNote, error) {
	id, _ := uuid.NewV4()
	note.ID = id.String()

	note.BaseEntity = BaseEntity{
		PK: UserPK(note.UserID),
		SK: NoteSK(note.CourseID),
	}

	err := c.db.Put(ctx, courseNoteTableName, note)
	if err != nil {
		return nil, err
	}

	return &CourseNote{
		ID:       note.ID,
		CourseID: note.CourseID,
		UserID:   note.UserID,
		Value:    note.Value,
	}, nil
}

// Update ...
func (c CourseNoteStore) Update(ctx context.Context, note CourseNote) (*CourseNote, error) {
	p := CourseNote{
		ID:       note.ID,
		CourseID: note.CourseID,
		UserID:   note.UserID,
		Value:    note.Value,
	}

	upBuilder := expression.Set(expression.Name("Value"), expression.Value(note.Value))

	expr, err := expression.NewBuilder().WithUpdate(upBuilder).Build()
	if err != nil {
		return nil, err
	}

	res := CourseNote{}
	err = c.db.Update(ctx, courseNoteTableName, UserPK(note.UserID), NoteSK(note.CourseID), expr, &res)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
