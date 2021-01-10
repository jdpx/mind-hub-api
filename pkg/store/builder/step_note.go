package builder

import (
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
)

// StepNoteBuilder ...
type StepNoteBuilder struct {
	note store.StepNote
}

// NewStepNoteBuilder ...
func NewStepNoteBuilder() *StepNoteBuilder {
	return &StepNoteBuilder{
		note: store.StepNote{
			ID:     fake.CharactersN(10),
			StepID: fake.CharactersN(10),
			UserID: fake.CharactersN(10),
			Value:  fake.Sentences(),
		},
	}
}

// WithID ...
func (c StepNoteBuilder) WithID(id string) StepNoteBuilder {
	c.note.ID = id
	return c
}

// WithUserID ...
func (c StepNoteBuilder) WithUserID(id string) StepNoteBuilder {
	c.note.UserID = id
	return c
}

// WithStepID ...
func (c StepNoteBuilder) WithStepID(id string) StepNoteBuilder {
	c.note.StepID = id
	return c
}

// WithValue ...
func (c StepNoteBuilder) WithValue(t string) StepNoteBuilder {
	c.note.Value = t
	return c
}

// Build ...
func (c StepNoteBuilder) Build() store.StepNote {
	return c.note
}
