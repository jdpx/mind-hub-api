package graph

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// ResolverOption ...
type ResolverOption func(*Resolver)

// CMSRequester ...
type CMSRequester interface {
	Run(ctx context.Context, req *graphcms.Request, resp interface{}) error
}

// Storer ...
type Storer interface {
	Get(ctx context.Context, key string) interface{}
	Put(ctx context.Context, key string, i interface{}) error
}

// CMSClienter ...
type CMSClienter interface {
	ResolveCourses(ctx context.Context) ([]*graphcms.Course, error)
	ResolveCourse(ctx context.Context, id string) (*graphcms.Course, error)
	ResolveCourseSessions(ctx context.Context, id string) ([]*graphcms.Session, error)
	ResolveSession(ctx context.Context, id string) (*graphcms.Session, error)
}

// Resolver ...
type Resolver struct {
	graphcms CMSClienter
	store    Storer
}

// NewResolver ...
func NewResolver(opts ...ResolverOption) *Resolver {
	r := &Resolver{}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

// WithCMSClient ...
func WithCMSClient(c CMSClienter) func(*Resolver) {
	return func(r *Resolver) {
		r.graphcms = c
	}
}

// WithStore ...
func WithStore(s Storer) func(*Resolver) {
	return func(r *Resolver) {
		r.store = s
	}
}
