//go:generate mockgen -source=step_progress_handler.go -destination=./mocks/step_progress_handler.go -package=storemocks

package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/gommon/log"
)

const (
	stepProgressTableName = "step_progress"
)

// StepProgressRepositor ...
type StepProgressRepositor interface {
	GetStepProgress(ctx context.Context, sID, uID string) (*StepProgress, error)
	GetCompletedProgressByStepID(ctx context.Context, uID string, ids ...string) ([]*StepProgress, error)
	StartStep(ctx context.Context, sID, uID string) (*StepProgress, error)
	CompleteStep(ctx context.Context, sID, uID string) (*StepProgress, error)
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
func (c StepProgressHandler) GetStepProgress(ctx context.Context, sID, uID string) (*StepProgress, error) {
	p := map[string]string{
		"stepID": sID,
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

// GetStepProgress ...
func (c StepProgressHandler) GetCompletedProgressByStepID(ctx context.Context, uID string, ids ...string) ([]*StepProgress, error) {
	res := []*StepProgress{}

	keys := []map[string]string{}

	for _, id := range ids {
		m := map[string]string{
			"userID": uID,
			"stepID": id,
		}

		keys = append(keys, m)
	}

	err := c.db.Query(ctx, stepProgressTableName, keys, &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err
	}

	fP := []*StepProgress{}

	for _, p := range res {
		if p.State == STATUS_COMPLETED {
			fP = append(fP, p)
		}
	}

	return fP, nil
}

// StartStep ...
func (c StepProgressHandler) StartStep(ctx context.Context, sID, uID string) (*StepProgress, error) {
	id, _ := uuid.NewV4()

	now := time.Now()
	input := StepProgress{
		ID:          id.String(),
		StepID:      sID,
		UserID:      uID,
		State:       STATUS_STARTED,
		DateStarted: &now,
	}

	res := StepProgress{}
	err := c.db.Put(ctx, stepProgressTableName, input)
	if err != nil {
		log.Error(fmt.Sprintf("error completing Step %s in store", sID), err)
		return nil, err
	}

	return &res, nil
}

// StartStep ...
func (c StepProgressHandler) CompleteStep(ctx context.Context, sID, uID string) (*StepProgress, error) {
	now := time.Now()

	input := map[string]interface{}{
		":dateCompleted": now,
		":progressState": STATUS_COMPLETED,
	}

	keys := map[string]string{
		"stepID": sID,
		"userID": uID,
	}

	expression := "SET dateCompleted = :dateCompleted, progressState = :progressState"

	res := StepProgress{}
	err := c.db.Update(ctx, stepProgressTableName, keys, expression, input, &res)
	if err != nil {
		log.Error(fmt.Sprintf("error completing Step %s in store", sID), err)
		return nil, err
	}

	return &res, nil
}