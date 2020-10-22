package event

// CourseStarted ...
type CourseStarted struct {
	CourseID string `json:"courseID"`
	UserID   string `json:"userID"`
}
