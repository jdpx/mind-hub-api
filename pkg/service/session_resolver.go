package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
)

type SessionServicer interface {
	GetByID(ctx context.Context, id string) (*Session, error)
	GetByCourseID(ctx context.Context, id string) ([]*Session, error)
	CountByCourseID(ctx context.Context, id string) (int, error)
}

type SessionResolver struct {
	graphcms graphcms.Resolverer
}

// SessionResolverOption ...
type SessionResolverOption func(*SessionResolver)

// NewSessionService ...
func NewSessionService(cms graphcms.Resolverer) *SessionResolver {
	r := &SessionResolver{
		graphcms: cms,
	}

	return r
}

func (s SessionResolver) GetByID(ctx context.Context, id string) (*Session, error) {
	gss, err := s.graphcms.ResolveSession(ctx, id)
	if err != nil {
		return nil, err
	}

	if gss == nil {
		return nil, ErrNotFound
	}

	ss := SessionFromCMS(gss)

	return ss, nil
}

func (s SessionResolver) GetByCourseID(ctx context.Context, id string) ([]*Session, error) {
	gss, err := s.graphcms.ResolveCourseSessions(ctx, id)
	if err != nil {
		return nil, err
	}

	ss := SessionsFromCMS(gss)

	return ss, nil
}

func (s SessionResolver) CountByCourseID(ctx context.Context, id string) (int, error) {
	gss, err := s.graphcms.ResolveCourseSessions(ctx, id)
	if err != nil {
		return 0, err
	}

	return len(gss), nil
}
