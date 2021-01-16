package request

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jdpx/mind-hub-api/pkg/auth"
)

const (
	subjectDelimiter = "|"
	authDelimiter    = " "
)

// GetUserID ...
func GetUserID(ctx context.Context) (string, error) {
	ts, err := AuthTokenFromContext(ctx)
	if err != nil {
		return "", fmt.Errorf("no auth token in context %w", err)
	}

	c, err := auth.GetUserClaims(ts)
	if err != nil {
		return "", fmt.Errorf("no user claims in token %w", err)
	}

	sub := strings.Split(c.Subject, subjectDelimiter)
	if len(sub) < 2 {
		return "", fmt.Errorf("token user ID is an invalid Auth0 user ID")
	}

	return sub[1], nil
}

func GetOrganisationID(ctx context.Context) (string, error) {
	ts, err := AuthTokenFromContext(ctx)
	if err != nil {
		return "", fmt.Errorf("no auth token in context %w", err)
	}

	c, err := auth.GetUserClaims(ts)
	if err != nil {
		return "", fmt.Errorf("no user claims in token %w", err)
	}

	orgScope, err := auth.GetOrganisationScope(c.Scope)
	if err != nil {
		return "", fmt.Errorf("no organisation scope claim in token %w", err)
	}

	return orgScope, nil
}

func AuthTokenFromContext(ctx context.Context) (string, error) {
	gc, err := GinContext(ctx)
	if err != nil {
		return "", err
	}

	return AuthTokenFromGinContext(gc)
}

func AuthTokenFromGinContext(gc *gin.Context) (string, error) {
	header := GetAuthorizationHeader(gc.Request.Header)
	if header == "" {
		return "", fmt.Errorf("no authorization header present in request")
	}

	s := strings.Split(header, authDelimiter)

	return s[1], nil
}
