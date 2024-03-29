// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Course struct {
	ID           string          `json:"id"`
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	SessionCount int             `json:"sessionCount"`
	StepCount    int             `json:"stepCount"`
	FirstSession *string         `json:"firstSession"`
	Sessions     []*Session      `json:"sessions"`
	Note         *CourseNote     `json:"note"`
	Progress     *CourseProgress `json:"progress"`
}

type CourseNote struct {
	ID       string `json:"id"`
	CourseID string `json:"courseID"`
	UserID   string `json:"userID"`
	Value    string `json:"value"`
}

type CourseProgress struct {
	ID             string `json:"id"`
	State          string `json:"state"`
	CompletedSteps int    `json:"completedSteps"`
	DateStarted    string `json:"dateStarted"`
}

type CourseQuery struct {
	ID string `json:"id"`
}

type CourseStarted struct {
	CourseID string `json:"courseID"`
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

type SessionsByCourseIDQuery struct {
	ID string `json:"id"`
}

type Step struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Type        string        `json:"type"`
	VideoURL    *string       `json:"videoUrl"`
	AudioURL    *string       `json:"audioUrl"`
	Question    *string       `json:"question"`
	Session     *Session      `json:"session"`
	Note        *StepNote     `json:"note"`
	Progress    *StepProgress `json:"progress"`
}

type StepCompleted struct {
	ID string `json:"id"`
}

type StepNote struct {
	ID     string `json:"id"`
	StepID string `json:"stepID"`
	UserID string `json:"userID"`
	Value  string `json:"value"`
}

type StepProgress struct {
	ID            string `json:"id"`
	State         string `json:"state"`
	DateStarted   string `json:"dateStarted"`
	DateCompleted string `json:"dateCompleted"`
}

type StepQuery struct {
	ID string `json:"id"`
}

type StepStarted struct {
	ID string `json:"id"`
}

type Timemap struct {
	ID        string `json:"id"`
	Map       string `json:"map"`
	UpdatedAt string `json:"updatedAt"`
}

type UpdatedCourseNote struct {
	ID       *string `json:"id"`
	CourseID string  `json:"courseID"`
	Value    string  `json:"value"`
}

type UpdatedStepNote struct {
	ID     *string `json:"id"`
	StepID string  `json:"stepID"`
	Value  string  `json:"value"`
}

type UpdatedTimemap struct {
	Map string `json:"map"`
}
