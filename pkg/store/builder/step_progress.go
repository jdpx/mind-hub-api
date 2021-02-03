package builder

import (
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
)

// StepProgressBuilder ...
type StepProgressBuilder struct {
	progress store.Progress
}

// NewStepProgressBuilder ...
func NewStepProgressBuilder() *StepProgressBuilder {
	return &StepProgressBuilder{
		progress: store.Progress{
			ID:       fake.CharactersN(10),
			EntityID: fake.CharactersN(10),
			UserID:   fake.CharactersN(10),
			State:    store.STATUS_STARTED,
		},
	}
}

// WithID ...
func (c StepProgressBuilder) WithID(id string) StepProgressBuilder {
	c.progress.ID = id
	return c
}

// WithUserID ...
func (c StepProgressBuilder) WithUserID(id string) StepProgressBuilder {
	c.progress.UserID = id
	return c
}

// WithStepID ...
func (c StepProgressBuilder) WithStepID(id string) StepProgressBuilder {
	c.progress.EntityID = id
	return c
}

// Completed ...
func (c StepProgressBuilder) Completed() StepProgressBuilder {
	c.progress.State = store.STATUS_COMPLETED
	return c
}

// Build ...
func (c StepProgressBuilder) Build() store.Progress {
	return c.progress
}
