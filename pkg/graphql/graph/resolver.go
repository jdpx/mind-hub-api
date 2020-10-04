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

type courseResponse struct {
	Courses []*model.Course `json:"courses"`
}

func (r Resolver) resolveCourses(ctx context.Context, query string) ([]*model.Course, error) {
	req := graphcms.NewRequest(query)
	res := courseResponse{}

	err := r.client.Run(ctx, req, &res)
	if err != nil {
		fmt.Println("Error occurred", err)
	}

	return res.Courses, err
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
