package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/event"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/generated"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/request"
)

func (r *courseResolver) Sessions(ctx context.Context, obj *model.Course) ([]*model.Session, error) {
	log := logging.NewFromResolver(ctx)

	log.Info("course sessions resolver got called", obj.ID)

	gss, err := r.graphcms.ResolveCourseSessions(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	ss := SessionsFromCMS(gss)

	return ss, nil
}

func (r *courseResolver) Progress(ctx context.Context, obj *model.Course) (*model.Progress, error) {
	log := logging.NewFromResolver(ctx)

	log.Info("progress resolver got called")

	return &model.Progress{Value: "foo"}, nil
}

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

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	cgs, err := r.graphcms.ResolveCourses(ctx)
	if err != nil {
		return nil, err
	}

	cs := CoursesFromCMS(cgs)

	return cs, nil
}

func (r *queryResolver) Course(ctx context.Context, where model.CourseQuery) (*model.Course, error) {
	cg, err := r.graphcms.ResolveCourse(ctx, where.ID)
	if err != nil {
		return nil, err
	}

	c := CourseFromCMS(cg)

	return c, nil
}

func (r *queryResolver) Session(ctx context.Context, where model.SessionQuery) (*model.Session, error) {
	gs, err := r.graphcms.ResolveSession(ctx, where.ID)
	if err != nil {
		return nil, err
	}

	s := SessionFromCMS(gs)

	return s, nil
}

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
