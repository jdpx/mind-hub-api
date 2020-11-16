//go:generate mockgen -source=resolvers.go -destination=./mocks/resolvers.go -package=graphcmsmocks

package graphcms

import (
	"context"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/logging"
)

// Resolver defines the resolver used to retreive data from GraphCMS
type Resolver struct {
	client Requester
}

// Resolverer ...
type Resolverer interface {
	ResolveCourses(ctx context.Context) ([]*Course, error)
	ResolveCourse(ctx context.Context, id string) (*Course, error)
	ResolveCourseSessions(ctx context.Context, id string) ([]*Session, error)
	ResolveSession(ctx context.Context, id string) (*Session, error)
	ResolveCourseStepIDs(ctx context.Context, id string) ([]string, error)
}

// NewResolver initialises a new Resolver
func NewResolver(client Requester) *Resolver {
	return &Resolver{
		client: client,
	}
}

// ResolveCourses retreives Courses from GraphCMS
func (r Resolver) ResolveCourses(ctx context.Context) ([]*Course, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Resolving Courses")

	req := NewRequest(getAllCoursesQuery)
	res := coursesResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error occurred getting GraphCMS Courses")
	}

	return res.Courses, nil
}

// ResolveCourse retreives Courses from GraphCMS based on the Course ID
func (r Resolver) ResolveCourse(ctx context.Context, id string) (*Course, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.CourseIDKey, id)
	log.Info(fmt.Sprintf("Resolving Graphcms Course %s", id))

	req := NewRequest(getCourseByID)
	req.Var("id", id)
	res := courseResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error occurred getting GraphCMS Course")
	}

	return res.Course, err
}

// ResolveCourseSessions retreives Sessions from GraphCMS based on the Course ID
func (r Resolver) ResolveCourseSessions(ctx context.Context, id string) ([]*Session, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.CourseIDKey, id)
	log.Info(fmt.Sprintf("Resolving Graphcms Course %s Sessions", id))

	req := NewRequest(getSessionsByCourseID)
	req.Var("id", id)
	res := sessionsResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error occurred getting GraphCMS Course Sessions")
	}

	return res.Sessions, err
}

// ResolveSession retreives a Session from GraphCMS based on the Session ID
func (r Resolver) ResolveSession(ctx context.Context, id string) (*Session, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.SessionIDKey, id)
	log.Info(fmt.Sprintf("Resolving Graphcms Session %s", id))

	req := NewRequest(getSessionByID)
	req.Var("id", id)
	res := sessionResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error occurred getting GraphCMS Session")
	}

	return res.Session, err
}

// ResolveCourseStepIDs retreives a Session from GraphCMS based on the Session ID
func (r Resolver) ResolveCourseStepIDs(ctx context.Context, id string) ([]string, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.SessionIDKey, id)
	log.Info(fmt.Sprintf("Resolving Graphcms Step IDs for Course %s", id))

	req := NewRequest(getSessionsByCourseID)
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
