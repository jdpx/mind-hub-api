package builder

import (
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
)

// StepBuilder ...
type StepBuilder struct {
	step graphcms.Step
}

// NewStepBuilder ...
func NewStepBuilder() *StepBuilder {
	return &StepBuilder{
		step: graphcms.Step{
			ID:          fake.CharactersN(10),
			Title:       fake.Title(),
			Description: fake.Sentences(),
		},
	}
}

// WithID ...
func (c StepBuilder) WithID(id string) StepBuilder {
	c.step.ID = id
	return c
}

// WithTitle ...
func (c StepBuilder) WithTitle(title string) StepBuilder {
	c.step.Title = title
	return c
}

// Build ...
func (c StepBuilder) Build() graphcms.Step {
	return c.step
}
