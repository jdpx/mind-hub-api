package request

import (
	"context"
	"fmt"
	"strings"

	"github.com/jdpx/mind-hub-api/pkg/auth"
)

const (
	// AuthorizationHeader ...
	AuthorizationHeader = "Authorization"
)

// GetUserID ...
func GetUserID(ctx context.Context) (string, error) {
	ts, err := authTokenFromContext(ctx)
	if err != nil {
		return "", fmt.Errorf("no auth token in context %w", err)
	}

	c, err := auth.GetUserClaims(ts)
	if err != nil {
		return "", fmt.Errorf("no user claims in token %w", err)
	}

	sub := strings.Split(c.Subject, "|")
	if len(sub) < 2 {
		return "", fmt.Errorf("token user ID is an invalid Auth0 user ID")
	}

	return sub[1], nil
}

func authTokenFromContext(ctx context.Context) (string, error) {
	gc, err := GinContext(ctx)
	if err != nil {
		return "", err
	}

	header := gc.Request.Header.Get(AuthorizationHeader)
	if header == "" {
		return "", fmt.Errorf("no authorization header present in request")
	}

	s := strings.Split(header, " ")

	return s[1], nil
}
