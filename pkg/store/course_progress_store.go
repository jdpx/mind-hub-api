//go:generate mockgen -source=course_progress_store.go -destination=./mocks/course_progress_store.go -package=storemocks

package store

import (
	"context"
	"errors"
	"time"

	"github.com/labstack/gommon/log"
)

const (
	userTableName = "user"
)

// CourseProgressRepositor ...
type CourseProgressRepositor interface {
	Get(ctx context.Context, cID, uID string) (*CourseProgress, error)
	Start(ctx context.Context, cID, uID string) (*CourseProgress, error)
}

// CourseProgressStore ...
type CourseProgressStore struct {
	db Storer
}

// NewCourseProgressStore ...
func NewCourseProgressStore(client Storer) CourseProgressStore {
	return CourseProgressStore{
		db: client,
	}
}

// Get ...
func (c CourseProgressStore) Get(ctx context.Context, cID, uID string) (*CourseProgress, error) {
	res := CourseProgress{}
	err := c.db.Get(ctx, userTableName, UserPK(uID), ProgressSK(cID), &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err

	}

	return &res, nil
}

// Start ...
func (c CourseProgressStore) Start(ctx context.Context, cID, uID string) (*CourseProgress, error) {
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

	err := c.db.Put(ctx, userTableName, input)
	if err != nil {
		log.Error("error getting item from store", err)
		return nil, err
	}

	return &input, nil
}
