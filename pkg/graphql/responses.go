package graphql

import (
	"github.com/jdpx/mind-hub-api/pkg/graphql/model"
	"github.com/jdpx/mind-hub-api/pkg/service"
)

// CourseFromCMS ...
func CourseFromCMS(gc *service.Course) *model.Course {
	return &model.Course{
		ID:          gc.ID,
		Title:       gc.Title,
		Description: gc.Description,
		Sessions:    []*model.Session{},
	}
}

// CourseProgressFromService ...
func CourseProgressFromService(cp *service.CourseProgress) *model.CourseProgress {
	return &model.CourseProgress{
		ID:             cp.ID,
		State:          cp.State,
		CompletedSteps: cp.CompletedSteps,
		DateStarted:    cp.DateStarted.String(),
	}
}

// CourseNoteFromService ...
func CourseNoteFromService(cn *service.CourseNote) *model.CourseNote {
	return &model.CourseNote{
		ID:       cn.ID,
		CourseID: cn.CourseID,
		UserID:   cn.UserID,
		Value:    cn.Value,
	}
}

// CoursesFromCMS ...
func CoursesFromCMS(gcs []*service.Course) []*model.Course {
	cs := []*model.Course{}

	for _, gc := range gcs {
		cs = append(cs, CourseFromCMS(gc))
	}

	return cs
}

// SessionFromCMS ...
func SessionFromCMS(gs *service.Session) *model.Session {
	s := &model.Session{
		ID:          gs.ID,
		Title:       gs.Title,
		Description: gs.Description,
		Steps:       []*model.Step{},
	}

	for _, step := range gs.Steps {
		s.Steps = append(s.Steps, StepFromCMS(step))
	}

	if gs.Course != nil {
		s.Course = CourseFromCMS(gs.Course)
	}

	return s
}

// SessionsFromCMS ...
func SessionsFromCMS(gcs []*service.Session) []*model.Session {
	cs := []*model.Session{}

	for _, gc := range gcs {
		cs = append(cs, SessionFromCMS(gc))
	}

	return cs
}

// StepFromCMS ...
func StepFromCMS(gc *service.Step) *model.Step {
	return &model.Step{
		ID:          gc.ID,
		Title:       gc.Title,
		Description: gc.Description,
		Type:        gc.Type,
		VideoURL:    gc.VideoURL,
		AudioURL:    gc.AudioURL,
		Question:    gc.Question,
	}
}
