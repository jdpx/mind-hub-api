package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/store"
)

type StepProgressServicer interface {
	Get(ctx context.Context, sID, userID string) (*StepProgress, error)
	Start(ctx context.Context, sID, userID string) (*StepProgress, error)
	Complete(ctx context.Context, sID, userID string) (*StepProgress, error)
}

type StepProgressService struct {
	stepProgressHandler store.StepProgressRepositor
}

// NewStepProgressService ...
func NewStepProgressService(s store.StepProgressRepositor) *StepProgressService {
	return &StepProgressService{
		stepProgressHandler: s,
	}
}

// // WithCMSClient ...
// func WithCMSClient(c graphcms.Resolverer) func(*StepProgressResolver) {
// 	return func(r *StepProgressResolver) {
// 		r.graphcms = c
// 	}
// }

// // WithStepProgressRepository ...
// func WithStepProgressRepository(s store.StepProgressRepositor) func(*StepProgressResolver) {
// 	return func(r *StepProgressResolver) {
// 		r.courseProgressHandler = s
// 	}
// }

// // WithStepProgressRepository ...
// func WithStepProgressRepository(s store.StepProgressRepositor) func(*StepProgressResolver) {
// 	return func(r *StepProgressResolver) {
// 		r.stepProgressHandler = s
// 	}
// }

// Get ...
func (r StepProgressService) Get(ctx context.Context, sID, userID string) (*StepProgress, error) {
	sProgress, err := r.stepProgressHandler.Get(ctx, sID, userID)
	if err != nil {
		return nil, err
	}

	if sProgress == nil {
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
func (r StepProgressService) Start(ctx context.Context, sID, userID string) (*StepProgress, error) {
	sProgress, err := r.stepProgressHandler.Start(ctx, sID, userID)
	if err != nil {
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
func (r StepProgressService) Complete(ctx context.Context, sID, userID string) (*StepProgress, error) {
	sProgress, err := r.stepProgressHandler.Complete(ctx, sID, userID)
	if err != nil {
		return nil, err
	}

	p := StepProgress{
		// ID:          progress.ID,
		State:       sProgress.State,
		DateStarted: sProgress.DateStarted,
	}

	return &p, nil
}
