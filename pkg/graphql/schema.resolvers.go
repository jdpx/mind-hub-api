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

func (r *courseResolver) StepCount(ctx context.Context, obj *model.Course) (int, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("course step count resolver got called", obj.ID)

	gss, err := r.graphcms.ResolveCourseStepIDs(ctx, obj.ID)
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

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	note, err := r.courseNoteHandler.Get(ctx, obj.ID, userID)
	if err != nil {
		return nil, err
	}

	if note == nil {
		return nil, nil
	}

	return &model.CourseNote{
		ID:       note.ID,
		CourseID: note.CourseID,
		UserID:   note.UserID,
		Value:    &note.Value,
	}, nil
}

func (r *courseResolver) Progress(ctx context.Context, obj *model.Course) (*model.CourseProgress, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("get progress resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting course progress %w", err)
	}

	cp, err := r.courseProgressService.Get(ctx, obj.ID, userID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return nil, nil
		}

		log.Error("Error occurred getting Course Progress %w", err)
		return nil, err
	}

	p := CourseProgressFromService(cp)

	return p, nil

	// progress, err := r.courseProgressHandler.Get(ctx, obj.ID, userID)
	// if err != nil {
	// 	return nil, err
	// }

	// if progress == nil {
	// 	return nil, nil
	// }

	// res := &model.CourseProgress{
	// 	// ID:          progress.ID,
	// 	State:       progress.State,
	// 	DateStarted: progress.DateStarted.String(),
	// }

	// courseStepIDs, err := r.graphcms.ResolveCourseStepIDs(ctx, obj.ID)
	// if err != nil {
	// 	log.Error("error getting course steps", err)
	// 	return nil, fmt.Errorf("error occurred getting course progress %w", err)
	// }

	// if len(courseStepIDs) == 0 {
	// 	return res, nil
	// }

	// completedSteps, err := r.stepProgressHandler.GetCompletedByStepID(ctx, userID, courseStepIDs...)
	// if err != nil {
	// 	log.Error("error getting completed steps", err)
	// 	return nil, fmt.Errorf("error occurred getting course progress %w", err)
	// }

	// res.CompletedSteps = len(completedSteps)

	// return p, nil
}

func (r *mutationResolver) CourseStarted(ctx context.Context, input model.CourseStarted) (*model.Course, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("course started resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	_, err = r.courseProgressService.Start(ctx, input.CourseID, userID)
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
		return nil, err
	}

	var note *store.CourseNote

	if input.ID == nil {
		m := store.CourseNote{
			CourseID: input.CourseID,
			UserID:   userID,
			Value:    input.Value,
		}

		note, err = r.courseNoteHandler.Create(ctx, m)
		if err != nil {
			log.Error("An error occurred creating Note", err)

			return nil, err
		}
	} else {
		m := store.CourseNote{
			ID:       *input.ID,
			CourseID: input.CourseID,
			UserID:   userID,
			Value:    input.Value,
		}

		note, err = r.courseNoteHandler.Update(ctx, m)
		if err != nil {
			log.Error("An error occurred updating Note", err)

			return nil, err
		}
	}

	return &model.Course{
		ID: input.CourseID,
		Note: &model.CourseNote{
			ID:       note.ID,
			CourseID: note.CourseID,
			UserID:   note.UserID,
			Value:    &note.Value,
		},
	}, nil
}

func (r *mutationResolver) StepStarted(ctx context.Context, input model.StepStarted) (*model.Step, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("step started resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	_, err = r.stepProgressHandler.Start(ctx, input.ID, userID)
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
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	_, err = r.stepProgressHandler.Complete(ctx, input.ID, userID)
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
		return nil, err
	}

	var note *store.StepNote

	if input.ID == nil {
		m := store.StepNote{
			StepID: input.StepID,
			UserID: userID,
			Value:  input.Value,
		}

		note, err = r.stepNoteHandler.Create(ctx, m)
		if err != nil {
			log.Error("An error occurred creating Note", err)

			return nil, err
		}
	} else {
		m := store.StepNote{
			ID:     *input.ID,
			StepID: input.StepID,
			UserID: userID,
			Value:  input.Value,
		}

		note, err = r.stepNoteHandler.Update(ctx, m)
		if err != nil {
			log.Error("An error occurred updating Note", err)

			return nil, err
		}
	}

	return &model.Step{
		ID: input.StepID,
		Note: &model.StepNote{
			ID:     note.ID,
			StepID: note.StepID,
			UserID: note.UserID,
			Value:  &note.Value,
		},
	}, nil
}

