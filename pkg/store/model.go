package store

import "time"

// Progress ...
type Progress struct {
	CourseID    string    `json:"courseID"`
	UserID      string    `json:"userID"`
	DateStarted time.Time `json:"dateStarted"`
}
