package graph

import (
	"context"
	"fmt"

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

type coursesResponse struct {
	Courses []*model.Course `json:"courses"`
}

type courseResponse struct {
	Course *model.Course `json:"course"`
}

type sessionsResponse struct {
	Sessions []*model.Session `json:"sessions"`
}

type sessionResponse struct {
	Session *model.Session `json:"session"`
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
	res := coursesResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		fmt.Println("Error occurred", err)

		return nil, nil
	}

	return res.Courses, err
}

func (r Resolver) resolveCourse(ctx context.Context, query string) (*model.Course, error) {
	req := graphcms.NewRequest(query)
	res := courseResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		fmt.Println("Error occurred", err)

		return nil, nil
	}

	return res.Course, err
}

func (r Resolver) resolveSessions(ctx context.Context, query string) ([]*model.Session, error) {
	req := graphcms.NewRequest(query)
	res := sessionsResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		fmt.Println("Error occurred", err)

		return nil, nil
	}

	return res.Sessions, err
}

func (r Resolver) resolveSession(ctx context.Context, query string) (*model.Session, error) {
	req := graphcms.NewRequest(query)
	res := sessionResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		fmt.Println("Error occurred", err)

		return nil, nil
	}

	return res.Session, err
}
