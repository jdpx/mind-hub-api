package graphql

import (
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/store"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// ResolverOption ...
type ResolverOption func(*Resolver)

// Resolver ...
type Resolver struct {
	graphcms              graphcms.Resolverer
	courseProgressHandler store.CourseProgressRepositor
	courseNoteHandler     store.CourseNoteRepositor
}

// NewResolver ...
func NewResolver(opts ...ResolverOption) *Resolver {
	r := &Resolver{}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

// WithCMSClient ...
func WithCMSClient(c graphcms.Resolverer) func(*Resolver) {
	return func(r *Resolver) {
		r.graphcms = c
	}
}

// WithCourseProgressHandler ...
func WithCourseProgressHandler(s store.CourseProgressRepositor) func(*Resolver) {
	return func(r *Resolver) {
		r.courseProgressHandler = s
	}
}

// WithCourseNoteRepositor ...
func WithCourseNoteRepositor(s store.CourseNoteRepositor) func(*Resolver) {
	return func(r *Resolver) {
		r.courseNoteHandler = s
	}
}
