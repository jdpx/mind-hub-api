package store

import "time"

// CourseProgress ...
type CourseProgress struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"courseID"`
	UserID      string    `json:"userID"`
	DateStarted time.Time `json:"dateStarted"`
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
	ID          string    `json:"id"`
	StepID      string    `json:"stepID"`
	UserID      string    `json:"userID"`
	DateStarted time.Time `json:"dateStarted"`
}

// StepNote ...
type StepNote struct {
	ID     string `json:"id"`
	StepID string `json:"stepID"`
	UserID string `json:"userID"`
	Value  string `json:"value"`
}
