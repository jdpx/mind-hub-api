package builder

import (
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
)

// CourseNoteBuilder ...
type CourseNoteBuilder struct {
	note store.CourseNote
}

// NewCourseNoteBuilder ...
func NewCourseNoteBuilder() *CourseNoteBuilder {
	return &CourseNoteBuilder{
		note: store.CourseNote{
			ID:       fake.CharactersN(10),
			CourseID: fake.CharactersN(10),
			UserID:   fake.CharactersN(10),
			Value:    fake.Sentences(),
		},
	}
}

// WithID ...
func (c CourseNoteBuilder) WithID(id string) CourseNoteBuilder {
	c.note.ID = id
	return c
}

// WithUserID ...
func (c CourseNoteBuilder) WithUserID(id string) CourseNoteBuilder {
	c.note.UserID = id
	return c
}

// WithCourseID ...
func (c CourseNoteBuilder) WithCourseID(id string) CourseNoteBuilder {
	c.note.CourseID = id
	return c
}

// WithValue ...
func (c CourseNoteBuilder) WithValue(t string) CourseNoteBuilder {
	c.note.Value = t
	return c
}

// Build ...
func (c CourseNoteBuilder) Build() store.CourseNote {
	return c.note
}
