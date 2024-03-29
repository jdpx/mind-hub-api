//go:generate mockgen -source=step_service.go -destination=./mocks/step_service.go -package=servicemocks
package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/sirupsen/logrus"
)

type StepServicer interface {
	GetByID(ctx context.Context, id string) (*Step, error)
	CountByCourseID(ctx context.Context, id string) (int, error)
}

type StepService struct {
	graphcms graphcms.CMSResolver
}

// NewStepService ...
func NewStepService(cms graphcms.CMSResolver) *StepService {
	r := &StepService{
		graphcms: cms,
	}

	return r
}

func (s StepService) GetByID(ctx context.Context, id string) (*Step, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.StepIDKey: id,
	})

	gss, err := s.graphcms.GetStepsByID(ctx, id)
	if err != nil {
		log.Error("error getting step by id from cms", err)

		return nil, err
	}

	if gss == nil {
		log.Error("step not found in cms")

		return nil, ErrNotFound
	}

	ss := StepFromCMS(gss)

	return ss, nil
}

func (s StepService) CountByCourseID(ctx context.Context, id string) (int, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.CourseIDKey: id,
	})

	gss, err := s.graphcms.GetStepIDsByCourseID(ctx, id)
	if err != nil {
		log.Error("error getting step count by course id from cms", err)

		return 0, err
	}

	return len(gss), nil
}
