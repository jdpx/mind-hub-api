package request

import (
	"context"

	"github.com/google/uuid"
)

// ContextCorrelationID extracts and returns the correlation ID from a Gin Wrapped context
// If there is no CorrelationID in the context, a random ID is generated
func ContextCorrelationID(ctx context.Context) (string, error) {
	gc, err := GinContext(ctx)
	if err != nil {
		return "", err
	}

	cID := GetCorrelationIDHeader(gc.Request.Header)
	if cID == "" {
		id, _ := uuid.NewUUID()
		cID = id.String()
	}

	return cID, nil
}
