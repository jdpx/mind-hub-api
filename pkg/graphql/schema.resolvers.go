package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/graphql/generated"
	"github.com/jdpx/mind-hub-api/pkg/graphql/model"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/jdpx/mind-hub-api/pkg/service"
)

func (r *courseResolver) SessionCount(ctx context.Context, obj *model.Course) (int, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("course sessions count resolver got called", obj.ID)

	count, err := r.service.Session.CountByCourseID(ctx, obj.ID)
	if err != nil {
		log.Error("error occurred getting session count", err)

		return 0, err
	}

	return count, nil
}

func (r *courseResolver) StepCount(ctx context.Context, obj *model.Course) (int, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("course step count resolver got called", obj.ID)

	count, err := r.service.Step.CountByCourseID(ctx, obj.ID)
	if err != nil {
		log.Error("error occurred getting step count", err)

		return 0, err
	}

	return count, nil
}

func (r *courseResolver) Sessions(ctx context.Context, obj *model.Course) ([]*model.Session, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("course sessions resolver got called", obj.ID)

	gss, err := r.service.Session.GetByCourseID(ctx, obj.ID)
	if err != nil {
		log.Error("error occurred getting sessions by course ID", err)

		return nil, err
	}

	ss := SessionsFromCMS(gss)

	return ss, nil
}

func (r *courseResolver) Note(ctx context.Context, obj *model.Course) (*model.CourseNote, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Course Note resolver got called", obj.ID)

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error occurred getting request user", err)

		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	note, err := r.service.CourseNote.Get(ctx, obj.ID, userID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Info("course note not found")

			return nil, nil
		}

		log.Error("Error occurred getting Course Note %w", err)
		return nil, err
	}

	return CourseNoteFromService(note), nil
}

func (r *courseResolver) Progress(ctx context.Context, obj *model.Course) (*model.CourseProgress, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("get progress resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error occurred getting request user", err)

		return nil, fmt.Errorf("error occurred getting course progress1 %w", err)
	}

	cp, err := r.service.CourseProgress.Get(ctx, obj.ID, userID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Info("course progress not found")

			return nil, nil
		}

		log.Error("Error occurred getting Course Progress1 %w", err)
		return nil, err
	}

	return CourseProgressFromService(cp), nil
}

func (r *mutationResolver) CourseStarted(ctx context.Context, input model.CourseStarted) (*model.Course, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("course started resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error occurred getting request user", err)

		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	_, err = r.service.CourseProgress.Start(ctx, input.CourseID, userID)
	if err != nil {
		log.Error("error starting Course", err)

		return nil, err
	}

	return &model.Course{
		ID: input.CourseID,
	}, nil
}

func (r *mutationResolver) UpdateCourseNote(ctx context.Context, input model.UpdatedCourseNote) (*model.Course, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Update Course Note resolver called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error occurred getting request user", err)

		return nil, err
	}

	note, err := r.service.CourseNote.Update(ctx, input.CourseID, userID, input.Value)
	if err != nil {
		return nil, err
	}

	return &model.Course{
		ID:   input.CourseID,
		Note: CourseNoteFromService(note),
	}, nil
}

func (r *mutationResolver) StepStarted(ctx context.Context, input model.StepStarted) (*model.Step, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("step started resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error occurred getting request user", err)

		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	_, err = r.service.StepProgress.Start(ctx, input.ID, userID)
	if err != nil {
		log.Error("error putting record in store", err)
		return nil, err
	}

	return &model.Step{
		ID: input.ID,
	}, nil
}

func (r *mutationResolver) StepCompleted(ctx context.Context, input model.StepCompleted) (*model.Step, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("step completed resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error occurred getting request user", err)

		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	_, err = r.service.StepProgress.Complete(ctx, input.ID, userID)
	if err != nil {
		log.Error("error putting record in store", err)
		return nil, err
	}

	return &model.Step{
		ID: input.ID,
	}, nil
}

func (r *mutationResolver) UpdateStepNote(ctx context.Context, input model.UpdatedStepNote) (*model.Step, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Update Step Note resolver called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error occurred getting request user", err)

		return nil, err
	}

	note, err := r.service.StepNote.Update(ctx, input.StepID, userID, input.Value)
	if err != nil {
		return nil, err
	}

	return &model.Step{
		ID: input.StepID,
		Note: &model.StepNote{
			ID:     note.ID,
			StepID: note.StepID,
			UserID: note.UserID,
			Value:  note.Value,
		},
	}, nil
}

