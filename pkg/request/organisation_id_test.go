package request_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/icrowley/fake"
	"github.com/jdpx/mind-hub-api/pkg/request"
	tools "github.com/jdpx/mind-hub-api/tools/testing"
	"github.com/stretchr/testify/assert"
)

func TestGetOrganisationID(t *testing.T) {
	orgID := fake.CharactersN(10)
	orgIDScope := tools.GenerateTestOrganisationIDScope(orgID)
	testScopes := fmt.Sprintf("%s %s", fake.Words(), orgIDScope)

	ctx := context.Background()

	testCases := []struct {
		desc string
		ctx  context.Context

		expectedOrgID string
		expectedErr   error
	}{
		{
			desc: "given a valid Org ID scope in the scope, id returned",
			ctx: tools.GenerateTestGinContextWithToken(ctx, jwt.MapClaims{
				"scope": testScopes,
			}),

			expectedOrgID: orgID,
		},
		{
			desc: "given there is no gin context, error returned",
			ctx:  context.Background(),

			expectedErr: fmt.Errorf("no auth token in context could not retrieve gin.Context from context"),
		},
		{
			desc: "given the org ID in the token scopes, error returned",
			ctx: tools.GenerateTestGinContextWithToken(ctx, jwt.MapClaims{
				"scope": fake.Words(),
			}),

			expectedErr: fmt.Errorf("no organisation scope claim in token no organisation scopes present"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			id, err := request.GetOrganisationID(tt.ctx)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedOrgID, id)
			}
		})
	}
}
