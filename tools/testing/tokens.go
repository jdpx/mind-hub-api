package testing

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/icrowley/fake"
)

// GenerateTestUserID generates a test Auth0 User ID
func GenerateTestUserID() string {
	return fmt.Sprintf("auth0|%s", fake.CharactersN(10))
}

// GenerateTestToken ...
func GenerateTestToken(tokenClaims jwt.MapClaims) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	return token
}

// GenerateTestTokenString ...
func GenerateTestTokenString(tokenClaims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	ts, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Printf("Error creating test token: %s", err)
	}
	return ts
}