package store

import (
	"fmt"
	"time"
)

const (
	STATUS_STARTED   = "STARTED"
	STATUS_COMPLETED = "COMPLETED"
)

func UserPK(id string) string {
	return fmt.Sprintf("USER#%s", id)
}

func ProgressSK(id string) string {
	return fmt.Sprintf("PROGRESS#%s", id)
}

func NoteSK(id string) string {
	return fmt.Sprintf("NOTE#%s", id)
}

func TimemapSK() string {
	return "TIMEMAP"
}

type BaseEntity struct {
	PK string `json:"PK"`
	SK string `json:"SK"`
}

// CourseProgress ...
type CourseProgress struct {
	BaseEntity
	CourseID    string    `json:"courseID,omitempty"`
	UserID      string    `json:"userID,omitempty"`
	State       string    `json:"progressState,omitempty"`
	DateStarted time.Time `json:"dateStarted,omitempty"`
}

// CourseNote ...
type CourseNote struct {
	BaseEntity
	ID       string `json:"id"`
	CourseID string `json:"courseID"`
	UserID   string `json:"userID"`
	Value    string `json:"value"`
}

// StepProgress ...
type StepProgress struct {
	BaseEntity
	ID            string     `json:"id,omitempty"`
	StepID        string     `json:"stepID,omitempty"`
	UserID        string     `json:"userID,omitempty"`
	State         string     `json:"progressState,omitempty"`
	DateStarted   *time.Time `json:"dateStarted,omitempty"`
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`
}

// StepNote ...
type StepNote struct {
	BaseEntity
	ID     string `json:"id"`
	StepID string `json:"stepID"`
	UserID string `json:"userID"`
	Value  string `json:"value"`
}

// Timemap ...
type Timemap struct {
	BaseEntity
	ID        string    `json:"id"`
	UserID    string    `json:"userID"`
	Map       string    `json:"map"`
	UpdatedAt time.Time `json:"updatedAt"`
}
