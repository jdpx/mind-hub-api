//go:generate mockgen -source=timemap_handler.go -destination=./mocks/timemap_handler.go -package=storemocks
package store

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/gofrs/uuid"
)

const (
	timemapTableName = "user"
)

// TimemapRepositor ...
type TimemapRepositor interface {
	Get(ctx context.Context, uID string) (*Timemap, error)
	Create(ctx context.Context, tm Timemap) (*Timemap, error)
	Update(ctx context.Context, tm *Timemap) (*Timemap, error)
}

// TimemapStore ...
type TimemapStore struct {
	db Storer
}

// NewTimemapStore ...
func NewTimemapStore(client Storer) TimemapStore {
	return TimemapStore{
		db: client,
	}
}

// Get ...
func (c TimemapStore) Get(ctx context.Context, uID string) (*Timemap, error) {
	res := Timemap{}
	err := c.db.Get(ctx, userTableName, UserPK(uID), TimemapSK(), &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

// Create ...
func (c TimemapStore) Create(ctx context.Context, tm Timemap) (*Timemap, error) {
	id, _ := uuid.NewV4()
	tm.ID = id.String()

	tm.BaseEntity = BaseEntity{
		PK: UserPK(tm.UserID),
		SK: TimemapSK(),
	}

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
func (c TimemapStore) Update(ctx context.Context, tm *Timemap) (*Timemap, error) {
	now := time.Now()

	p := Timemap{
		ID:        tm.ID,
		Map:       tm.Map,
		UpdatedAt: now,
	}

	upBuilder := expression.Set(
		expression.Name("map"),
		expression.Value(tm.Map),
	).Set(
		expression.Name("updatedAt"),
		expression.Value(now),
	)

	expr, err := expression.NewBuilder().WithUpdate(upBuilder).Build()
	if err != nil {
		return nil, err
	}

	res := Timemap{}
	err = c.db.Update(ctx, timemapTableName, UserPK(tm.UserID), TimemapSK(), expr, &res)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
