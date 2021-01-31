package graphql

import (
	"github.com/jdpx/mind-hub-api/pkg/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// ResolverOption ...
type ResolverOption func(*Resolver)

// Resolver ...
type Resolver struct {
	service *service.Service
}

// NewResolver ...
func NewResolver(opts ...ResolverOption) *Resolver {
	r := &Resolver{}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

// WithService ...
func WithService(c *service.Service) func(*Resolver) {
	return func(r *Resolver) {
		r.service = c
	}
}
