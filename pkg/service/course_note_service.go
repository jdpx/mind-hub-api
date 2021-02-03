//go:generate mockgen -source=course_note_service.go -destination=./mocks/course_note_service.go -package=servicemocks

package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/sirupsen/logrus"
)

type CourseNoteServicer interface {
	Get(ctx context.Context, cID, uID string) (*CourseNote, error)
	Update(ctx context.Context, cID, uID, value string) (*CourseNote, error)
}

type CourseNoteService struct {
	store store.NoteRepositor
}

// NewCourseNoteService ...
func NewCourseNoteService(rep store.NoteRepositor) *CourseNoteService {
	r := &CourseNoteService{
		store: rep,
	}

	return r
}

func (s CourseNoteService) Get(ctx context.Context, cID, uID string) (*CourseNote, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.CourseIDKey: cID,
		logging.UserIDKey:   uID,
	})

	cn, err := s.store.Get(ctx, cID, uID)
	if err != nil {
		log.Error("error occurred getting course note from store", err)

		return nil, err
	}

	if cn == nil {
		log.Info("course note not found in store")

		return nil, ErrNotFound
	}

	return &CourseNote{
		ID:       cn.ID,
		CourseID: cn.EntityID,
		UserID:   cn.UserID,
		Value:    cn.Value,
	}, nil
}

func (s CourseNoteService) Update(ctx context.Context, cID, uID, value string) (*CourseNote, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.CourseIDKey: cID,
		logging.UserIDKey:   uID,
	})

	m := store.Note{
		EntityID: cID,
		UserID:   uID,
		Value:    value,
	}

	cn, err := s.store.Update(ctx, m)
	if err != nil {
		log.Error("An error occurred updating course note", err)

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
		CourseID: cn.EntityID,
		UserID:   cn.UserID,
		Value:    cn.Value,
	}, nil
}
