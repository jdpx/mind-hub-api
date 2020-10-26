package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jdpx/mind-hub-api/pkg/event"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/generated"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/request"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	preloads := graphql.GetOperationContext(ctx)

	return r.resolveCourses(ctx, preloads.RawQuery)
}

func (r *queryResolver) Course(ctx context.Context, where model.CourseQuery) (*model.Course, error) {
	preloads := graphql.GetOperationContext(ctx)

	return r.resolveCourse(ctx, preloads)
}

func (r *queryResolver) Session(ctx context.Context, where model.SessionQuery) (*model.Session, error) {
	preloads := graphql.GetOperationContext(ctx)

	return r.resolveSession(ctx, preloads)
}

func (r *queryResolver) Step(ctx context.Context, where model.StepQuery) (*model.Step, error) {
	preloads := graphql.GetOperationContext(ctx)

	return r.resolveStep(ctx, preloads)
}

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type courseResolver struct{ *Resolver }

func (r *courseResolver) Sessions(ctx context.Context, obj *model.Course) ([]*model.Session, error) {
	log := logging.NewFromResolver(ctx)

	log.Info("course sessions resolver got called", obj.ID)

	return r.resolveCourseSessions(ctx, obj.ID)
}

func (r *courseResolver) Progress(ctx context.Context, obj *model.Course) (*model.Progress, error) {
	log := logging.NewFromResolver(ctx)

	log.Info("progress resolver got called")

	return &model.Progress{Value: "foo"}, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CourseStarted(ctx context.Context, input model.CourseStarted) (bool, error) {
	log := logging.NewFromResolver(ctx)

	userID, err := request.GetUserID(ctx)
	if err != nil {
		return false, err
	}

	event := event.CourseStarted{
		CourseID: input.CourseID,
		UserID:   userID,
	}

	log.Info("CourseStarted called", userID)

	err = r.store.Put(ctx, event)
	if err != nil {
		return false, err
	}

	return false, nil
}
