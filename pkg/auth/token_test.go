package auth_test

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/auth"
	tTools "github.com/jdpx/mind-hub-api/tools/testing"
	"github.com/stretchr/testify/assert"
)

func TestTokenGetUserClaims(t *testing.T) {
	testScopes := fake.Words()
	testUserID := tTools.GenerateTestUserID()
	token := tTools.GenerateTestTokenString(jwt.MapClaims{
		"scope": testScopes,
		"sub":   testUserID,
	})

	testCases := []struct {
		desc        string
		tokenString string

		expectedClaims auth.CustomClaims
		expectedErr    error
	}{
		{
			desc:        "given a valid token string, correct claims returned",
			tokenString: token,

			expectedClaims: auth.CustomClaims{
				Scope: testScopes,
				StandardClaims: jwt.StandardClaims{
					Subject: testUserID,
				},
			},
		},
		{
			desc:        "given an invalid token string, error returned",
			tokenString: token[0:10],

			expectedErr: fmt.Errorf("invalid token"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			c, err := auth.GetUserClaims(tt.tokenString)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, &tt.expectedClaims, c)
			}
		})
	}
}

func TestTokenGetOrganisationID(t *testing.T) {
	orgID := fake.CharactersN(10)
	orgIDScope := tTools.GenerateTestOrganisationIDScope(orgID)
	testScopes := fmt.Sprintf("%s %s", fake.Words(), orgIDScope)

	testCases := []struct {
		desc   string
		scopes string

		expectedOrgID string
		expectedErr   error
	}{
		{
			desc: "given a valid scope containing org id, correct id returned",
			scopes: tTools.GenerateTestTokenString(jwt.MapClaims{
				"scope": testScopes,
			}),

			expectedOrgID: orgID,
		},
		{
			desc:   "given an invalid token string, error returned",
			scopes: "testing",

			expectedErr: fmt.Errorf("no user claims in token invalid token"),
		},
		{
			desc: "given no org scope appears in scopes, error returned",
			scopes: tTools.GenerateTestTokenString(jwt.MapClaims{
				"scope": fake.Words(),
			}),

			expectedErr: fmt.Errorf("no organisation scopes present"),
		},
		{
			desc: "given an invalid org scope, error returned",
			scopes: tTools.GenerateTestTokenString(jwt.MapClaims{
				"scope": fmt.Sprintf("read:organisation:%s:foo", orgID),
			}),

			expectedErr: fmt.Errorf("invalid organisation scope"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			id, err := auth.GetTokenOrganisationID(tt.scopes)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedOrgID, id)
			}
		})
	}
}
