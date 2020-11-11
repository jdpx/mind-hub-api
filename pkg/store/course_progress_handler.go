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
	GetCourseProgress(ctx context.Context, cID, uID string) (*CourseProgress, error)
	StartCourse(ctx context.Context, cID, uID string) (*CourseProgress, error)
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

// GetCourseProgress ...
func (c CourseProgressHandler) GetCourseProgress(ctx context.Context, cID, uID string) (*CourseProgress, error) {
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

// StartCourse ...
func (c CourseProgressHandler) StartCourse(ctx context.Context, cID, uID string) (*CourseProgress, error) {
	id, _ := uuid.NewV4()

	input := CourseProgress{
		ID:          id.String(),
		CourseID:    cID,
		UserID:      uID,
		DateStarted: time.Now(),
	}

	res := CourseProgress{}
	err := c.db.Put(ctx, courseProgressTableName, input)
	if err != nil {
		log.Error("error getting item from store", err)
		return nil, err
	}

	return &res, nil
}
