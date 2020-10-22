package graph

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"
	"github.com/jdpx/mind-hub-api/pkg/logging"
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
	Put(ctx context.Context, i interface{}) error
}

// Resolver ...
type Resolver struct {
	client CMSRequester
	store  Storer
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
func WithCMSClient(c CMSRequester) func(*Resolver) {
	return func(r *Resolver) {
		r.client = c
	}
}

// WithStore ...
func WithStore(s Storer) func(*Resolver) {
	return func(r *Resolver) {
		r.store = s
	}
}

func (r Resolver) resolveCourses(ctx context.Context, query string) ([]*model.Course, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.QueryKey, query)
	log.Info("Resolving Courses")

	req := graphcms.NewRequest(query)
	res := CoursesResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		fmt.Println("Error occurred", err)

		return nil, err
	}

	return res.Courses, err
}

func (r Resolver) resolveCourse(ctx context.Context, query *graphql.OperationContext) (*model.Course, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.QueryKey, query)
	log.Info("Resolving Course")

	req := graphcms.NewQueryRequest(query.RawQuery, query.Variables)
	res := CourseResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		fmt.Println("Error occurred", err)

		return nil, err
	}

	return res.Course, err
}

func (r Resolver) resolveSessions(ctx context.Context, query string) ([]*model.Session, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.QueryKey, query)
	log.Info("Resolving Sessions")

	req := graphcms.NewRequest(query)
	res := SessionsResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		fmt.Println("Error occurred", err)

		return nil, err
	}

	return res.Sessions, err
}

func (r Resolver) resolveSession(ctx context.Context, query *graphql.OperationContext) (*model.Session, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.QueryKey, query)
	log.Info("Resolving Session")

	req := graphcms.NewQueryRequest(query.RawQuery, query.Variables)
	res := SessionResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		fmt.Println("Error occurred", err)

		return nil, err
	}

	return res.Session, err
}

func (r Resolver) resolveStep(ctx context.Context, query *graphql.OperationContext) (*model.Step, error) {
	log := logging.NewFromResolver(ctx).WithField(logging.QueryKey, query)
	log.Info("Resolving Step")

	req := graphcms.NewQueryRequest(query.RawQuery, query.Variables)
	res := StepResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		fmt.Println("Error occurred", err)

		return nil, err
	}

	return res.Step, err
}
