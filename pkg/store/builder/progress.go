package builder

import (
	"time"

	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
)

// ProgressBuilder ...
type ProgressBuilder struct {
	progress store.Progress
}

// NewProgressBuilder ...
func NewProgressBuilder() *ProgressBuilder {
	return &ProgressBuilder{
		progress: store.Progress{
			BaseEntity: store.BaseEntity{},
			ID:         fake.CharactersN(10),
			UserID:     fake.CharactersN(10),
			EntityID:   fake.CharactersN(10),
			State:      store.STATUS_STARTED,
		},
	}
}

// WithID ...
func (c ProgressBuilder) WithID(id string) ProgressBuilder {
	c.progress.ID = id
	return c
}

// WithPK ...
func (c ProgressBuilder) WithPK(pk string) ProgressBuilder {
	c.progress.BaseEntity.PK = pk
	return c
}

// WithSK ...
func (c ProgressBuilder) WithSK(sk string) ProgressBuilder {
	c.progress.BaseEntity.SK = sk
	return c
}

// WithUserID ...
func (c ProgressBuilder) WithUserID(id string) ProgressBuilder {
	c.progress.UserID = id
	return c
}

// WithEntityID ...
func (c ProgressBuilder) WithEntityID(id string) ProgressBuilder {
	c.progress.EntityID = id
	return c
}

// Completed ...
func (c ProgressBuilder) Completed() ProgressBuilder {
	c.progress.State = store.STATUS_COMPLETED
	return c
}

// WithDateStarted ...
func (c ProgressBuilder) WithDateStarted(t time.Time) ProgressBuilder {
	c.progress.DateStarted = t
	return c
}

// Build ...
func (c ProgressBuilder) Build() store.Progress {
	return c.progress
}
