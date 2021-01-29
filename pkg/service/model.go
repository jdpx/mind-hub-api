package service

import (
	"fmt"
	"time"
)

var (
	// ErrNotFound ...
	ErrNotFound = fmt.Errorf("record not found")
)

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

// StepProgress ...
type StepProgress struct {
	ID            string     `json:"id,omitempty"`
	StepID        string     `json:"stepID,omitempty"`
	UserID        string     `json:"userID,omitempty"`
	State         string     `json:"progressState,omitempty"`
	DateStarted   *time.Time `json:"dateStarted,omitempty"`
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
