package builder

import (
	"time"

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
			BaseEntity: store.BaseEntity{},
			ID:         fake.CharactersN(10),
			CourseID:   fake.CharactersN(10),
			UserID:     fake.CharactersN(10),
			Map:        fake.CharactersN(10),
		},
	}
}

// WithID ...
func (c TimemapBuilder) WithID(id string) TimemapBuilder {
	c.timemap.ID = id
	return c
}

// WithPK ...
func (c TimemapBuilder) WithPK(pk string) TimemapBuilder {
	c.timemap.BaseEntity.PK = pk
	return c
}

// WithSK ...
func (c TimemapBuilder) WithSK(sk string) TimemapBuilder {
	c.timemap.BaseEntity.SK = sk
	return c
}

// WithCourseID ...
func (c TimemapBuilder) WithCourseID(id string) TimemapBuilder {
	c.timemap.CourseID = id
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

// WithDateCreated ...
func (c TimemapBuilder) WithDateCreated(t time.Time) TimemapBuilder {
	c.timemap.DateCreated = t
	return c
}

// Build ...
func (c TimemapBuilder) Build() store.Timemap {
	return c.timemap
}
