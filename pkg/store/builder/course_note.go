package builder

import (
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
)

// CourseNoteBuilder ...
type CourseNoteBuilder struct {
	note store.Note
}

// NewCourseNoteBuilder ...
func NewCourseNoteBuilder() *CourseNoteBuilder {
	return &CourseNoteBuilder{
		note: store.Note{
			ID:       fake.CharactersN(10),
			EntityID: fake.CharactersN(10),
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

// WithEntityID ...
func (c CourseNoteBuilder) WithEntityID(id string) CourseNoteBuilder {
	c.note.EntityID = id
	return c
}

// WithValue ...
func (c CourseNoteBuilder) WithValue(t string) CourseNoteBuilder {
	c.note.Value = t
	return c
}

// Build ...
func (c CourseNoteBuilder) Build() store.Note {
	return c.note
}
