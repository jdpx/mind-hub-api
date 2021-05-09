package service

import (
	"fmt"
	"time"
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
	ID             string     `json:"id"`
	CourseID       string     `json:"courseID,omitempty"`
	UserID         string     `json:"userID,omitempty"`
	State          string     `json:"progressState,omitempty"`
	CompletedSteps int        `json:"completedSteps"`
	DateStarted    time.Time  `json:"dateStarted,omitempty"`
	DateCompleted  *time.Time `json:"dateCompleted,omitempty"`
}

// CourseNote ...
type CourseNote struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"courseID"`
	UserID      string    `json:"userID"`
	Value       string    `json:"value"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
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
	ID          string    `json:"id"`
	StepID      string    `json:"stepID"`
	UserID      string    `json:"userID"`
	Value       string    `json:"value"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
}

// Timemap ...
type Timemap struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"courseID"`
	UserID      string    `json:"userID"`
	Map         string    `json:"map"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
}
