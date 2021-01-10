package store

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
)

const (
	courseNoteTableName = "course_note"
)

// CourseNoteRepositor ...
type CourseNoteRepositor interface {
	Get(ctx context.Context, cID, uID string) (*CourseNote, error)
	Create(ctx context.Context, note CourseNote) (*CourseNote, error)
	Update(ctx context.Context, note CourseNote) (*CourseNote, error)
}

// CourseNoteHandler ...
type CourseNoteHandler struct {
	db Storer
}

// NewCourseNoteHandler ...
func NewCourseNoteHandler(client Storer) CourseNoteHandler {
	return CourseNoteHandler{
		db: client,
	}
}

// Get ...
func (c CourseNoteHandler) Get(ctx context.Context, cID, uID string) (*CourseNote, error) {
	p := map[string]string{
		"courseID": cID,
		"userID":   uID,
	}

	res := CourseNote{}
	err := c.db.Get(ctx, courseNoteTableName, p, &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

// Create ...
func (c CourseNoteHandler) Create(ctx context.Context, note CourseNote) (*CourseNote, error) {
	id, _ := uuid.NewV4()
	note.ID = id.String()

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
func (c CourseNoteHandler) Update(ctx context.Context, note CourseNote) (*CourseNote, error) {
	p := CourseNote{
		ID:       note.ID,
		CourseID: note.CourseID,
		UserID:   note.UserID,
		Value:    note.Value,
	}

	keys := map[string]string{
		"courseID": note.CourseID,
		"userID":   note.UserID,
	}

	exporession := "set info.value = :value"

	res := CourseNote{}
	err := c.db.Update(ctx, courseNoteTableName, keys, exporession, p, &res)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
