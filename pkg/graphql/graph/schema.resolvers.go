package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/generated"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"
)

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	preloads := graphql.GetOperationContext(ctx)

	return r.resolveCourses(ctx, preloads.RawQuery)
}

func (r *queryResolver) Sessions(ctx context.Context) ([]*model.Session, error) {
	preloads := graphql.GetOperationContext(ctx)

	return r.resolveSessions(ctx, preloads.RawQuery)
}

func (r *queryResolver) Course(ctx context.Context, where model.CourseQuery) (*model.Course, error) {
	preloads := graphql.GetOperationContext(ctx)

	return r.resolveCourse(ctx, preloads)
}

func (r *queryResolver) Session(ctx context.Context, where model.SessionQuery) (*model.Session, error) {
	preloads := graphql.GetOperationContext(ctx)

	return r.resolveSession(ctx, preloads.RawQuery)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
