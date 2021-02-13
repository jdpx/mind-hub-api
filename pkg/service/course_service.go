//go:generate mockgen -source=course_service.go -destination=./mocks/course_service.go -package=servicemocks

package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/sirupsen/logrus"
)

type CourseServicer interface {
	GetAll(ctx context.Context) ([]*Course, error)
	GetByID(ctx context.Context, id string) (*Course, error)
}

type CourseService struct {
	graphcms graphcms.CMSResolver
}

// NewCourseService ...
func NewCourseService(cms graphcms.CMSResolver) *CourseService {
	r := &CourseService{
		graphcms: cms,
	}

	return r
}

func (s CourseService) GetAll(ctx context.Context) ([]*Course, error) {
	log := logging.NewFromResolver(ctx)

	c, err := s.graphcms.GetCourses(ctx)
	if err != nil {
		log.Error("error getting all courses from cms", err)

		return nil, err
	}

	ss := CoursesFromCMS(c)

	return ss, nil
}

func (s CourseService) GetByID(ctx context.Context, id string) (*Course, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.CourseIDKey: id,
	})

	c, err := s.graphcms.GetCourseByID(ctx, id)
	if err != nil {
		log.Error("error getting course by id from cms", err)

		return nil, err
	}

	if c == nil {
		log.Error("course not found in cms")

		return nil, ErrNotFound
	}

	ss := CourseFromCMS(c)

	return ss, nil
}
