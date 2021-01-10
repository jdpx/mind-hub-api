package builder

import (
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/store"
)

// TimemapBuilder ...
type TimemapBuilder struct {
	timemap store.Timemap
}

// NewTimemapBuilder ...
func NewTimemapBuilder() *TimemapBuilder {
	return &TimemapBuilder{
		timemap: store.Timemap{
			ID:     fake.CharactersN(10),
			UserID: fake.CharactersN(10),
			Map:    fake.CharactersN(10),
		},
	}
}

// WithID ...
func (c TimemapBuilder) WithID(id string) TimemapBuilder {
	c.timemap.ID = id
	return c
}

// WithUserID ...
func (c TimemapBuilder) WithUserID(id string) TimemapBuilder {
	c.timemap.UserID = id
	return c
}

// WithMap ...
func (c TimemapBuilder) WithMap(t string) TimemapBuilder {
	c.timemap.Map = t
	return c
}

// Build ...
func (c TimemapBuilder) Build() store.Timemap {
	return c.timemap
}
