package store

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/gofrs/uuid"
)

const (
	stepNoteTableName = "user"
)

// StepNoteRepositor ...
type StepNoteRepositor interface {
	Get(ctx context.Context, cID, uID string) (*StepNote, error)
	Create(ctx context.Context, note StepNote) (*StepNote, error)
	Update(ctx context.Context, note StepNote) (*StepNote, error)
}

// StepNoteHandler ...
type StepNoteHandler struct {
	db StorerV2
}

// NewStepNoteHandler ...
func NewStepNoteHandler(client StorerV2) StepNoteHandler {
	return StepNoteHandler{
		db: client,
	}
}

// Get ...
func (c StepNoteHandler) Get(ctx context.Context, cID, uID string) (*StepNote, error) {
	res := StepNote{}
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
func (c StepNoteHandler) Create(ctx context.Context, note StepNote) (*StepNote, error) {
	id, _ := uuid.NewV4()
	note.ID = id.String()

	note.BaseEntity = BaseEntity{
		PK: UserPK(note.UserID),
		SK: NoteSK(note.StepID),
	}

	err := c.db.Put(ctx, stepNoteTableName, note)
	if err != nil {
		return nil, err
	}

	return &StepNote{
		ID:     note.ID,
		StepID: note.StepID,
		UserID: note.UserID,
		Value:  note.Value,
	}, nil
}

// Update ...
func (c StepNoteHandler) Update(ctx context.Context, note StepNote) (*StepNote, error) {
	p := StepNote{
		ID:     note.ID,
		StepID: note.StepID,
		UserID: note.UserID,
		Value:  note.Value,
	}

	upBuilder := expression.Set(expression.Name("Value"), expression.Value(note.Value))

	expr, err := expression.NewBuilder().WithUpdate(upBuilder).Build()
	if err != nil {
		return nil, err
	}

	res := StepNote{}
	err = c.db.Update(ctx, stepNoteTableName, UserPK(note.UserID), NoteSK(note.StepID), expr, &res)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
