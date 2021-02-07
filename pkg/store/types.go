package store

import (
	"time"
)

const (
	STATUS_STARTED   = "STARTED"
	STATUS_COMPLETED = "COMPLETED"
)

type BaseEntity struct {
	PK string `dynamodbav:"PK"`
	SK string `dynamodbav:"SK"`
}

// CourseProgress ...
type CourseProgress struct {
	BaseEntity
	CourseID    string    `dynamodbav:"courseID,omitempty"`
	UserID      string    `dynamodbav:"userID,omitempty"`
	State       string    `dynamodbav:"progressState,omitempty"`
	DateStarted time.Time `dynamodbav:"dateStarted,omitempty"`
}

// Note ...
type Note struct {
	BaseEntity
	ID          string    `dynamodbav:"id"`
	EntityID    string    `dynamodbav:"entityID"`
	UserID      string    `dynamodbav:"userID"`
	Value       string    `dynamodbav:"value"`
	DateCreated time.Time `dynamodbav:"dateCreated"`
	DateUpdated time.Time `dynamodbav:"dateUpdated"`
}

// Progress ...
type Progress struct {
	BaseEntity
	ID            string     `dynamodbav:"id,omitempty"`
	EntityID      string     `dynamodbav:"entityID,omitempty"`
	UserID        string     `dynamodbav:"userID,omitempty"`
	State         string     `dynamodbav:"state,omitempty"`
	DateStarted   time.Time  `dynamodbav:"dateStarted,omitempty"`
	DateCompleted *time.Time `dynamodbav:"dateCompleted,omitempty"`
}

// StepNote ...
type StepNote struct {
	BaseEntity
	ID     string `dynamodbav:"id"`
	StepID string `dynamodbav:"stepID"`
	UserID string `dynamodbav:"userID"`
	Value  string `dynamodbav:"value"`
}

// Timemap ...
type Timemap struct {
	BaseEntity
	ID          string    `dynamodbav:"id"`
	UserID      string    `dynamodbav:"userID"`
	Map         string    `dynamodbav:"map"`
	DateCreated time.Time `dynamodbav:"dateCreated"`
	DateUpdated time.Time `dynamodbav:"dateUpdated"`
}
