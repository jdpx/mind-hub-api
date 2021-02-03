//go:generate mockgen -source=timemap_store.go -destination=./mocks/timemap_store.go -package=storemocks

package store

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
)

// TimemapRepositor ...
type TimemapRepositor interface {
	Get(ctx context.Context, uID string) (*Timemap, error)
	Create(ctx context.Context, tm Timemap) (*Timemap, error)
	Update(ctx context.Context, tm *Timemap) (*Timemap, error)
}

// TimemapStoreOption ...
type TimemapStoreOption func(*TimemapStore)

// TimemapStore ...
type TimemapStore struct {
	db          Storer
	idGenerator IDGenerator
	timer       Timer
}

// NewTimemapStore ...
func NewTimemapStore(client Storer, opts ...TimemapStoreOption) TimemapStore {
	s := TimemapStore{
		db:          client,
		idGenerator: GenerateID,
		timer:       time.Now,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

// WithTimemapIDGenerator ...
func WithTimemapIDGenerator(c IDGenerator) func(*TimemapStore) {
	return func(r *TimemapStore) {
		r.idGenerator = c
	}
}

// WithTimemapTimer ...
func WithTimemapTimer(c Timer) func(*TimemapStore) {
	return func(r *TimemapStore) {
		r.timer = c
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
	tm.ID = c.idGenerator()
	tm.DateUpdated = c.timer()

	tm.BaseEntity = BaseEntity{
		PK: UserPK(tm.UserID),
		SK: TimemapSK(),
	}

	err := c.db.Put(ctx, userTableName, tm)
	if err != nil {
		return nil, err
	}

	return &Timemap{
		ID:          tm.ID,
		UserID:      tm.UserID,
		Map:         tm.Map,
		DateUpdated: tm.DateUpdated,
	}, nil
}

// Update ...
func (c TimemapStore) Update(ctx context.Context, tm *Timemap) (*Timemap, error) {
	upBuilder := expression.Set(
		expression.Name("map"),
		expression.Value(tm.Map),
	).Set(
		expression.Name("dateUpdated"),
		expression.Value(c.timer()),
	)

	expr, err := expression.NewBuilder().WithUpdate(upBuilder).Build()
	if err != nil {
		return nil, err
	}

	res := Timemap{}
	err = c.db.Update(ctx, userTableName, UserPK(tm.UserID), TimemapSK(), expr, &res)
	if err != nil {
		return nil, err
	}

	return &Timemap{
		ID:          res.ID,
		UserID:      res.UserID,
		Map:         res.Map,
		DateUpdated: res.DateUpdated,
	}, nil
}
