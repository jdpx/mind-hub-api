package logging

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/sirupsen/logrus"
)

const (
	// APIVersionKey ...
	APIVersionKey = "api_version"
	// APIEnvironmentKey ...
	APIEnvironmentKey = "api_environment"
	// APICMSMapping ...
	APICMSMapping = "api_cms_mapping"
	// CorrelationIDKey ...
	CorrelationIDKey = "correlation_id"
	// OrganisationIDKey ...
	OrganisationIDKey = "organisation_id"
	// QueryKey ...
	QueryKey = "query"
	// CourseIDKey ...
	CourseIDKey = "course_id"
	// SessionKey ...
	SessionIDKey = "session_id"
	// StepKey ...
	StepIDKey = "step_id"
	// UserKey ...
	UserIDKey = "user_id"
	// PKKey ...
	PKKey = "pk"
	// SKKey ...
	SKKey = "sk"

	// RequestDurationKey ...
	RequestDurationKey = "request_duration"
	// HTTPMethodKey ...
	HTTPMethodKey = "http_method"
	// HTTPStatusKey ...
	HTTPStatusKey = "http_status"
	// HTTPClientNameKey ...
	HTTPClientNameKey = "http_client_name"
	// HTTPURLKey ...
	HTTPURLKey = "http_url"
)

// New ...
func New() *logrus.Entry {
	log := logrus.New()
	// log.SetFormatter(&logrus.JSONFormatter{})

	log.SetFormatter(&logrus.TextFormatter{
		// DisableColors: true,
		FullTimestamp: true,
	})

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