func (r *mutationResolver) UpdateTimemap(ctx context.Context, input model.UpdatedTimemap) (*model.Timemap, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Update Timemap resolver called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error occurred getting request user", err)

		return nil, err
	}

	timemap, err := r.service.Timemap.Update(ctx, userID, input.Map)
	if err != nil {
		log.Error("An error occurred getting Timemap", err)

		return nil, err
	}

	return &model.Timemap{
		Map:       timemap.Map,
		UpdatedAt: timemap.DateUpdated.String(),
	}, nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("courses resolver called")

	cgs, err := r.service.Course.GetAll(ctx)
	if err != nil {
		log.Error("error occurred getting all courses", err)

		return nil, err
	}

	cs := CoursesFromCMS(cgs)

	return cs, nil
}

func (r *queryResolver) Course(ctx context.Context, where model.CourseQuery) (*model.Course, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("course by id resolver called", where.ID)

	cg, err := r.service.Course.GetByID(ctx, where.ID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Error("course not found")

			return nil, nil
		}

		log.Error("error occurred getting course by id", err)

		return nil, err
	}

	c := CourseFromCMS(cg)

	return c, nil
}

func (r *queryResolver) Session(ctx context.Context, where model.SessionQuery) (*model.Session, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("session by id resolver called", where.ID)

	gs, err := r.service.Session.GetByID(ctx, where.ID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Error("session not found")

			return nil, nil
		}

		log.Error("error occurred getting session by id", err)

		return nil, err
	}

	s := SessionFromCMS(gs)

	return s, nil
}

func (r *queryResolver) Step(ctx context.Context, where model.StepQuery) (*model.Step, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("step by id resolver called", where.ID)

	gs, err := r.service.Step.GetByID(ctx, where.ID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Error("step not found")

			return nil, nil
		}

		log.Error("error occurred getting step by id", err)

		return nil, err
	}

	s := StepFromCMS(gs)

	return s, nil
}

func (r *queryResolver) SessionsByCourseID(ctx context.Context, where model.SessionsByCourseIDQuery) ([]*model.Session, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Sessions By Course ID resolver got called", where.ID)

	gss, err := r.service.Session.GetByCourseID(ctx, where.ID)
	if err != nil {
		log.Error("error occurred getting sessions by course id", err)

		return nil, err
	}

	ss := SessionsFromCMS(gss)

	return ss, nil
}

func (r *queryResolver) Timemap(ctx context.Context) (*model.Timemap, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Timemap resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error occurred getting request user", err)

		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	timemap, err := r.service.Timemap.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Error("timemap not found")

			return nil, nil
		}

		log.Error("error getting Timemap", err)

		return nil, fmt.Errorf("error occurred getting Timemap %w", err)
	}

	return &model.Timemap{
		ID:        timemap.ID,
		Map:       timemap.Map,
		UpdatedAt: timemap.DateUpdated.String(),
	}, nil
}

func (r *stepResolver) Note(ctx context.Context, obj *model.Step) (*model.StepNote, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Step Note resolver got called", obj.ID)

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error occurred getting request user", err)

		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	note, err := r.service.StepNote.Get(ctx, obj.ID, userID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Info("step note not found")

			return nil, nil
		}
		log.Error("error getting step note", err)

		return nil, err
	}

	return &model.StepNote{
		ID:     note.ID,
		StepID: note.StepID,
		UserID: note.UserID,
		Value:  note.Value,
	}, nil
}

func (r *stepResolver) Progress(ctx context.Context, obj *model.Step) (*model.StepProgress, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Step Progress resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error occurred getting request user", err)

		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	progress, err := r.service.StepProgress.Get(ctx, obj.ID, userID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Info("step progress not found")

			return nil, nil
		}
		log.Error("error getting step progress", err)

		return nil, err
	}

	res := &model.StepProgress{
		ID:          progress.ID,
		State:       progress.State,
		DateStarted: progress.DateStarted.String(),
	}

	if progress.DateCompleted != nil {
		res.DateCompleted = progress.DateCompleted.String()
	}

	return res, nil
}

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Step returns generated.StepResolver implementation.
func (r *Resolver) Step() generated.StepResolver { return &stepResolver{r} }

type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type stepResolver struct{ *Resolver }
