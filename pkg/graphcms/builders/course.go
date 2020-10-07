package builder

import "github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"

// CourseBuilder ...
type CourseBuilder struct {
	course model.Course
}

// NewCourseBuilder ...
func NewCourseBuilder() *CourseBuilder {
	return &CourseBuilder{
		course: model.Course{},
	}
}

// WithTitle ...
func (c CourseBuilder) WithTitle(title string) CourseBuilder {
	c.course.Title = title
	return c
}

// Build ...
func (c CourseBuilder) Build() model.Course {
	return c.course
}
