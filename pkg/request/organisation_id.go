package request

import (
	"context"
	"fmt"

	"github.com/jdpx/mind-hub-api/pkg/auth"
)

func GetOrganisationID(ctx context.Context) (string, error) {
	ts, err := getAuthTokenFromContext(ctx)
	if err != nil {
		return "", fmt.Errorf("no auth token in context %w", err)
	}

	orgScope, err := auth.GetTokenOrganisationID(ts)
	if err != nil {
		return "", fmt.Errorf("no organisation scope claim in token %w", err)
	}

	return orgScope, nil
}
