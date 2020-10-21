package request

import (
	"context"
	"fmt"
	"strings"

	"github.com/jdpx/mind-hub-api/pkg/auth"
)

const (
	authorizationHeader = "Authorization"
)

// AuthTokenFromContext ...
func AuthTokenFromContext(ctx context.Context) (string, error) {
	gc, err := GinContext(ctx)
	if err != nil {
		return "", err
	}

	header := gc.Request.Header.Get(authorizationHeader)
	if header == "" {
		return "", fmt.Errorf("user not authenticated")
	}

	s := strings.Split(header, " ")

	return s[1], nil
}

// GetUserID ...
func GetUserID(ctx context.Context) (string, error) {
	ts, err := AuthTokenFromContext(ctx)
	if err != nil {
		return "", err
	}

	c, err := auth.GetUserClaims(ts)
	if err != nil {
		return "", err
	}

	sub := strings.Split(c.Subject, "|")

	return sub[1], nil
}
