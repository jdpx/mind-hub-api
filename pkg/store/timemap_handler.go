//go:generate mockgen -source=timemap_handler.go -destination=./mocks/timemap_handler.go -package=storemocks
package store

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

const (
	timemapTableName = "timemap"
)

// TimemapRepositor ...
type TimemapRepositor interface {
	Get(ctx context.Context, uID string) (*Timemap, error)
	Create(ctx context.Context, tm Timemap) (*Timemap, error)
	Update(ctx context.Context, tm *Timemap) (*Timemap, error)
}

// TimemapHandler ...
type TimemapHandler struct {
	db Storer
}

// NewTimemapHandler ...
func NewTimemapHandler(client Storer) TimemapHandler {
	return TimemapHandler{
		db: client,
	}
}

// Get ...
func (c TimemapHandler) Get(ctx context.Context, uID string) (*Timemap, error) {
	p := map[string]string{
		"userID": uID,
	}

	res := Timemap{}
	err := c.db.Get(ctx, timemapTableName, p, &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

// Create ...
func (c TimemapHandler) Create(ctx context.Context, tm Timemap) (*Timemap, error) {
	id, _ := uuid.NewV4()
	tm.ID = id.String()

	err := c.db.Put(ctx, timemapTableName, tm)
	if err != nil {
		return nil, err
	}

	return &Timemap{
		ID:        tm.ID,
		UserID:    tm.UserID,
		Map:       tm.Map,
		UpdatedAt: tm.UpdatedAt,
	}, nil
}

// Update ...
func (c TimemapHandler) Update(ctx context.Context, tm *Timemap) (*Timemap, error) {
	now := time.Now()

	p := Timemap{
		ID:        tm.ID,
		Map:       tm.Map,
		UpdatedAt: now,
	}

	keys := map[string]string{
		"userID": tm.UserID,
	}

	exp := "set info.map = :map, updatedAt = :updatedAt"

	res := Timemap{}
	err := c.db.Update(ctx, timemapTableName, keys, exp, p, &res)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
