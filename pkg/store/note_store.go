//go:generate mockgen -source=note_store.go -destination=./mocks/note_store.go -package=storemocks

package store

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
)

const (
	userTableName = "user"
)

// NoteRepositor ...
type NoteRepositor interface {
	Get(ctx context.Context, id, uID string) (*Note, error)
	Create(ctx context.Context, note Note) (*Note, error)
	Update(ctx context.Context, note Note) (*Note, error)
}

// NoteStore ...
type NoteStore struct {
	db          Storer
	idGenerator IDGenerator
	timer       Timer
}

// NoteStoreOption ...
type NoteStoreOption func(*NoteStore)

// NewNoteStore ...
func NewNoteStore(client Storer, opts ...NoteStoreOption) NoteStore {
	s := NoteStore{
		db:          client,
		idGenerator: GenerateID,
		timer:       time.Now,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

// WithNoteIDGenerator ...
func WithNoteIDGenerator(c IDGenerator) func(*NoteStore) {
	return func(r *NoteStore) {
		r.idGenerator = c
	}
}

// WithNoteTimer ...
func WithNoteTimer(c Timer) func(*NoteStore) {
	return func(r *NoteStore) {
		r.timer = c
	}
}

// Get ...
func (c NoteStore) Get(ctx context.Context, id, uID string) (*Note, error) {
	res := Note{}
	err := c.db.Get(ctx, userTableName, UserPK(uID), NoteSK(id), &res)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

// Create ...
func (c NoteStore) Create(ctx context.Context, note Note) (*Note, error) {
	id := c.idGenerator()
	note.ID = id
	note.DateCreated = c.timer()
	note.DateUpdated = c.timer()

	note.BaseEntity = BaseEntity{
		PK: UserPK(note.UserID),
		SK: NoteSK(note.EntityID),
	}

	err := c.db.Put(ctx, userTableName, note)
	if err != nil {
		return nil, err
	}

	return &Note{
		ID:          note.ID,
		EntityID:    note.EntityID,
		UserID:      note.UserID,
		Value:       note.Value,
		DateCreated: note.DateCreated,
		DateUpdated: note.DateUpdated,
	}, nil
}

// Update ...
func (c NoteStore) Update(ctx context.Context, note Note) (*Note, error) {
	upBuilder := expression.
		Set(expression.Name("id"), expression.Name("id").IfNotExists(expression.Value(c.idGenerator()))).
		Set(expression.Name("entityID"), expression.Name("entityID").IfNotExists(expression.Value(note.EntityID))).
		Set(expression.Name("value"), expression.Value(note.Value)).
		Set(expression.Name("dateCreated"), expression.Name("dateCreated").IfNotExists(expression.Value(c.timer()))).
		Set(expression.Name("dateUpdated"), expression.Value(c.timer()))

	expr, err := expression.NewBuilder().WithUpdate(upBuilder).Build()
	if err != nil {
		return nil, err
	}

	res := Note{}
	err = c.db.Update(ctx, userTableName, UserPK(note.UserID), NoteSK(note.EntityID), expr, &res)
	if err != nil {
		return nil, err
	}

	p := Note{
		ID:          res.ID,
		EntityID:    res.EntityID,
		UserID:      res.UserID,
		Value:       res.Value,
		DateCreated: res.DateCreated,
		DateUpdated: res.DateUpdated,
	}

	return &p, nil
}
