package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
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

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}
