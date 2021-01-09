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
	GetTimemap(ctx context.Context, uID string) (*Timemap, error)
	CreateTimemap(ctx context.Context, note Timemap) (*Timemap, error)
	UpdateTimemap(ctx context.Context, note *Timemap) (*Timemap, error)
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

// GetTimemap ...
func (c TimemapHandler) GetTimemap(ctx context.Context, uID string) (*Timemap, error) {
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

// CreateTimemap ...
func (c TimemapHandler) CreateTimemap(ctx context.Context, note Timemap) (*Timemap, error) {
	id, _ := uuid.NewV4()
	note.ID = id.String()

	err := c.db.Put(ctx, timemapTableName, note)
	if err != nil {
		return nil, err
	}

	return &Timemap{
		ID:        note.ID,
		UserID:    note.UserID,
		Map:       note.Map,
		UpdatedAt: note.UpdatedAt,
	}, nil
}

// UpdateTimemap ...
func (c TimemapHandler) UpdateTimemap(ctx context.Context, note *Timemap) (*Timemap, error) {
	now := time.Now()

	p := Timemap{
		ID:        note.ID,
		Map:       note.Map,
		UpdatedAt: now,
	}

	keys := map[string]string{
		"userID": note.UserID,
	}

	exp := "set info.map = :map, updatedAt = :updatedAt"

	res := Timemap{}
	err := c.db.Update(ctx, timemapTableName, keys, exp, p, &res)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
