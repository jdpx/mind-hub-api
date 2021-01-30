package service

import (
	"context"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/labstack/gommon/log"
)

type CourseProgressServicer interface {
	Get(ctx context.Context, cID, userID string) (*CourseProgress, error)
	Start(ctx context.Context, cID, userID string) (*CourseProgress, error)
}

type CourseProgressResolver struct {
	graphcms              graphcms.Resolverer
	courseProgressHandler store.CourseProgressRepositor
	stepProgressHandler   store.StepProgressRepositor
}

// ResolverOption ...
type ResolverOption func(*CourseProgressResolver)

// NewCourseProgressService ...
func NewCourseProgressService(c graphcms.Resolverer, cpr store.CourseProgressRepositor, spr store.StepProgressRepositor) *CourseProgressResolver {
	return &CourseProgressResolver{
		graphcms:              c,
		courseProgressHandler: cpr,
		stepProgressHandler:   spr,
	}

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
		log.Error("error getting course progress", err)
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
		return nil, fmt.Errorf("error occurred getting course progress2 %w", err)
	}

	if len(courseStepIDs) == 0 {
		return &p, nil
	}

	completedSteps, err := r.stepProgressHandler.GetCompletedByStepID(ctx, userID, courseStepIDs...)
	if err != nil {
		log.Error("error getting completed steps", err)
		return nil, fmt.Errorf("error occurred getting course progress3 %w", err)
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

	p := CourseProgress{
		// ID:          progress.ID,
		State:       sProgress.State,
		DateStarted: sProgress.DateStarted,
	}

	return &p, nil
}
