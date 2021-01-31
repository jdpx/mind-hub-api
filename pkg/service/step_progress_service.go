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
	store store.StepProgressRepositor
}

// NewStepProgressService ...
func NewStepProgressService(s store.StepProgressRepositor) *StepProgressService {
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
		log.Error("error getting session progress from store", err)

		return nil, err
	}

	if sProgress == nil {
		log.Info("session progress not found in store")

		return nil, ErrNotFound
	}

	p := StepProgress{
		// ID:          progress.ID,
		State: sProgress.State,
	}

	if sProgress.DateStarted != nil {
		p.DateStarted = sProgress.DateStarted
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
		log.Error("error starting session progress in store", err)

		return nil, err
	}

	p := StepProgress{
		// ID:          progress.ID,
		State:       sProgress.State,
		DateStarted: sProgress.DateStarted,
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
		log.Error("error completing session progress in store", err)

		return nil, err
	}

	p := StepProgress{
		// ID:          progress.ID,
		State:       sProgress.State,
		DateStarted: sProgress.DateStarted,
	}

	return &p, nil
}
