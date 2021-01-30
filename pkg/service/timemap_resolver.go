package service

import (
	"context"

	"github.com/jdpx/mind-hub-api/pkg/store"
	"github.com/labstack/gommon/log"
)

type TimemapServicer interface {
	Get(ctx context.Context, uID string) (*Timemap, error)
	Update(ctx context.Context, uID, value string) (*Timemap, error)
}

type TimemapResolver struct {
	store store.TimemapRepositor
}

// TimemapResolverOption ...
type TimemapResolverOption func(*TimemapResolver)

// NewTimemapService ...
func NewTimemapService(cms store.TimemapRepositor) *TimemapResolver {
	r := &TimemapResolver{
		store: cms,
	}

	return r
}

func (s TimemapResolver) Get(ctx context.Context, uID string) (*Timemap, error) {
	timemap, err := s.store.Get(ctx, uID)

	if err != nil {
		log.Error("An error occurred getting Timemap", err)

		return nil, err
	}

	if timemap == nil {
		return nil, ErrNotFound
	}

	return &Timemap{
		ID:        timemap.ID,
		UserID:    timemap.UserID,
		Map:       timemap.Map,
		UpdatedAt: timemap.UpdatedAt,
	}, nil
}

func (s TimemapResolver) Update(ctx context.Context, uID, value string) (*Timemap, error) {
	timemap, err := s.store.Get(ctx, uID)

	if err != nil {
		log.Error("An error occurred getting Timemap", err)

		return nil, err
	}

	if timemap == nil {
		sTm := store.Timemap{
			UserID: uID,
			Map:    value,
		}

		timemap, err = s.store.Create(ctx, sTm)
		if err != nil {
			log.Error("An error occurred creating Timemap", err)

			return nil, err
		}
	} else {
		timemap.Map = value

		timemap, err = s.store.Update(ctx, timemap)
		if err != nil {
			log.Error("An error occurred updating Timemap", err)

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
