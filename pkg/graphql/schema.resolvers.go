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

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	note, err := r.courseNoteHandler.GetNote(ctx, obj.ID, userID)
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
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	progress, err := r.courseProgressHandler.GetCourseProgress(ctx, obj.ID, userID)
	if err != nil {
		return nil, err
	}

	if progress == nil {
		return nil, nil
	}

	return &model.CourseProgress{
		ID:          progress.ID,
		DateStarted: progress.DateStarted.String(),
	}, nil
}

func (r *mutationResolver) CourseStarted(ctx context.Context, input model.CourseStarted) (*model.Course, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("course started resolver got called")

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	_, err = r.courseProgressHandler.StartCourse(ctx, input.CourseID, userID)
	if err != nil {
		log.Error("error putting record in store", err)
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

		note, err = r.courseNoteHandler.CreateNote(ctx, m)
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

		note, err = r.courseNoteHandler.UpdateNote(ctx, m)
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

	_, err = r.stepProgressHandler.StartStep(ctx, input.ID, userID)
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

	_, err = r.stepProgressHandler.CompleteStep(ctx, input.ID, userID)
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

		note, err = r.stepNoteHandler.CreateNote(ctx, m)
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

		note, err = r.stepNoteHandler.UpdateNote(ctx, m)
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

func (r *stepResolver) Note(ctx context.Context, obj *model.Step) (*model.StepNote, error) {
	log := logging.NewFromResolver(ctx)
	log.Info("Step Note resolver got called", obj.ID)

	userID, err := request.GetUserID(ctx)
	if err != nil {
		log.Error("error getting user", err)
		return nil, fmt.Errorf("error occurred getting request user ID %w", err)
	}

	note, err := r.stepNoteHandler.GetNote(ctx, obj.ID, userID)
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

	progress, err := r.stepProgressHandler.GetStepProgress(ctx, obj.ID, userID)
	if err != nil {
		return nil, err
	}

	if progress == nil {
		return nil, nil
	}

	res := &model.StepProgress{
		ID:     progress.ID,
		Status: progress.Status,
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
