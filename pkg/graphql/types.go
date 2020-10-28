package graphql

import "github.com/jdpx/mind-hub-api/pkg/graphql/model"

// CoursesResponse ...
type CoursesResponse struct {
	Courses []*model.Course `json:"courses"`
}

// CourseResponse ...
type CourseResponse struct {
	Course *model.Course `json:"course"`
}

// SessionsResponse ...
type SessionsResponse struct {
	Sessions []*model.Session `json:"sessions"`
}

// SessionResponse ...
type SessionResponse struct {
	Session *model.Session `json:"session"`
}

// StepResponse ...
type StepResponse struct {
	Step *model.Step `json:"step"`
}
