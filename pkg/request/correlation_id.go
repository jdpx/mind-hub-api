package request

import (
	"context"

	"github.com/google/uuid"
)

const (
	correlationIDHeader = "X-Correlation-Id"
)

// ContextCorrelationID ...
func ContextCorrelationID(ctx context.Context) (string, error) {
	gc, err := GinContext(ctx)
	if err != nil {
		return "", err
	}

	cID := gc.Request.Header.Get(correlationIDHeader)
	if cID == "" {
		id, _ := uuid.NewUUID()
		cID = id.String()
	}

	return cID, nil
}
