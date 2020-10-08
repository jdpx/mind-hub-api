package builder

import "github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"

// StepBuilder ...
type StepBuilder struct {
	step model.Step
}

// NewStepBuilder ...
func NewStepBuilder() *StepBuilder {
	return &StepBuilder{
		step: model.Step{},
	}
}

// WithTitle ...
func (c StepBuilder) WithTitle(title string) StepBuilder {
	c.step.Title = title
	return c
}

// Build ...
func (c StepBuilder) Build() model.Step {
	return c.step
}
