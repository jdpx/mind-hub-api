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
	Get(ctx context.Context, uID, cID, tID string) (*Timemap, error)
	GetByCourseID(ctx context.Context, uID, cID string) ([]Timemap, error)
	Create(ctx context.Context, tm Timemap) (*Timemap, error)
	Update(ctx context.Context, tm Timemap) (*Timemap, error)
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
func (c TimemapStore) Get(ctx context.Context, uID, cID, tID string) (*Timemap, error) {
	res := Timemap{}
	err := c.db.Get(ctx, userTableName, UserPK(uID), TimemapSK(cID, tID), &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

// GetByCourseID ...
func (c TimemapStore) GetByCourseID(ctx context.Context, uID, cID string) ([]Timemap, error) {
	expr, _ := expression.NewBuilder().
		WithKeyCondition(expression.Key("PK").Equal(expression.Value(UserPK(uID)))).
		WithFilter(expression.Name("SK").BeginsWith(CourseTimemapsSK(cID))).
		Build()

	res := []Timemap{}
	err := c.db.Query(ctx, userTableName, expr, &res)
	if err != nil {
		return []Timemap{}, err
	}

	return res, nil
}

// Create ...
func (c TimemapStore) Create(ctx context.Context, tm Timemap) (*Timemap, error) {
	tm.ID = c.idGenerator()
	tm.DateCreated = c.timer()
	tm.DateUpdated = c.timer()

	tm.BaseEntity = BaseEntity{
		PK: UserPK(tm.UserID),
		SK: TimemapSK(tm.CourseID, tm.ID),
	}

	err := c.db.Put(ctx, userTableName, tm)
	if err != nil {
		return nil, err
	}

	return &Timemap{
		ID:          tm.ID,
		CourseID:    tm.CourseID,
		UserID:      tm.UserID,
		Map:         tm.Map,
		DateCreated: tm.DateCreated,
		DateUpdated: tm.DateUpdated,
	}, nil
}

// Update ...
func (c TimemapStore) Update(ctx context.Context, tm Timemap) (*Timemap, error) {
	if tm.ID == "" {
		tm.ID = c.idGenerator()
	}

	upBuilder := expression.
		Set(expression.Name("id"), expression.Name("id").IfNotExists(expression.Value(tm.ID))).
		Set(expression.Name("courseID"), expression.Name("courseID").IfNotExists(expression.Value(tm.CourseID))).
		Set(expression.Name("userID"), expression.Name("userID").IfNotExists(expression.Value(tm.UserID))).
		Set(expression.Name("map"), expression.Value(tm.Map)).
		Set(expression.Name("dateCreated"), expression.Name("dateCreated").IfNotExists(expression.Value(c.timer()))).
		Set(expression.Name("dateUpdated"), expression.Value(c.timer()))

	expr, err := expression.NewBuilder().WithUpdate(upBuilder).Build()
	if err != nil {
		return nil, err
	}

	res := Timemap{}
	err = c.db.Update(ctx, userTableName, UserPK(tm.UserID), TimemapSK(tm.CourseID, tm.ID), expr, &res)
	if err != nil {
		return nil, err
	}

	return &Timemap{
		ID:          res.ID,
		CourseID:    res.CourseID,
		UserID:      res.UserID,
		Map:         res.Map,
		DateCreated: res.DateCreated,
		DateUpdated: res.DateUpdated,
	}, nil
}
