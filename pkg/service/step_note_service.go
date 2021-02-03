//go:generate mockgen -source=step_note_service.go -destination=./mocks/step_note_service.go -package=servicemocks

package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/sirupsen/logrus"
)

type StepNoteServicer interface {
	Get(ctx context.Context, sID, uID string) (*StepNote, error)
	Update(ctx context.Context, sID, uID, value string) (*StepNote, error)
}

type StepNoteService struct {
	store store.NoteRepositor
}

// NewStepNoteService ...
func NewStepNoteService(rep store.NoteRepositor) *StepNoteService {
	r := &StepNoteService{
		store: rep,
	}

	return r
}

func (s StepNoteService) Get(ctx context.Context, sID, uID string) (*StepNote, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.SessionIDKey: sID,
		logging.UserIDKey:    uID,
	})

	cn, err := s.store.Get(ctx, sID, uID)
	if err != nil {
		log.Error("error occurred getting session note from store", err)

		return nil, err
	}

	if cn == nil {
		log.Info("session note not found in store")

		return nil, ErrNotFound
	}

	return &StepNote{
		ID:     cn.ID,
		StepID: cn.EntityID,
		UserID: cn.UserID,
		Value:  cn.Value,
	}, nil
}

func (s StepNoteService) Update(ctx context.Context, sID, uID, value string) (*StepNote, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.SessionIDKey: sID,
		logging.UserIDKey:    uID,
	})

	m := store.Note{
		// ID:       *input.ID,
		EntityID: sID,
		UserID:   uID,
		Value:    value,
	}

	cn, err := s.store.Update(ctx, m)
	if err != nil {
		log.Error("error occurred updating step note in store", err)

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
		StepID: cn.EntityID,
		UserID: cn.UserID,
		Value:  cn.Value,
	}, nil
}
