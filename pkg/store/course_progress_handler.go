//go:generate mockgen -source=course_progress_handler.go -destination=./mocks/course_progress_handler.go -package=storemocks

package store

import (
	"context"
	"errors"
	"time"

	"github.com/labstack/gommon/log"
)

const (
	courseProgressTableName = "user"
)

// CourseProgressRepositor ...
type CourseProgressRepositor interface {
	Get(ctx context.Context, cID, uID string) (*CourseProgress, error)
	Start(ctx context.Context, cID, uID string) (*CourseProgress, error)
}

// CourseProgressHandler ...
type CourseProgressHandler struct {
	db StorerV2
}

// NewCourseProgressHandler ...
func NewCourseProgressHandler(client StorerV2) CourseProgressHandler {
	return CourseProgressHandler{
		db: client,
	}
}

// Get ...
func (c CourseProgressHandler) Get(ctx context.Context, cID, uID string) (*CourseProgress, error) {
	res := CourseProgress{}
	err := c.db.Get(ctx, courseProgressTableName, UserPK(uID), ProgressSK(cID), &res)
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
	input := CourseProgress{
		BaseEntity: BaseEntity{
			PK: UserPK(uID),
			SK: ProgressSK(cID),
		},

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
