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
	Get(ctx context.Context, sID, uID string) (*Progress, error)
	GetCompletedByIDs(ctx context.Context, uID string, ids ...string) ([]*Progress, error)
	Start(ctx context.Context, sID, uID string) (*Progress, error)
	Complete(ctx context.Context, sID, uID string) (*Progress, error)
}

// ProgressStore ...
type ProgressStore struct {
	db          Storer
	idGenerator IDGenerator
}

// NewNoteStore ...
func NewProgressStore(client Storer, gen IDGenerator) ProgressStore {
	return ProgressStore{
		db:          client,
		idGenerator: gen,
	}
}

// Get ...
func (c ProgressStore) Get(ctx context.Context, sID, uID string) (*Progress, error) {
	res := Progress{}
	err := c.db.Get(ctx, userTableName, UserPK(uID), ProgressSK(sID), &res)
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

	// fP := []*StepProgress{}

	// for _, p := range res {
	// 	if p.State == STATUS_COMPLETED {
	// 		fP = append(fP, p)
	// 	}
	// }

	// return fP, nil
}

// Start ...
func (c ProgressStore) Start(ctx context.Context, sID, uID string) (*Progress, error) {
	id := c.idGenerator()

	input := Progress{
		BaseEntity: BaseEntity{
			PK: UserPK(uID),
			SK: ProgressSK(sID),
		},

		ID:          id,
		EntityID:    sID,
		UserID:      uID,
		State:       STATUS_STARTED,
		DateStarted: time.Now(),
	}

	err := c.db.Put(ctx, userTableName, input)
	if err != nil {
		log.Error(fmt.Sprintf("error completing Step %s in store", sID), err)
		return nil, err
	}

	return &input, nil
}

// Complete ...
func (c ProgressStore) Complete(ctx context.Context, sID, uID string) (*Progress, error) {
	now := time.Now()
	builder := expression.NewBuilder()

	upBuilder := expression.Set(
		expression.Name("DateCompleted"),
		expression.Value(now),
	).Set(
		expression.Name("State"),
		expression.Value(STATUS_COMPLETED),
	)

	builder = builder.WithUpdate(upBuilder)

	expr, err := builder.Build()
	if err != nil {
		return nil, fmt.Errorf("error creating complete step expression %w", err)
	}

	res := Progress{}
	err = c.db.Update(ctx, userTableName, UserPK(uID), ProgressSK(sID), expr, &res)
	if err != nil {
		log.Error(fmt.Sprintf("error completing Step %s in store", sID), err)

		return nil, err
	}

	return &res, nil
}
