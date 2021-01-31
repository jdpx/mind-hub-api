//go:generate mockgen -source=course_progress_service.go -destination=./mocks/course_progress_service.go -package=servicemocks

package service

import (
	"context"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/sirupsen/logrus"
)

type CourseProgressServicer interface {
	Get(ctx context.Context, cID, userID string) (*CourseProgress, error)
	Start(ctx context.Context, cID, userID string) (*CourseProgress, error)
}

type CourseProgressService struct {
	graphcms            graphcms.Resolverer
	courseProgressStore store.CourseProgressRepositor
	stepProgressStore   store.StepProgressRepositor
}

// NewCourseProgressService ...
func NewCourseProgressService(c graphcms.Resolverer, cpr store.CourseProgressRepositor, spr store.StepProgressRepositor) *CourseProgressService {
	return &CourseProgressService{
		graphcms:            c,
		courseProgressStore: cpr,
		stepProgressStore:   spr,
	}
}

// Get ...
func (r CourseProgressService) Get(ctx context.Context, cID, uID string) (*CourseProgress, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.CourseIDKey: cID,
		logging.UserIDKey:   uID,
	})

	cProgress, err := r.courseProgressStore.Get(ctx, cID, uID)
	if err != nil {
		log.Error("error getting course progress from store", err)

		return nil, err
	}

	if cProgress == nil {
		log.Info("course progress not found in store")

		return nil, ErrNotFound
	}

	p := CourseProgress{
		// ID:          progress.ID,
		State:       cProgress.State,
		DateStarted: cProgress.DateStarted,
	}

	courseStepIDs, err := r.graphcms.ResolveCourseStepIDs(ctx, cID)
	if err != nil {
		log.Error("error getting course steps for course progress from store", err)

		return nil, fmt.Errorf("error occurred getting course progress %w", err)
	}

	if len(courseStepIDs) == 0 {
		return &p, nil
	}

	completedSteps, err := r.stepProgressStore.GetCompletedByStepID(ctx, uID, courseStepIDs...)
	if err != nil {
		log.Error("error getting completed steps for course progress from store", err)
		return nil, fmt.Errorf("error occurred getting course progress3 %w", err)
	}

	p.CompletedSteps = len(completedSteps)

	return &p, nil
}

// Start ...
func (r CourseProgressService) Start(ctx context.Context, cID, uID string) (*CourseProgress, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.CourseIDKey: cID,
		logging.UserIDKey:   uID,
	})

	sProgress, err := r.courseProgressStore.Start(ctx, cID, uID)
	if err != nil {
		log.Error("error starting course progress from store", err)

		return nil, err
	}

	p := CourseProgress{
		// ID:          progress.ID,
		State:       sProgress.State,
		DateStarted: sProgress.DateStarted,
	}

	return &p, nil
}
