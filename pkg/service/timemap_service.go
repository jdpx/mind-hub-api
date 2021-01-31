//go:generate mockgen -source=timemap_service.go -destination=./mocks/timemap_service.go -package=servicemocks

package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/sirupsen/logrus"
)

type TimemapServicer interface {
	Get(ctx context.Context, uID string) (*Timemap, error)
	Update(ctx context.Context, uID, value string) (*Timemap, error)
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

func (s TimemapService) Get(ctx context.Context, uID string) (*Timemap, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.UserIDKey: uID,
	})

	timemap, err := s.store.Get(ctx, uID)

	if err != nil {
		log.Error("error getting timemap by id from store", err)

		return nil, err
	}

	if timemap == nil {
		log.Error("timemap not found in store")

		return nil, ErrNotFound
	}

	return &Timemap{
		ID:        timemap.ID,
		UserID:    timemap.UserID,
		Map:       timemap.Map,
		UpdatedAt: timemap.UpdatedAt,
	}, nil
}

func (s TimemapService) Update(ctx context.Context, uID, value string) (*Timemap, error) {
	log := logging.NewFromResolver(ctx).WithFields(logrus.Fields{
		logging.UserIDKey: uID,
	})

	timemap, err := s.store.Get(ctx, uID)

	if err != nil {
		log.Error("error getting timemap from store", err)

		return nil, err
	}

	if timemap == nil {
		log.Info("creating new timemap in store")

		sTm := store.Timemap{
			UserID: uID,
			Map:    value,
		}

		timemap, err = s.store.Create(ctx, sTm)
		if err != nil {
			log.Error("error creating timemap from store", err)

			return nil, err
		}
	} else {
		log.Info("updating timemap in store")
		timemap.Map = value

		timemap, err = s.store.Update(ctx, timemap)
		if err != nil {
			log.Error("error updating timemap from store", err)

			return nil, err
		}
	}

	return &Timemap{
		ID:        timemap.ID,
		UserID:    timemap.UserID,
		Map:       timemap.Map,
		UpdatedAt: timemap.UpdatedAt,
	}, nil
}
