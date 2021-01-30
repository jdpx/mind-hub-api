package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/labstack/gommon/log"
)

type StepNoteServicer interface {
	Get(ctx context.Context, sID, uID string) (*StepNote, error)
	Update(ctx context.Context, sID, uID, value string) (*StepNote, error)
}

type StepNoteResolver struct {
	store store.StepNoteRepositor
}

// StepNoteResolverOption ...
type StepNoteResolverOption func(*StepNoteResolver)

// NewStepNoteService ...
func NewStepNoteService(rep store.StepNoteRepositor) *StepNoteResolver {
	r := &StepNoteResolver{
		store: rep,
	}

	return r
}

func (s StepNoteResolver) Get(ctx context.Context, sID, uID string) (*StepNote, error) {
	cn, err := s.store.Get(ctx, sID, uID)
	if err != nil {
		return nil, err
	}

	if cn == nil {
		return nil, ErrNotFound
	}

	return &StepNote{
		ID:     cn.ID,
		StepID: cn.StepID,
		UserID: cn.UserID,
		Value:  cn.Value,
	}, nil
}

func (s StepNoteResolver) Update(ctx context.Context, sID, uID, value string) (*StepNote, error) {
	m := store.StepNote{
		// ID:       *input.ID,
		StepID: sID,
		UserID: uID,
		Value:  value,
	}

	cn, err := s.store.Update(ctx, m)
	if err != nil {
		log.Error("An error occurred updating Note", err)

		return nil, err
	}

	// cn, err := s.store.Get(ctx, sID, uID)
	// if err != nil {
	// 	return nil, err
	// }

	// if cn == nil {
	// 	return nil, ErrNotFound
	// }

	return &StepNote{
		ID:     cn.ID,
		StepID: cn.StepID,
		UserID: cn.UserID,
		Value:  cn.Value,
	}, nil
}
