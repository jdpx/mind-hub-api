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
	graphcms      graphcms.CMSResolver
	progressStore store.ProgressRepositor
}

// NewCourseProgressService ...
func NewCourseProgressService(c graphcms.CMSResolver, spr store.ProgressRepositor) *CourseProgressService {
	return &CourseProgressService{
		graphcms:      c,
		progressStore: spr,
	}
}

// Get ...
func (r CourseProgressService) Get(ctx context.Context, cID, uID string) (*CourseProgress, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.CourseIDKey: cID,
		logging.UserIDKey:   uID,
	})

	cProgress, err := r.progressStore.Get(ctx, cID, uID)
	if err != nil {
		log.Error("error getting course progress from store", err)

		return nil, err
	}

	if cProgress == nil {
		log.Info("course progress not found in store")

		return nil, ErrNotFound
	}

	p := CourseProgress{
		ID:            cProgress.ID,
		CourseID:      cProgress.EntityID,
		UserID:        cProgress.UserID,
		State:         cProgress.State,
		DateStarted:   cProgress.DateStarted,
		DateCompleted: cProgress.DateCompleted,
	}

	courseStepIDs, err := r.graphcms.ResolveCourseStepIDs(ctx, cID)
	if err != nil {
		log.Error("error getting course steps for course progress from store", err)

		return nil, fmt.Errorf("error occurred getting course step ids %w", err)
	}

	if len(courseStepIDs) == 0 {
		return &p, nil
	}

	completedSteps, err := r.progressStore.GetCompletedByIDs(ctx, uID, courseStepIDs...)
	if err != nil {
		log.Error("error getting completed steps for course progress from store", err)
		return nil, fmt.Errorf("error occurred getting course progress %w", err)
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

	sProgress, err := r.progressStore.Start(ctx, cID, uID)
	if err != nil {
		log.Error("error starting course progress from store", err)

		return nil, err
	}

	p := CourseProgress{
		ID:          sProgress.ID,
		CourseID:    sProgress.EntityID,
		UserID:      sProgress.UserID,
		State:       sProgress.State,
		DateStarted: sProgress.DateStarted,
	}

	return &p, nil
}
