package service

import (
	"fmt"
	"time"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
)

var (
	// ErrNotFound ...
	ErrNotFound = fmt.Errorf("record not found")
)

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

// CourseProgress ...
type CourseProgress struct {
	ID             string    `json:"id"`
	CourseID       string    `json:"courseID,omitempty"`
	UserID         string    `json:"userID,omitempty"`
	State          string    `json:"progressState,omitempty"`
	CompletedSteps int       `json:"completedSteps"`
	DateStarted    time.Time `json:"dateStarted,omitempty"`
}

// CourseNote ...
type CourseNote struct {
	ID       string `json:"id"`
	CourseID string `json:"courseID"`
	UserID   string `json:"userID"`
	Value    string `json:"value"`
}

type Session struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Steps       []*Step `json:"steps"`
	Course      *Course `json:"course"`
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

// StepProgress ...
type StepProgress struct {
	ID            string     `json:"id,omitempty"`
	StepID        string     `json:"stepID,omitempty"`
	UserID        string     `json:"userID,omitempty"`
	State         string     `json:"progressState,omitempty"`
	DateStarted   time.Time  `json:"dateStarted,omitempty"`
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`
}

// StepNote ...
type StepNote struct {
	ID     string `json:"id"`
	StepID string `json:"stepID"`
	UserID string `json:"userID"`
	Value  string `json:"value"`
}

// Timemap ...
type Timemap struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userID"`
	Map       string    `json:"map"`
	UpdatedAt time.Time `json:"updatedAt"`
}

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
