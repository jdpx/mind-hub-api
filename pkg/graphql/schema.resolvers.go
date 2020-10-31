package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/graphql/generated"
	"github.com/jdpx/mind-hub-api/pkg/graphql/model"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/jdpx/mind-hub-api/pkg/store"
)

func (r *courseResolver) SessionCount(ctx context.Context, obj *model.Course) (int, error) {
	log := logging.NewFromResolver(ctx)

	log.Info("course sessions count resolver got called", obj.ID)

	gss, err := r.graphcms.ResolveCourseSessions(ctx, obj.ID)
	if err != nil {
		return 0, err
	}

	return len(gss), nil
}

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

func (r *courseResolver) Note(ctx context.Context, obj *model.Course) (*model.CourseNote, error) {
	log := logging.NewFromResolver(ctx)

	log.Info("Course Note resolver got called", obj.ID)

	n, err := r.store.Get(ctx, "course_note", obj.ID)
	if err != nil {
		return nil, err
	}

	if n == nil {
		return nil, nil
	}

	return n.(*model.CourseNote), nil
}

func (r *courseResolver) Progress(ctx context.Context, obj *model.Course) (*model.Progress, error) {
	log := logging.NewFromResolver(ctx)

	log.Info("get progress resolver got called")

	progress, err := r.store.Get(ctx, "course_progress", obj.ID)
	if err != nil {
		return nil, err
	}

	return &model.Progress{SessionsCompleted: 0, Started: progress != nil}, nil
}

func (r *mutationResolver) CourseStarted(ctx context.Context, input model.CourseStarted) (*model.Course, error) {
	log := logging.NewFromResolver(ctx)

	log.Info("course started resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	event := store.Progress{
		CourseID: input.CourseID,
		UserID:   userID,
	}

	err = r.store.Put(ctx, "course_progress", input.CourseID, event)
	if err != nil {
		log.Error("error putting record in store", err)
		return nil, err
	}

	return &model.Course{
		ID: input.CourseID,

		Progress: &model.Progress{
			Started: true,
		},
	}, nil
}

func (r *mutationResolver) UpdateCourseNote(ctx context.Context, input model.UpdatedCourseNote) (*model.CourseNote, error) {
	log := logging.NewFromResolver(ctx)

	log.Info("Update Course Note resolver called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	note := model.CourseNote{
		ID:       "111",
		CourseID: input.CourseID,
		UserID:   userID,
		Value:    &input.Value,
	}

	err = r.store.Put(ctx, "course_note", input.CourseID, &note)

	return &note, err
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
