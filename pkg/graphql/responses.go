package graphql

import (
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/graphql/model"
	"github.com/jdpx/mind-hub-api/pkg/service"
)

// CourseFromCMS ...
func CourseFromCMS(gc *graphcms.Course) *model.Course {
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

// CoursesFromCMS ...
func CoursesFromCMS(gcs []*graphcms.Course) []*model.Course {
	cs := []*model.Course{}

	for _, gc := range gcs {
		cs = append(cs, CourseFromCMS(gc))
	}

	return cs
}

// SessionFromCMS ...
func SessionFromCMS(gs *graphcms.Session) *model.Session {
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
func SessionsFromCMS(gcs []*graphcms.Session) []*model.Session {
	cs := []*model.Session{}

	for _, gc := range gcs {
		cs = append(cs, SessionFromCMS(gc))
	}

	return cs
}

// StepFromCMS ...
func StepFromCMS(gc *graphcms.Step) *model.Step {
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
