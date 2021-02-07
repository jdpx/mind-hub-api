package builder

import (
	"time"

	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/service"
)

// TimemapBuilder ...
type TimemapBuilder struct {
	timemap service.Timemap
}

// NewTimemapBuilder ...
func NewTimemapBuilder() *TimemapBuilder {
	return &TimemapBuilder{
		timemap: service.Timemap{
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

// WithDateUpdated ...
func (c TimemapBuilder) WithDateUpdated(t time.Time) TimemapBuilder {
	c.timemap.DateUpdated = t
	return c
}

// Build ...
func (c TimemapBuilder) Build() service.Timemap {
	return c.timemap
}
