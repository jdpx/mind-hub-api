package service

import (
	"context"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/labstack/gommon/log"
)

type CourseProgressRepositor interface {
	Get(ctx context.Context, cID, userID string) (*CourseProgress, error)
	Start(ctx context.Context, cID, userID string) (*CourseProgress, error) {
}

type CourseProgressResolver struct {
	graphcms              graphcms.Resolverer
	courseProgressHandler store.CourseProgressRepositor
	stepProgressHandler   store.StepProgressRepositor
}

// ResolverOption ...
type ResolverOption func(*CourseProgressResolver)

// New ...
func New(opts ...ResolverOption) *CourseProgressResolver {
	r := &CourseProgressResolver{}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

// WithCMSClient ...
func WithCMSClient(c graphcms.Resolverer) func(*CourseProgressResolver) {
	return func(r *CourseProgressResolver) {
		r.graphcms = c
	}
}

// WithCourseProgressRepository ...
func WithCourseProgressRepository(s store.CourseProgressRepositor) func(*CourseProgressResolver) {
	return func(r *CourseProgressResolver) {
		r.courseProgressHandler = s
	}
}

// WithStepProgressRepository ...
func WithStepProgressRepository(s store.StepProgressRepositor) func(*CourseProgressResolver) {
	return func(r *CourseProgressResolver) {
		r.stepProgressHandler = s
	}
}

// Get ...
func (r CourseProgressResolver) Get(ctx context.Context, cID, userID string) (*CourseProgress, error) {
	sProgress, err := r.courseProgressHandler.Get(ctx, cID, userID)
	if err != nil {
		return nil, err
	}

	if sProgress == nil {
		return nil, ErrNotFound
	}

	p := CourseProgress{
		// ID:          progress.ID,
		State:       sProgress.State,
		DateStarted: sProgress.DateStarted,
	}

	courseStepIDs, err := r.graphcms.ResolveCourseStepIDs(ctx, cID)
	if err != nil {
		log.Error("error getting course steps", err)
		return nil, fmt.Errorf("error occurred getting course progress %w", err)
	}

	if len(courseStepIDs) == 0 {
		return &p, nil
	}

	completedSteps, err := r.stepProgressHandler.GetCompletedByStepID(ctx, userID, courseStepIDs...)
	if err != nil {
		log.Error("error getting completed steps", err)
		return nil, fmt.Errorf("error occurred getting course progress %w", err)
	}

	p.CompletedSteps = len(completedSteps)

	return &p, nil
}

// Start ...
func (r CourseProgressResolver) Start(ctx context.Context, cID, userID string) (*CourseProgress, error) {
	sProgress, err := r.courseProgressHandler.Start(ctx, cID, userID)
	if err != nil {
		return nil, err
	}

	

}