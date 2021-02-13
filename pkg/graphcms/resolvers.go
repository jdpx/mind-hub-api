//go:generate mockgen -source=resolvers.go -destination=./mocks/resolvers.go -package=graphcmsmocks

package graphcms

import (
	"context"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/logging"
)

// CMSResolver ...
type CMSResolver interface {
	GetCourses(ctx context.Context) ([]*Course, error)
	GetCourseByID(ctx context.Context, id string) (*Course, error)
	GetSessionsByCourseID(ctx context.Context, id string) ([]*Session, error)
	GetSessionByID(ctx context.Context, id string) (*Session, error)
	GetStepIDsByCourseID(ctx context.Context, id string) ([]string, error)
	GetStepsByID(ctx context.Context, id string) (*Step, error)
}

// Resolver defines the resolver used to retrieve data from GraphCMS
type Resolver struct {
	client Requester
}

// NewResolver initialises a new Resolver
func NewResolver(client Requester) *Resolver {
	return &Resolver{
		client: client,
	}
}

// GetCourses retrieves Courses from GraphCMS
func (r Resolver) GetCourses(ctx context.Context) ([]*Course, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Resolving Courses")

	req := NewRequest(ctx, getAllCoursesQuery)
	res := coursesResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error occurred getting GraphCMS Courses")
	}

	return res.Courses, nil
}

// GetCourseByID retrieves Courses from GraphCMS based on the Course ID
func (r Resolver) GetCourseByID(ctx context.Context, id string) (*Course, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.CourseIDKey, id)
	log.Info(fmt.Sprintf("Resolving GraphCMS Course %s", id))

	req := NewRequest(ctx, getCourseByID)
	req.Var("id", id)
	res := courseResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error occurred getting GraphCMS Course")
	}

	return res.Course, err
}

// GetSessionsByCourseID retrieves Sessions from GraphCMS based on the Course ID
func (r Resolver) GetSessionsByCourseID(ctx context.Context, id string) ([]*Session, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.CourseIDKey, id)
	log.Info(fmt.Sprintf("Resolving GraphCMS Course %s Sessions", id))

	req := NewRequest(ctx, getSessionsByCourseID)
	req.Var("id", id)
	res := sessionsResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error occurred getting GraphCMS Course Sessions")
	}

	return res.Sessions, err
}

// GetSessionByID retrieves a Session from GraphCMS based on the Session ID
func (r Resolver) GetSessionByID(ctx context.Context, id string) (*Session, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.SessionIDKey, id)
	log.Info(fmt.Sprintf("Resolving GraphCMS Session %s", id))

	req := NewRequest(ctx, getSessionByID)
	req.Var("id", id)
	res := sessionResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error occurred getting GraphCMS Session")
	}

	return res.Session, err
}

// GetStepIDsByCourseID retreives a Session from GraphCMS based on the Session ID
func (r Resolver) GetStepIDsByCourseID(ctx context.Context, id string) ([]string, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.SessionIDKey, id)
	log.Info(fmt.Sprintf("Resolving GraphCMS Step IDs for Course %s", id))

	req := NewRequest(ctx, getSessionsByCourseID)
	req.Var("id", id)
	res := sessionsResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error occurred getting GraphCMS Session By Course ID")
	}

	ids := []string{}

	for _, session := range res.Sessions {
		for _, step := range session.Steps {
			ids = append(ids, step.ID)
		}
	}

	return ids, err
}

// GetStepsByID retreives Steps from GraphCMS based on the Step ID
func (r Resolver) GetStepsByID(ctx context.Context, id string) (*Step, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.StepIDKey, id)
	log.Info(fmt.Sprintf("Resolving GraphCMS Step %s", id))

	req := NewRequest(ctx, getStepByID)
	req.Var("id", id)
	res := stepResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error occurred getting GraphCMS Step")
	}

	return res.Step, err
}
