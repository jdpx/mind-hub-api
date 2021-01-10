package builder

import (
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
)

// CourseProgressBuilder ...
type CourseProgressBuilder struct {
	progress store.CourseProgress
}

// NewCourseProgressBuilder ...
func NewCourseProgressBuilder() *CourseProgressBuilder {
	return &CourseProgressBuilder{
		progress: store.CourseProgress{
			ID:       fake.CharactersN(10),
			CourseID: fake.CharactersN(10),
			UserID:   fake.CharactersN(10),
			State:    store.STATUS_STARTED,
		},
	}
}

// WithID ...
func (c CourseProgressBuilder) WithID(id string) CourseProgressBuilder {
	c.progress.ID = id
	return c
}

// WithUserID ...
func (c CourseProgressBuilder) WithUserID(id string) CourseProgressBuilder {
	c.progress.UserID = id
	return c
}

// WithCourseID ...
func (c CourseProgressBuilder) WithCourseID(id string) CourseProgressBuilder {
	c.progress.CourseID = id
	return c
}

// Completed ...
func (c CourseProgressBuilder) Completed() CourseProgressBuilder {
	c.progress.State = store.STATUS_COMPLETED
	return c
}

// Build ...
func (c CourseProgressBuilder) Build() store.CourseProgress {
	return c.progress
}
