package graphql

import (
	"github.com/jdpx/mind-hub-api/pkg/graphql/model"
	"github.com/jdpx/mind-hub-api/pkg/service"
)

// CourseFromServices ...
func CourseFromServices(gc *service.Course) *model.Course {
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

// CoursesFromServices ...
func CoursesFromServices(gcs []*service.Course) []*model.Course {
	cs := []*model.Course{}

	for _, gc := range gcs {
		cs = append(cs, CourseFromServices(gc))
	}

	return cs
}

// SessionFromServices ...
func SessionFromServices(gs *service.Session) *model.Session {
	s := &model.Session{
		ID:          gs.ID,
		Title:       gs.Title,
		Description: gs.Description,
		Steps:       []*model.Step{},
	}

	for _, step := range gs.Steps {
		s.Steps = append(s.Steps, StepFromServices(step))
	}

	if gs.Course != nil {
		s.Course = CourseFromServices(gs.Course)
	}

	return s
}

// SessionsFromServices ...
func SessionsFromServices(gcs []*service.Session) []*model.Session {
	cs := []*model.Session{}

	for _, gc := range gcs {
		cs = append(cs, SessionFromServices(gc))
	}

	return cs
}

// StepFromServices ...
func StepFromServices(gc *service.Step) *model.Step {
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

// TimemapsFromServices ...
func TimemapsFromServices(stm []service.Timemap) []*model.Timemap {
	tms := []*model.Timemap{}

	for _, gc := range stm {
		tms = append(tms, TimemapFromServices(gc))
	}

	return tms
}

// TimemapFromServices ...
func TimemapFromServices(tm service.Timemap) *model.Timemap {
	return &model.Timemap{
		ID:        tm.ID,
		Map:       tm.Map,
		UpdatedAt: tm.DateUpdated.String(),
	}
}
