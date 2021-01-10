//go:generate mockgen -source=course_progress_handler.go -destination=./mocks/course_progress_handler.go -package=storemocks

package store

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/gommon/log"
)

const (
	courseProgressTableName = "course_progress"
)

// CourseProgressRepositor ...
type CourseProgressRepositor interface {
	Get(ctx context.Context, cID, uID string) (*CourseProgress, error)
	Start(ctx context.Context, cID, uID string) (*CourseProgress, error)
}

// CourseProgressHandler ...
type CourseProgressHandler struct {
	db Storer
}

// NewCourseProgressHandler ...
func NewCourseProgressHandler(client Storer) CourseProgressHandler {
	return CourseProgressHandler{
		db: client,
	}
}

// Get ...
func (c CourseProgressHandler) Get(ctx context.Context, cID, uID string) (*CourseProgress, error) {
	p := map[string]string{
		"courseID": cID,
		"userID":   uID,
	}

	res := CourseProgress{}
	err := c.db.Get(ctx, courseProgressTableName, p, &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err

	}

	return &res, nil
}

// Start ...
func (c CourseProgressHandler) Start(ctx context.Context, cID, uID string) (*CourseProgress, error) {
	id, _ := uuid.NewV4()

	input := CourseProgress{
		ID:          id.String(),
		CourseID:    cID,
		UserID:      uID,
		State:       STATUS_STARTED,
		DateStarted: time.Now(),
	}

	err := c.db.Put(ctx, courseProgressTableName, input)
	if err != nil {
		log.Error("error getting item from store", err)
		return nil, err
	}

	return &input, nil
}
