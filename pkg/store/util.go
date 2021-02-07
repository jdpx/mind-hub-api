package store

import (
	"time"

	"github.com/segmentio/ksuid"
)

type IDGenerator func() string

func GenerateID() string {
	return ksuid.New().String()
}

type Timer func() time.Time
