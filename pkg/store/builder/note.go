package builder

import (
	"time"

	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
)

// NoteBuilder ...
type NoteBuilder struct {
	note store.Note
}

// NewNoteBuilder ...
func NewNoteBuilder() *NoteBuilder {
	return &NoteBuilder{
		note: store.Note{
			BaseEntity: store.BaseEntity{},
			ID:         fake.CharactersN(10),
			EntityID:   fake.CharactersN(10),
			UserID:     fake.CharactersN(10),
			Value:      fake.Sentences(),
		},
	}
}

// WithID ...
func (c NoteBuilder) WithID(id string) NoteBuilder {
	c.note.ID = id
	return c
}

// WithPK ...
func (c NoteBuilder) WithPK(pk string) NoteBuilder {
	c.note.BaseEntity.PK = pk
	return c
}

// WithPK ...
func (c NoteBuilder) WithSK(sk string) NoteBuilder {
	c.note.BaseEntity.SK = sk
	return c
}

// WithUserID ...
func (c NoteBuilder) WithUserID(id string) NoteBuilder {
	c.note.UserID = id
	return c
}

// WithEntityID ...
func (c NoteBuilder) WithEntityID(id string) NoteBuilder {
	c.note.EntityID = id
	return c
}

// WithValue ...
func (c NoteBuilder) WithValue(t string) NoteBuilder {
	c.note.Value = t
	return c
}

// WithDateCreated ...
func (c NoteBuilder) WithDateCreated(t time.Time) NoteBuilder {
	c.note.DateCreated = t
	return c
}

// WithDateUpdated ...
func (c NoteBuilder) WithDateUpdated(t time.Time) NoteBuilder {
	c.note.DateUpdated = t
	return c
}

// Build ...
func (c NoteBuilder) Build() store.Note {
	return c.note
}
