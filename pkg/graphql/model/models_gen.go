// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Course struct {
	ID           string      `json:"id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	SessionCount int         `json:"sessionCount"`
	Sessions     []*Session  `json:"sessions"`
	Note         *CourseNote `json:"note"`
	Progress     *Progress   `json:"progress"`
}

type CourseNote struct {
	ID       string  `json:"id"`
	CourseID string  `json:"courseID"`
	UserID   string  `json:"userID"`
	Value    *string `json:"value"`
}

type CourseQuery struct {
	ID string `json:"id"`
}

type CourseStarted struct {
	CourseID string `json:"courseID"`
}

type Progress struct {
	Started           bool `json:"started"`
	SessionsCompleted int  `json:"sessionsCompleted"`
}

type Session struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Steps       []*Step `json:"steps"`
	Course      *Course `json:"course"`
}

type SessionQuery struct {
	ID string `json:"id"`
}

type Step struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	VideoURL    *string  `json:"videoUrl"`
	AudioURL    *string  `json:"audioUrl"`
	Question    *string  `json:"question"`
	Session     *Session `json:"session"`
}

type StepQuery struct {
	ID string `json:"id"`
}

type UpdatedCourseNote struct {
	CourseID string `json:"courseID"`
	Value    string `json:"value"`
}