func (r *mutationResolver) UpdateTimemap(ctx context.Context, input model.UpdatedTimemap) (*model.Timemap, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Update Timemap resolver called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	timemap, err := r.timemapHandler.Get(ctx, userID)

	if err != nil {
		log.Error("An error occurred getting Timemap", err)

		return nil, err
	}

	if timemap == nil {
		sTm := store.Timemap{
			UserID: userID,
			Map:    input.Map,
		}

		timemap, err = r.timemapHandler.Create(ctx, sTm)
		if err != nil {
			log.Error("An error occurred creating Timemap", err)

			return nil, err
		}
	} else {
		timemap.Map = input.Map

		timemap, err = r.timemapHandler.Update(ctx, timemap)
		if err != nil {
			log.Error("An error occurred updating Timemap", err)

			return nil, err
		}
	}

	return &model.Timemap{
		Map:       timemap.Map,
		UpdatedAt: timemap.UpdatedAt.String(),
	}, nil
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

	if cg == nil {
		return nil, nil
	}

	c := CourseFromCMS(cg)

	return c, nil
}

func (r *queryResolver) Session(ctx context.Context, where model.SessionQuery) (*model.Session, error) {
	gs, err := r.graphcms.ResolveSession(ctx, where.ID)
	if err != nil {
		return nil, err
	}

	if gs == nil {
		return nil, nil
	}

	s := SessionFromCMS(gs)

	return s, nil
}

func (r *queryResolver) Step(ctx context.Context, where model.StepQuery) (*model.Step, error) {
	gs, err := r.graphcms.ResolveStep(ctx, where.ID)
	if err != nil {
		return nil, err
	}

	if gs == nil {
		return nil, nil
	}

	s := StepFromCMS(gs)

	return s, nil
}

func (r *queryResolver) SessionsByCourseID(ctx context.Context, where model.SessionsByCourseIDQuery) ([]*model.Session, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Sessions By Course ID resolver got called", where.ID)

	gss, err := r.graphcms.ResolveCourseSessions(ctx, where.ID)
	if err != nil {
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
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	timemap, err := r.timemapHandler.Get(ctx, userID)
	if err != nil {
		log.Error("error getting Timemap", err)
		return nil, fmt.Errorf("error occurred getting Timemap %w", err)
	}

	return &model.Timemap{
		ID:        timemap.ID,
		Map:       timemap.Map,
		UpdatedAt: timemap.UpdatedAt.String(),
	}, nil
}

func (r *stepResolver) Note(ctx context.Context, obj *model.Step) (*model.StepNote, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Step Note resolver got called", obj.ID)

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	note, err := r.stepNoteHandler.Get(ctx, obj.ID, userID)
	if err != nil {
		return nil, err
	}

	if note == nil {
		return nil, nil
	}

	return &model.StepNote{
		ID:     note.ID,
		StepID: note.StepID,
		UserID: note.UserID,
		Value:  &note.Value,
	}, nil
}

func (r *stepResolver) Progress(ctx context.Context, obj *model.Step) (*model.StepProgress, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Step Progress resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	progress, err := r.stepProgressHandler.Get(ctx, obj.ID, userID)
	if err != nil {
		return nil, err
	}

	if progress == nil {
		return nil, nil
	}

	res := &model.StepProgress{
		ID:    progress.ID,
		State: progress.State,
	}

	if progress.DateStarted != nil {
		res.DateStarted = progress.DateStarted.String()
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
