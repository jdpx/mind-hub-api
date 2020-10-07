package builder

import "github.com/jdpx/mind-hub-api/pkg/graphql/graph/model"

// SessionBuilder ...
type SessionBuilder struct {
	session model.Session
}

// NewSessionBuilder ...
func NewSessionBuilder() *SessionBuilder {
	return &SessionBuilder{
		session: model.Session{},
	}
}

// WithTitle ...
func (c SessionBuilder) WithTitle(title string) SessionBuilder {
	c.session.Title = title
	return c
}

// Build ...
func (c SessionBuilder) Build() model.Session {
	return c.session
}
