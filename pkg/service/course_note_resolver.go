package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/labstack/gommon/log"
)

type CourseNoteServicer interface {
	Get(ctx context.Context, cID, uID string) (*CourseNote, error)
	Update(ctx context.Context, cID, uID, value string) (*CourseNote, error)
}

type CourseNoteResolver struct {
	store store.CourseNoteRepositor
}

// CourseNoteResolverOption ...
type CourseNoteResolverOption func(*CourseNoteResolver)

// NewCourseNoteService ...
func NewCourseNoteService(rep store.CourseNoteRepositor) *CourseNoteResolver {
	r := &CourseNoteResolver{
		store: rep,
	}

	return r
}

func (s CourseNoteResolver) Get(ctx context.Context, cID, uID string) (*CourseNote, error) {
	cn, err := s.store.Get(ctx, cID, uID)
	if err != nil {
		return nil, err
	}

	if cn == nil {
		return nil, ErrNotFound
	}

	return &CourseNote{
		ID:       cn.ID,
		CourseID: cn.CourseID,
		UserID:   cn.UserID,
		Value:    cn.Value,
	}, nil
}

func (s CourseNoteResolver) Update(ctx context.Context, cID, uID, value string) (*CourseNote, error) {
	m := store.CourseNote{
		// ID:       *input.ID,
		CourseID: cID,
		UserID:   uID,
		Value:    value,
	}

	cn, err := s.store.Update(ctx, m)
	if err != nil {
		log.Error("An error occurred updating Note", err)

		return nil, err
	}

	// cn, err := s.store.Get(ctx, cID, uID)
	// if err != nil {
	// 	return nil, err
	// }

	// if cn == nil {
	// 	return nil, ErrNotFound
	// }

	return &CourseNote{
		ID:       cn.ID,
		CourseID: cn.CourseID,
		UserID:   cn.UserID,
		Value:    cn.Value,
	}, nil
}
