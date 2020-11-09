package store

import "time"

// Progress ...
type Progress struct {
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
