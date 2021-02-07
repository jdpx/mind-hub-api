//go:generate mockgen -source=progress_store.go -destination=./mocks/progress_store.go -package=storemocks

package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/labstack/gommon/log"
)

// ProgressRepositor ...
type ProgressRepositor interface {
	Get(ctx context.Context, eID, uID string) (*Progress, error)
	GetCompletedByIDs(ctx context.Context, uID string, ids ...string) ([]*Progress, error)
	Start(ctx context.Context, eID, uID string) (*Progress, error)
	Complete(ctx context.Context, eID, uID string) (*Progress, error)
}

// ProgressStoreOption ...
type ProgressStoreOption func(*ProgressStore)

// ProgressStore ...
type ProgressStore struct {
	db          Storer
	idGenerator IDGenerator
	timer       Timer
}

// NewNoteStore ...
func NewProgressStore(client Storer, opts ...ProgressStoreOption) ProgressStore {
	s := ProgressStore{
		db:          client,
		idGenerator: GenerateID,
		timer:       time.Now,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

// WithProgressIDGenerator ...
func WithProgressIDGenerator(c IDGenerator) func(*ProgressStore) {
	return func(r *ProgressStore) {
		r.idGenerator = c
	}
}

// WithProgressTimer ...
func WithProgressTimer(c Timer) func(*ProgressStore) {
	return func(r *ProgressStore) {
		r.timer = c
	}
}

// Get ...
func (c ProgressStore) Get(ctx context.Context, eID, uID string) (*Progress, error) {
	res := Progress{}
	err := c.db.Get(ctx, userTableName, UserPK(uID), ProgressSK(eID), &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err

	}

	return &res, nil
}

// GetCompletedByIDs ...
func (c ProgressStore) GetCompletedByIDs(ctx context.Context, uID string, ids ...string) ([]*Progress, error) {
	res := []*Progress{}

	return res, nil

	// fmt.Println("IDS", ids)

	// builder := expression.NewBuilder()
	// keyCond := expression.Key("PK").Equal(expression.Value(UserPK(uID))).And(expression.Key(""))

	// // for _, id := range ids {
	// // keyCond2 := expression.Key("SK").Equal(expression.Value(ProgressSK(id)))

	// builder = builder.WithKeyCondition(keyCond)
	// // builder = builder.WithKeyCondition(keyCond).WithKeyCondition(keyCond2)
	// // }

	// expr, err := builder.Build()
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("44444", expr)

	// err = c.db.Query(ctx, stepProgressTableName, expr, &res)
	// if err != nil {
	// 	if errors.Is(err, ErrNotFound) {
	// 		return nil, nil
	// 	}

	// 	return nil, err
	// }

	// fP := []*progressProgress{}

	// for _, p := range res {
	// 	if p.State == STATUS_COMPLETED {
	// 		fP = append(fP, p)
	// 	}
	// }

	// return fP, nil
}

// Start ...
func (c ProgressStore) Start(ctx context.Context, eID, uID string) (*Progress, error) {
	id := c.idGenerator()

	input := Progress{
		BaseEntity: BaseEntity{
			PK: UserPK(uID),
			SK: ProgressSK(eID),
		},

		ID:          id,
		EntityID:    eID,
		UserID:      uID,
		State:       STATUS_STARTED,
		DateStarted: c.timer(),
	}

	err := c.db.Put(ctx, userTableName, input)
	if err != nil {
		log.Error(fmt.Sprintf("error completing progress %s in store", eID), err)
		return nil, err
	}

	return &Progress{
		ID:          id,
		EntityID:    eID,
		UserID:      uID,
		State:       STATUS_STARTED,
		DateStarted: c.timer(),
	}, nil
}

// Complete ...
func (c ProgressStore) Complete(ctx context.Context, eID, uID string) (*Progress, error) {
	builder := expression.NewBuilder()

	upBuilder := expression.
		Set(expression.Name("id"), expression.Name("id").IfNotExists(expression.Value(c.idGenerator()))).
		Set(expression.Name("entityID"), expression.Name("entityID").IfNotExists(expression.Value(eID))).
		Set(expression.Name("userID"), expression.Name("userID").IfNotExists(expression.Value(uID))).
		Set(expression.Name("state"), expression.Value(STATUS_COMPLETED)).
		Set(expression.Name("dateStarted"), expression.Name("dateStarted").IfNotExists(expression.Value(c.timer()))).
		Set(expression.Name("dateCompleted"), expression.Value(c.timer()))

	builder = builder.WithUpdate(upBuilder)

	expr, err := builder.Build()
	if err != nil {
		return nil, fmt.Errorf("error creating complete progress expression %w", err)
	}

	res := Progress{}
	err = c.db.Update(ctx, userTableName, UserPK(uID), ProgressSK(eID), expr, &res)
	if err != nil {
		log.Error(fmt.Sprintf("error completing progress %s in store", eID), err)

		return nil, err
	}

	return &Progress{
		ID:            res.ID,
		EntityID:      res.EntityID,
		UserID:        res.UserID,
		State:         res.State,
		DateStarted:   res.DateStarted,
		DateCompleted: res.DateCompleted,
	}, nil
}
