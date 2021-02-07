//go:generate mockgen -source=session_service.go -destination=./mocks/session_service.go -package=servicemocks

package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/sirupsen/logrus"
)

type SessionServicer interface {
	GetByID(ctx context.Context, id string) (*Session, error)
	GetByCourseID(ctx context.Context, id string) ([]*Session, error)
	CountByCourseID(ctx context.Context, id string) (int, error)
}

type SessionService struct {
	graphcms graphcms.Resolverer
}

// NewSessionService ...
func NewSessionService(cms graphcms.Resolverer) *SessionService {
	r := &SessionService{
		graphcms: cms,
	}

	return r
}

func (s SessionService) GetByID(ctx context.Context, id string) (*Session, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.SessionIDKey: id,
	})

	gss, err := s.graphcms.ResolveSession(ctx, id)
	if err != nil {
		log.Error("error getting session by id from cms", err)

		return nil, err
	}

	if gss == nil {
		log.Error("session not found in cms")

		return nil, ErrNotFound
	}

	ss := SessionFromCMS(gss)

	return ss, nil
}

func (s SessionService) GetByCourseID(ctx context.Context, id string) ([]*Session, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.CourseIDKey: id,
	})

	gss, err := s.graphcms.ResolveCourseSessions(ctx, id)
	if err != nil {
		log.Error("error getting session by course id from cms", err)

		return nil, err
	}

	ss := SessionsFromCMS(gss)

	return ss, nil
}

func (s SessionService) CountByCourseID(ctx context.Context, id string) (int, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.CourseIDKey: id,
	})

	gss, err := s.graphcms.ResolveCourseSessions(ctx, id)
	if err != nil {
		log.Error("error getting session count by course id from cms", err)

		return 0, err
	}

	return len(gss), nil
}
