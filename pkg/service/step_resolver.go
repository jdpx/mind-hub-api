package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
)

type StepServicer interface {
	GetByID(ctx context.Context, id string) (*Step, error)
	CountByCourseID(ctx context.Context, id string) (int, error)
}

type StepResolver struct {
	graphcms graphcms.Resolverer
}

// StepResolverOption ...
type StepResolverOption func(*StepResolver)

// NewStepService ...
func NewStepService(cms graphcms.Resolverer) *StepResolver {
	r := &StepResolver{
		graphcms: cms,
	}

	return r
}

func (s StepResolver) GetByID(ctx context.Context, id string) (*Step, error) {
	gss, err := s.graphcms.ResolveStep(ctx, id)
	if err != nil {
		return nil, err
	}

	if gss == nil {
		return nil, ErrNotFound
	}

	ss := StepFromCMS(gss)

	return ss, nil
}

func (s StepResolver) CountByCourseID(ctx context.Context, id string) (int, error) {
	gss, err := s.graphcms.ResolveCourseStepIDs(ctx, id)
	if err != nil {
		return 0, err
	}

	return len(gss), nil
}
