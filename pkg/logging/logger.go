package logging

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/sirupsen/logrus"
)

const (
	correlationIDHeader = "X-Correlation-Id"

	// CorrelationIDKey ...
	CorrelationIDKey = "correlation_id"
	// QueryKey ...
	QueryKey = "query"
	// RequestDurationKey ...
	RequestDurationKey = "request_duration"
	// HTTPStatusKey ...
	HTTPStatusKey = "http_status"
)

// New ...
func New() *logrus.Entry {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	id, _ := uuid.NewUUID()
	cID := id.String()

	return log.WithField(CorrelationIDKey, cID)
}

// NewFromResolver ...
func NewFromResolver(ctx context.Context) *logrus.Entry {
	log := New()

	cID, err := request.ContextCorrelationID(ctx)
	if err != nil {
		log.WithContext(ctx)
	}

	return log.WithContext(ctx).
		WithField(CorrelationIDKey, cID)
}

// NewWithContext ...
func NewWithContext(ctx context.Context) *logrus.Entry {
	log := New()

	cID := ctx.Value(correlationIDHeader)
	if cID == "" {
		id, _ := uuid.NewUUID()
		cID = id.String()
	}

	return log.WithContext(ctx).
		WithField(CorrelationIDKey, cID)
}
