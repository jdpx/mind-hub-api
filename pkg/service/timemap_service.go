//go:generate mockgen -source=timemap_service.go -destination=./mocks/timemap_service.go -package=servicemocks

package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/sirupsen/logrus"
)

type TimemapServicer interface {
	Get(ctx context.Context, uID, cID, tID string) (*Timemap, error)
	GetByCourseID(ctx context.Context, uID, cID string) ([]Timemap, error)
	Update(ctx context.Context, uID, cID, tID, value string) (*Timemap, error)
}

type TimemapService struct {
	store store.TimemapRepositor
}

// NewTimemapService ...
func NewTimemapService(cms store.TimemapRepositor) *TimemapService {
	r := &TimemapService{
		store: cms,
	}

	return r
}

func (s TimemapService) Get(ctx context.Context, uID, cID, tID string) (*Timemap, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.UserIDKey:    uID,
		logging.CourseIDKey:  cID,
		logging.TimemapIDKey: tID,
	})

	timemap, err := s.store.Get(ctx, uID, cID, tID)

	if err != nil {
		log.Error("error getting timemap by id from store", err)

		return nil, err
	}

	if timemap == nil {
		log.Error("timemap not found in store")

		return nil, ErrNotFound
	}

	return &Timemap{
		ID:          timemap.ID,
		CourseID:    timemap.CourseID,
		UserID:      timemap.UserID,
		Map:         timemap.Map,
		DateCreated: timemap.DateCreated,
		DateUpdated: timemap.DateUpdated,
	}, nil
}

func (s TimemapService) GetByCourseID(ctx context.Context, uID, cID string) ([]Timemap, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.UserIDKey:   uID,
		logging.CourseIDKey: cID,
	})

	tms, err := s.store.GetByCourseID(ctx, uID, cID)

	if err != nil {
		log.Error("error getting timemap by courseID from store", err)

		return nil, err
	}

	timemaps := []Timemap{}

	if len(tms) == 0 {
		return timemaps, nil
	}

	for _, tm := range tms {
		timemaps = append(timemaps, Timemap{
			ID:          tm.ID,
			CourseID:    tm.CourseID,
			UserID:      tm.UserID,
			Map:         tm.Map,
			DateCreated: tm.DateCreated,
			DateUpdated: tm.DateUpdated,
		})
	}

	return timemaps, nil
}

func (s TimemapService) Update(ctx context.Context, uID, cID, tID, value string) (*Timemap, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.UserIDKey:    uID,
		logging.CourseIDKey:  cID,
		logging.TimemapIDKey: tID,
	})

	timemap, err := s.store.Update(ctx, &store.Timemap{
		ID:       tID,
		CourseID: cID,
		UserID:   uID,
		Map:      value,
	})
	if err != nil {
		log.Error("error updating timemap from store", err)

		return nil, err
	}

	return &Timemap{
		ID:          timemap.ID,
		CourseID:    timemap.CourseID,
		UserID:      timemap.UserID,
		Map:         timemap.Map,
		DateCreated: timemap.DateCreated,
		DateUpdated: timemap.DateUpdated,
	}, nil
}
