//go:generate mockgen -source=step_progress_handler.go -destination=./mocks/step_progress_handler.go -package=storemocks

package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/gofrs/uuid"
	"github.com/labstack/gommon/log"
)

const (
	stepProgressTableName = "user"
)

// StepProgressRepositor ...
type StepProgressRepositor interface {
	Get(ctx context.Context, sID, uID string) (*StepProgress, error)
	GetCompletedByStepID(ctx context.Context, uID string, ids ...string) ([]*StepProgress, error)
	Start(ctx context.Context, sID, uID string) (*StepProgress, error)
	Complete(ctx context.Context, sID, uID string) (*StepProgress, error)
}

// StepProgressHandler ...
type StepProgressHandler struct {
	db StorerV2
}

// NewStepProgressHandler ...
func NewStepProgressHandler(client StorerV2) StepProgressHandler {
	return StepProgressHandler{
		db: client,
	}
}

// Get ...
func (c StepProgressHandler) Get(ctx context.Context, sID, uID string) (*StepProgress, error) {
	res := StepProgress{}
	err := c.db.Get(ctx, courseProgressTableName, UserPK(uID), ProgressSK(sID), &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err

	}

	return &res, nil
}

// GetCompletedByStepID ...
func (c StepProgressHandler) GetCompletedByStepID(ctx context.Context, uID string, ids ...string) ([]*StepProgress, error) {
	res := []*StepProgress{}

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
func (c StepProgressHandler) Start(ctx context.Context, sID, uID string) (*StepProgress, error) {
	id, _ := uuid.NewV4()

	now := time.Now()
	input := StepProgress{
		ID:          id.String(),
		StepID:      sID,
		UserID:      uID,
		State:       STATUS_STARTED,
		DateStarted: &now,
	}

	err := c.db.Put(ctx, stepProgressTableName, input)
	if err != nil {
		log.Error(fmt.Sprintf("error completing Step %s in store", sID), err)
		return nil, err
	}

	return &input, nil
}

// Complete ...
func (c StepProgressHandler) Complete(ctx context.Context, sID, uID string) (*StepProgress, error) {
	now := time.Now()

	// input := map[string]interface{}{
	// 	":dateCompleted": now,
	// 	":progressState": STATUS_COMPLETED,
	// }

	// keys := map[string]string{
	// 	"stepID": c,
	// 	"userID": uID,
	// }

	builder := expression.NewBuilder()

	keyCond := expression.Key("PK").Equal(expression.Value(UserPK(uID)))
	keyCond2 := expression.Key("SK").Equal(expression.Value(ProgressSK(sID)))

	builder.WithKeyCondition(keyCond).WithKeyCondition(keyCond2)

	upBuilder := expression.Set(expression.Name("dateCompleted"), expression.IfNotExists(expression.Name("dateCompleted"), expression.Value(now))).
		Set(expression.Name("progressState"), expression.Value(STATUS_COMPLETED))

	// expression := "SET dateCompleted = :dateCompleted, progressState = :progressState"
	builder.WithUpdate(upBuilder)

	expr, err := builder.Build()
	if err != nil {
		return nil, err
	}

	res := StepProgress{}
	err = c.db.Update(ctx, stepProgressTableName, UserPK(uID), ProgressSK(sID), expr, &res)
	if err != nil {
		log.Error(fmt.Sprintf("error completing Step %s in store", sID), err)
		return nil, err
	}

	return &res, nil
}
