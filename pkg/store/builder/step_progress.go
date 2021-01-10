package builder

import (
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
)

// StepProgressBuilder ...
type StepProgressBuilder struct {
	progress store.StepProgress
}

// NewStepProgressBuilder ...
func NewStepProgressBuilder() *StepProgressBuilder {
	return &StepProgressBuilder{
		progress: store.StepProgress{
			ID:     fake.CharactersN(10),
			StepID: fake.CharactersN(10),
			UserID: fake.CharactersN(10),
			State:  store.STATUS_STARTED,
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
	c.progress.StepID = id
	return c
}

// Completed ...
func (c StepProgressBuilder) Completed() StepProgressBuilder {
	c.progress.State = store.STATUS_COMPLETED
	return c
}

// Build ...
func (c StepProgressBuilder) Build() store.StepProgress {
	return c.progress
}
