package graph

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// ResolverOption ...
type ResolverOption func(*Resolver)

// Requester ...
type Requester interface {
	Run(ctx context.Context, req *graphcms.Request, resp interface{}) error
}

// Resolver ...
type Resolver struct {
	client Requester
}

// CoursesResponse ...
type CoursesResponse struct {
	Courses []*model.Course `json:"courses"`
}

// CourseResponse ...
type CourseResponse struct {
	Course *model.Course `json:"course"`
}

// SessionsResponse ...
type SessionsResponse struct {
	Sessions []*model.Session `json:"sessions"`
}

// SessionResponse ...
type SessionResponse struct {
	Session *model.Session `json:"session"`
}

// StepResponse ...
type StepResponse struct {
	Step *model.Step `json:"step"`
}

// NewResolver ...
func NewResolver(opts ...ResolverOption) *Resolver {
	r := &Resolver{}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

// WithClient ...
func WithClient(c Requester) func(*Resolver) {
	return func(r *Resolver) {
		r.client = c
	}
}

func (r Resolver) resolveCourses(ctx context.Context, query string) ([]*model.Course, error) {
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
	req := graphcms.NewQueryRequest(query.RawQuery, query.Variables)
	res := StepResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		fmt.Println("Error occurred", err)

		return nil, err
	}

	return res.Step, err
}
