package auth

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const (
	scopeDelimiter    = " "
	orgScopeDelimiter = ":"
)

// CustomClaims ...
type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

// GetUserClaims ...
func GetUserClaims(ts string) (*CustomClaims, error) {
	token, _ := jwt.ParseWithClaims(ts, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	if token == nil {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}

func GetOrganisationScopeClaims(ts string) (string, error) {
	c, err := GetUserClaims(ts)
	if err != nil {
		return "", fmt.Errorf("no user claims in token %w", err)
	}

	sS := strings.Split(c.Scope, scopeDelimiter)

	for _, scope := range sS {
		if !strings.HasPrefix(scope, "read:organisation") {
			continue
		}

		a := strings.Split(scope, orgScopeDelimiter)

		if len(a) != 3 {
			return "", fmt.Errorf("invalid organisation scope")
		}

		return a[2], nil
	}

	return "", fmt.Errorf("no organisation scopes present")
}

func GetOrganisationScope(scopes string) (string, error) {
	sS := strings.Split(scopes, scopeDelimiter)

	for _, scope := range sS {
		if !strings.HasPrefix(scope, "read:organisation") {
			continue
		}

		a := strings.Split(scope, orgScopeDelimiter)

		if len(a) != 3 {
			return "", fmt.Errorf("invalid organisation scope")
		}

		return a[2], nil
	}

	return "", fmt.Errorf("no organisation scopes present")
}
