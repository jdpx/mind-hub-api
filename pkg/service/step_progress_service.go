//go:generate mockgen -source=step_progress_service.go -destination=./mocks/step_progress_service.go -package=servicemocks

package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/sirupsen/logrus"
)

type StepProgressServicer interface {
	Get(ctx context.Context, sID, userID string) (*StepProgress, error)
	Start(ctx context.Context, sID, userID string) (*StepProgress, error)
	Complete(ctx context.Context, sID, userID string) (*StepProgress, error)
}

type StepProgressService struct {
	store store.ProgressRepositor
}

// NewStepProgressService ...
func NewStepProgressService(s store.ProgressRepositor) *StepProgressService {
	return &StepProgressService{
		store: s,
	}
}

// Get ...
func (r StepProgressService) Get(ctx context.Context, sID, uID string) (*StepProgress, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.SessionIDKey: sID,
		logging.UserIDKey:    uID,
	})

	sProgress, err := r.store.Get(ctx, sID, uID)
	if err != nil {
		log.Error("error getting step progress from store", err)

		return nil, err
	}

	if sProgress == nil {
		log.Info("step progress not found in store")

		return nil, ErrNotFound
	}

	p := StepProgress{
		ID:          sProgress.ID,
		StepID:      sProgress.EntityID,
		UserID:      sProgress.UserID,
		State:       sProgress.State,
		DateStarted: sProgress.DateStarted,
	}

	if sProgress.DateCompleted != nil {
		p.DateCompleted = sProgress.DateCompleted
	}

	return &p, nil
}

// Start ...
func (r StepProgressService) Start(ctx context.Context, sID, uID string) (*StepProgress, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.SessionIDKey: sID,
		logging.UserIDKey:    uID,
	})

	sProgress, err := r.store.Start(ctx, sID, uID)
	if err != nil {
		log.Error("error starting step progress in store", err)

		return nil, err
	}

	p := StepProgress{
		ID:            sProgress.ID,
		StepID:        sProgress.EntityID,
		UserID:        sProgress.UserID,
		State:         sProgress.State,
		DateStarted:   sProgress.DateStarted,
		DateCompleted: sProgress.DateCompleted,
	}

	return &p, nil
}

// Start ...
func (r StepProgressService) Complete(ctx context.Context, sID, uID string) (*StepProgress, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.SessionIDKey: sID,
		logging.UserIDKey:    uID,
	})

	sProgress, err := r.store.Complete(ctx, sID, uID)
	if err != nil {
		log.Error("error completing step progress in store", err)

		return nil, err
	}

	p := StepProgress{
		ID:            sProgress.ID,
		StepID:        sProgress.EntityID,
		UserID:        sProgress.UserID,
		State:         sProgress.State,
		DateStarted:   sProgress.DateStarted,
		DateCompleted: sProgress.DateCompleted,
	}

	return &p, nil
}
