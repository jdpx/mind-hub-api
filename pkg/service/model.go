package service

import (
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
)

// CourseFromCMS ...
func CourseFromCMS(gc *graphcms.Course) *Course {
	return &Course{
		ID:          gc.ID,
		Title:       gc.Title,
		Description: gc.Description,
		Sessions:    []*Session{},
	}
}

// CoursesFromCMS ...
func CoursesFromCMS(gcs []*graphcms.Course) []*Course {
	cs := []*Course{}

	for _, gc := range gcs {
		cs = append(cs, CourseFromCMS(gc))
	}

	return cs
}

// SessionFromCMS ...
func SessionFromCMS(gs *graphcms.Session) *Session {
	s := &Session{
		ID:          gs.ID,
		Title:       gs.Title,
		Description: gs.Description,
		Steps:       []*Step{},
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
func SessionsFromCMS(gcs []*graphcms.Session) []*Session {
	cs := []*Session{}

	for _, gc := range gcs {
		cs = append(cs, SessionFromCMS(gc))
	}

	return cs
}

// StepFromCMS ...
func StepFromCMS(gc *graphcms.Step) *Step {
	s := &Step{
		ID:          gc.ID,
		Title:       gc.Title,
		Description: gc.Description,
		Type:        gc.Type,
		VideoURL:    gc.VideoURL,
		Question:    gc.Question,
	}

	if gc.Audio != nil {
		s.AudioURL = &gc.Audio.URL
	}

	return s
}
