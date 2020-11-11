//go:generate mockgen -source=step_progress_handler.go -destination=./mocks/step_progress_handler.go -package=storemocks

package store

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/gommon/log"
)

const (
	stepProgressTableName = "step_progress"
)

// StepProgressRepositor ...
type StepProgressRepositor interface {
	GetStepProgress(ctx context.Context, cID, uID string) (*StepProgress, error)
	StartStep(ctx context.Context, cID, uID string) (*StepProgress, error)
}

// StepProgressHandler ...
type StepProgressHandler struct {
	db Storer
}

// NewStepProgressHandler ...
func NewStepProgressHandler(client Storer) StepProgressHandler {
	return StepProgressHandler{
		db: client,
	}
}

// GetStepProgress ...
func (c StepProgressHandler) GetStepProgress(ctx context.Context, cID, uID string) (*StepProgress, error) {
	p := map[string]string{
		"stepID": cID,
		"userID": uID,
	}

	res := StepProgress{}
	err := c.db.Get(ctx, stepProgressTableName, p, &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err

	}

	return &res, nil
}

// StartStep ...
func (c StepProgressHandler) StartStep(ctx context.Context, cID, uID string) (*StepProgress, error) {
	id, _ := uuid.NewV4()

	input := StepProgress{
		ID:          id.String(),
		StepID:      cID,
		UserID:      uID,
		DateStarted: time.Now(),
	}

	res := StepProgress{}
	err := c.db.Put(ctx, stepProgressTableName, input)
	if err != nil {
		log.Error("error getting item from store", err)
		return nil, err
	}

	return &res, nil
}
