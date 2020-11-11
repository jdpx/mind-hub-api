package store

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
)

const (
	stepNoteTableName = "step_note"
)

// StepNoteRepositor ...
type StepNoteRepositor interface {
	GetNote(ctx context.Context, cID, uID string) (*StepNote, error)
	CreateNote(ctx context.Context, note StepNote) (*StepNote, error)
	UpdateNote(ctx context.Context, note StepNote) (*StepNote, error)
}

// StepNoteHandler ...
type StepNoteHandler struct {
	db Storer
}

// NewStepNoteHandler ...
func NewStepNoteHandler(client Storer) StepNoteHandler {
	return StepNoteHandler{
		db: client,
	}
}

// GetNote ...
func (c StepNoteHandler) GetNote(ctx context.Context, cID, uID string) (*StepNote, error) {
	p := map[string]string{
		"stepID": cID,
		"userID": uID,
	}

	res := StepNote{}
	err := c.db.Get(ctx, stepNoteTableName, p, &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

// CreateNote ...
func (c StepNoteHandler) CreateNote(ctx context.Context, note StepNote) (*StepNote, error) {
	id, _ := uuid.NewV4()
	note.ID = id.String()

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

// UpdateNote ...
func (c StepNoteHandler) UpdateNote(ctx context.Context, note StepNote) (*StepNote, error) {
	p := StepNote{
		ID:     note.ID,
		StepID: note.StepID,
		UserID: note.UserID,
		Value:  note.Value,
	}

	keys := map[string]string{
		"stepID": note.StepID,
		"userID": note.UserID,
	}

	exporession := "set info.value = :value"

	res := StepNote{}
	err := c.db.Update(ctx, stepNoteTableName, keys, exporession, p, &res)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
