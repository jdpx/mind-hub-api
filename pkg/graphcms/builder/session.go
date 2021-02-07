package builder

import (
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
)

// SessionBuilder ...
type SessionBuilder struct {
	session graphcms.Session
}

// NewSessionBuilder ...
func NewSessionBuilder() *SessionBuilder {
	return &SessionBuilder{
		session: graphcms.Session{
			ID:          fake.CharactersN(10),
			Title:       fake.Title(),
			Description: fake.Sentences(),
		},
	}
}

// WithID ...
func (c SessionBuilder) WithID(id string) SessionBuilder {
	c.session.ID = id
	return c
}

// WithTitle ...
func (c SessionBuilder) WithTitle(title string) SessionBuilder {
	c.session.Title = title
	return c
}

// WithCourse ...
func (c SessionBuilder) WithCourse(course *graphcms.Course) SessionBuilder {
	c.session.Course = course
	return c
}

// WithSteps ...
func (c SessionBuilder) WithSteps(steps ...*graphcms.Step) SessionBuilder {
	c.session.Steps = steps
	return c
}

// Build ...
func (c SessionBuilder) Build() graphcms.Session {
	return c.session
}
