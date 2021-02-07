package builder

import (
	"time"

	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/service"
)

// StepNoteBuilder ...
type StepNoteBuilder struct {
	note service.StepNote
}

// NewStepNoteBuilder ...
func NewStepNoteBuilder() *StepNoteBuilder {
	return &StepNoteBuilder{
		note: service.StepNote{
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

// WithDateCreated ...
func (c StepNoteBuilder) WithDateCreated(t time.Time) StepNoteBuilder {
	c.note.DateCreated = t
	return c
}

// WithDateUpdated ...
func (c StepNoteBuilder) WithDateUpdated(t time.Time) StepNoteBuilder {
	c.note.DateUpdated = t
	return c
}

// Build ...
func (c StepNoteBuilder) Build() service.StepNote {
	return c.note
}
