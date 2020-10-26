package builder

import (
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
)

// CourseBuilder ...
type CourseBuilder struct {
	course graphcms.Course
}

// NewCourseBuilder ...
func NewCourseBuilder() *CourseBuilder {
	return &CourseBuilder{
		course: graphcms.Course{
			ID:          fake.CharactersN(10),
			Title:       fake.Title(),
			Description: fake.Sentences(),
		},
	}
}

// WithID ...
func (c CourseBuilder) WithID(id string) CourseBuilder {
	c.course.ID = id
	return c
}

// WithTitle ...
func (c CourseBuilder) WithTitle(title string) CourseBuilder {
	c.course.Title = title
	return c
}

// Build ...
func (c CourseBuilder) Build() graphcms.Course {
	return c.course
}
