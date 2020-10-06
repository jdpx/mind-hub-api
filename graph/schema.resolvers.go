package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/generated"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"
)

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Sessions(ctx context.Context) ([]*model.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Course(ctx context.Context, where model.CourseQuery) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Session(ctx context.Context, where model.SessionQuery) (*model.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
