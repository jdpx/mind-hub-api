package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
)

type CourseServicer interface {
	GetAll(ctx context.Context) ([]*Course, error)
	GetByID(ctx context.Context, id string) (*Course, error)
}

type CourseResolver struct {
	graphcms graphcms.Resolverer
}

// CourseResolverOption ...
type CourseResolverOption func(*CourseResolver)

// NewCourseService ...
func NewCourseService(cms graphcms.Resolverer) *CourseResolver {
	r := &CourseResolver{
		graphcms: cms,
	}

	return r
}

func (s CourseResolver) GetAll(ctx context.Context) ([]*Course, error) {
	c, err := s.graphcms.ResolveCourses(ctx)
	if err != nil {
		return nil, err
	}

	ss := CoursesFromCMS(c)

	return ss, nil
}

func (s CourseResolver) GetByID(ctx context.Context, id string) (*Course, error) {
	c, err := s.graphcms.ResolveCourse(ctx, id)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, ErrNotFound
	}

	ss := CourseFromCMS(c)

	return ss, nil
}
